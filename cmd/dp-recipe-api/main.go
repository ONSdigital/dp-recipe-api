package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ONSdigital/go-ns/server"
	"github.com/ONSdigital/log.go/log"

	"os"

	"github.com/ONSdigital/dp-recipe-api/config"
	"github.com/ONSdigital/dp-recipe-api/recipe"
	"github.com/gorilla/mux"
)

func main() {
	log.Namespace = "recipe-api"
	configuration, configErr := config.Get()
	if configErr != nil {
		log.Event(context.Background(), "error loading app config", log.Error(configErr))
		os.Exit(1)
	}

	router := mux.NewRouter()
	router.Path("/recipes").HandlerFunc(recipeListHandler)
	router.Path("/recipes/{id}").HandlerFunc(recipeHandler)

	log.Event(context.Background(), "starting http server", log.Data{"bind_addr": configuration.BindAddr})
	srv := server.New(configuration.BindAddr, router)
	if err := srv.ListenAndServe(); err != nil {
		log.Event(context.Background(), "error starting http server for API", log.Error(err))
		os.Exit(1)
	}
}

func recipeListHandler(w http.ResponseWriter, req *http.Request) {
	list := &recipe.FullList
	c := len(list.Items)
	list.Count = c
	list.TotalCount = c
	list.ItemsPerPage = c

	b, err := json.Marshal(list)
	if err != nil {
		log.Event(req.Context(), "error returned from json marshall", log.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func recipeHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	recipeID := vars["id"]
	logD := log.Data{"recipe_id": recipeID}

	var r recipe.Response
	found := false

	for _, item := range recipe.FullList.Items {
		if item.ID == recipeID {
			r = item
			found = true
			break
		}
	}

	if !found {
		log.Event(req.Context(), "recipe not found", log.Error(errors.New("recipe not found")), logD)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	b, err := json.Marshal(r)
	if err != nil {
		log.Event(req.Context(), "error returned from json marshall", log.Error(err), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
