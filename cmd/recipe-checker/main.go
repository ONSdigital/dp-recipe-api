package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/ONSdigital/go-ns/log"
	"github.com/ONSdigital/go-ns/server"

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
		log.Info("URLs must be provided for app to start", log.Data{"dev": devURLFlag, "beta": betaURLFlag})
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
		data := getRecipeList()
		bootstrap.ExecuteTemplate(w, "bootstrap", data)
	})

	router.Path("/recipes").HandlerFunc(recipesStatusListHandler)
	router.Path("/recipes/{recipe}").HandlerFunc(recipesStatusHandler)
	router.Path("/recipes/{recipe}/codelists").HandlerFunc(codelistsListHandler)

	log.Debug("starting http server", log.Data{"bind_addr": *bindAddrFlag})
	srv := server.New(*bindAddrFlag, router)
	if err := srv.ListenAndServe(); err != nil {
		log.Error(err, nil)
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
	newList := getRecipeList()

	b, err := json.Marshal(newList)
	if err != nil {
		log.ErrorR(req, err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func getRecipeList() *status.RecipeList {
	origList := recipe.FullList
	newList := &status.RecipeList{}
	for _, i := range origList.Items {
		r := status.Recipe{
			ID:    i.ID,
			Alias: i.Alias,
		}

		for _, o := range i.OutputInstances {
			add := status.CheckRecipe(devURL, betaURL, o.DatasetID, o.CodeLists)
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
	vars := mux.Vars(req)
	recipeID := vars["recipe"]

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
			add := status.CheckRecipe(devURL, betaURL, o.DatasetID, o.CodeLists)
			add.DatasetName = o.Title
			r.Outputs = append(r.Outputs, *add)
		}
	}

	if r == nil {
		log.ErrorR(req, errors.New("recipe not found"), nil)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	b, err := json.Marshal(r)
	if err != nil {
		log.ErrorR(req, err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

// list codelists for a specific recipe and their status per env
func codelistsListHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	recipeID := vars["recipe"]

	newList := getCodelists(recipeID)

	b, err := json.Marshal(newList)
	if err != nil {
		log.ErrorR(req, err, nil)
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
