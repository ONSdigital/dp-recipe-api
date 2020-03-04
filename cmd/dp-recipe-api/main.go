package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"syscall"

	"github.com/ONSdigital/go-ns/server"
	"github.com/ONSdigital/log.go/log"

	"os"
	"os/signal"

	"github.com/ONSdigital/dp-recipe-api/config"
	"github.com/ONSdigital/dp-recipe-api/mongo"
	"github.com/ONSdigital/dp-recipe-api/recipe"
	"github.com/gorilla/mux"
)

func main() {
	log.Namespace = "recipe-api"
	ctx := context.Background()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	cfg, configErr := config.Get()
	if configErr != nil {
		log.Event(ctx, "error loading app config", log.FATAL, log.Error(configErr))
		os.Exit(1)
	}

	log.Event(ctx, "config on startup", log.INFO, log.Data{"config": cfg})

	mongodb := &mongo.Mongo{
		Collection: cfg.MongoConfig.Collection,
		Database:   cfg.MongoConfig.Database,
		URI:        cfg.MongoConfig.BindAddr,
	}

	var err error
	mongodb.Session, err = mongodb.Init()
	if err != nil {
		log.Event(ctx, "failed to initialise mongo", log.FATAL, log.Error(err))
		os.Exit(1)
	}

	router := mux.NewRouter()
	router.Methods("GET").Path("/recipes").HandlerFunc(recipeListHandler)
	router.Methods("GET").Path("/recipes/{id}").HandlerFunc(recipeHandler)

	log.Event(ctx, "starting http server", log.INFO, log.Data{"bind_addr": cfg.BindAddr})
	srv := server.New(cfg.BindAddr, router)
	if err := srv.ListenAndServe(); err != nil {
		log.Event(ctx, "error starting http server for API", log.FATAL, log.Error(err))
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
		log.Event(req.Context(), "error returned from json marshall", log.ERROR, log.Error(err))
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
		log.Event(req.Context(), "recipe not found", log.ERROR, log.Error(errors.New("recipe not found")), logD)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	b, err := json.Marshal(r)
	if err != nil {
		log.Event(req.Context(), "error returned from json marshall", log.ERROR, log.Error(err), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
