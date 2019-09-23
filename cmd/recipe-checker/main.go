package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/ONSdigital/dp-api-clients-go/headers"
	"github.com/ONSdigital/go-ns/server"
	"github.com/ONSdigital/log.go/log"

	"os"

	"github.com/ONSdigital/dp-recipe-api/recipe"
	"github.com/ONSdigital/dp-recipe-api/status"
	"github.com/gorilla/mux"
)

var bootstrap *template.Template
var devURL string
var betaURL string
var urlFMT = "https://%s/v1"

var devURLFlag = flag.String("dev", "", "the hostname for the develop environment")
var betaURLFlag = flag.String("beta", "", "the hostname for the production environment")
var bindAddrFlag = flag.String("bind", ":2222", "the desired port for the application to run on")

func main() {
	log.Namespace = "recipe-checker"
	flag.Parse()

	if len(*devURLFlag) == 0 || len(*betaURLFlag) == 0 {
		log.Event(context.Background(), "URLs must be provided for app to start", log.Data{"dev": devURLFlag, "beta": betaURLFlag})
		os.Exit(1)
	}

	devURL = fmt.Sprintf(urlFMT, *devURLFlag)
	betaURL = fmt.Sprintf(urlFMT, *betaURLFlag)

	var err error

	bootstrap, err = template.New("checker").Funcs(template.FuncMap{
		"getCodelists": getCodelists,
	}).ParseFiles(layoutFiles()...)

	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.Path("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		userAuthToken, _ := headers.GetUserAuthToken(req)
		serviceAuthToken, _ := headers.GetServiceAuthToken(req)
		collectionID, _ := headers.GetCollectionID(req)

		data := getRecipeList(ctx, userAuthToken, serviceAuthToken, collectionID)
		bootstrap.ExecuteTemplate(w, "bootstrap", data)
	})

	router.Path("/recipes").HandlerFunc(recipesStatusListHandler)
	router.Path("/recipes/{recipe}").HandlerFunc(recipesStatusHandler)
	router.Path("/recipes/{recipe}/codelists").HandlerFunc(codelistsListHandler)

	log.Event(context.Background(), "starting http server", log.Data{"bind_addr": *bindAddrFlag})
	srv := server.New(*bindAddrFlag, router)
	if err := srv.ListenAndServe(); err != nil {
		log.Event(context.Background(), "error starting http server", log.Error(err))
		os.Exit(1)
	}
}

func layoutFiles() []string {
	files, err := filepath.Glob("pages/*.gohtml")
	if err != nil {
		panic(err)
	}
	return files
}

//list each recipe and a status per env
func recipesStatusListHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	userAuthToken, _ := headers.GetUserAuthToken(req)
	serviceAuthToken, _ := headers.GetServiceAuthToken(req)
	collectionID, _ := headers.GetCollectionID(req)

	newList := getRecipeList(ctx, userAuthToken, serviceAuthToken, collectionID)

	b, err := json.Marshal(newList)
	if err != nil {
		log.Event(ctx, "error returned from json marshall", log.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func getRecipeList(ctx context.Context, userAuthToken, serviceAuthToken, collectionID string) *status.RecipeList {
	origList := recipe.FullList
	newList := &status.RecipeList{}
	for _, i := range origList.Items {
		r := status.Recipe{
			ID:    i.ID,
			Alias: i.Alias,
		}

		for _, o := range i.OutputInstances {
			checkReq := status.CheckRequest{
				UserAuthToken:    userAuthToken,
				ServiceAuthToken: serviceAuthToken,
				DevURL:           devURL,
				BetaURL:          betaURL,
				CollectionID:     collectionID,
				DatasetID:        o.DatasetID,
				CodeLists:        o.CodeLists,
			}

			add := status.CheckRecipe(ctx, checkReq)
			add.DatasetName = o.Title
			r.Outputs = append(r.Outputs, *add)
		}

		newList.Items = append(newList.Items, r)
	}

	c := len(newList.Items)
	newList.Count = c
	newList.TotalCount = c
	newList.ItemsPerPage = c
	return newList
}

//return a specific recipe and its statuses
func recipesStatusHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)
	recipeID := vars["recipe"]
	logD := log.Data{"recipe_id": recipeID}
	userAuthToken, _ := headers.GetUserAuthToken(req)
	serviceAuthToken, _ := headers.GetServiceAuthToken(req)
	collectionID, _ := headers.GetCollectionID(req)

	var r *status.Recipe
	origList := recipe.FullList
	for _, i := range origList.Items {
		if i.ID != recipeID {
			continue
		}

		r = &status.Recipe{
			ID:    i.ID,
			Alias: i.Alias,
		}

		for _, o := range i.OutputInstances {
			checkReq := status.CheckRequest{
				UserAuthToken:    userAuthToken,
				ServiceAuthToken: serviceAuthToken,
				DevURL:           devURL,
				BetaURL:          betaURL,
				CollectionID:     collectionID,
				DatasetID:        o.DatasetID,
				CodeLists:        o.CodeLists,
			}

			add := status.CheckRecipe(ctx, checkReq)
			add.DatasetName = o.Title
			r.Outputs = append(r.Outputs, *add)
		}
	}

	if r == nil {
		log.Event(ctx, "recipe not found", log.Error(errors.New("recipe not found")), logD)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	b, err := json.Marshal(r)
	if err != nil {
		log.Event(ctx, "error returned from json marshal", log.Error(err), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

// list codelists for a specific recipe and their status per env
func codelistsListHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)
	recipeID := vars["recipe"]
	logD := log.Data{"recipe_id": recipeID}

	newList := getCodelists(recipeID)

	b, err := json.Marshal(newList)
	if err != nil {
		log.Event(ctx, "error returned from json marshal", log.Error(err), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func getCodelists(recipeID string) status.CodelistList {
	origList := recipe.FullList
	newList := &status.CodelistList{}
	for _, i := range origList.Items {
		if i.ID != recipeID {
			continue
		}

		var instances sync.WaitGroup
		for _, o := range i.OutputInstances {
			instances.Add(1)
			var wg sync.WaitGroup

			for _, codelist := range o.CodeLists {
				wg.Add(1)

				c := &status.Codelist{
					ID:          codelist.ID,
					Name:        codelist.Name,
					IsHierarchy: codelist.IsHierarchy,
				}

				go func(c *status.Codelist) {
					c.CheckCodelist(devURL, betaURL, &wg)
					newList.Items = append(newList.Items, *c)
				}(c)
			}
			wg.Wait()
			instances.Done()
		}
		instances.Wait()
	}

	c := len(newList.Items)
	newList.Count = c
	newList.TotalCount = c
	newList.ItemsPerPage = c

	return *newList
}
