package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ONSdigital/go-ns/log"
	"github.com/ONSdigital/go-ns/server"

	"os"

	"github.com/ONSdigital/dp-recipe-api/config"
	"github.com/ONSdigital/dp-recipe-api/recipe"
	"github.com/gorilla/mux"
)

func main() {
	log.Namespace = "dp-recipe-api"
	configuration, configErr := config.Get()
	if configErr != nil {
		log.Error(configErr, nil)
		os.Exit(1)
	}

	router := mux.NewRouter()
	router.Path("/recipes").HandlerFunc(recipeListHandler)
	router.Path("/recipes/{id}").HandlerFunc(recipeHandler)

	log.Debug("starting http server", log.Data{"bind_addr": configuration.BindAddr})
	srv := server.New(configuration.BindAddr, router)
	if err := srv.ListenAndServe(); err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}
}

func recipeListHandler(w http.ResponseWriter, req *http.Request) {
	b, err := json.Marshal(&recipe.FullList)
	if err != nil {
		log.ErrorR(req, err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func recipeHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	recipeID := vars["id"]

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
