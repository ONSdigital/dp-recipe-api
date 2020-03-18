package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ONSdigital/go-ns/common"
	"github.com/ONSdigital/go-ns/server"
	"github.com/ONSdigital/log.go/log"

	"os"

	"github.com/ONSdigital/dp-api-audit-spike/auditing"
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

	auditor := &auditing.Stub{}

	router := mux.NewRouter()
	router.Path("/recipes").Handler(recipeListHandlerWithAudit(auditor))
	router.Path("/recipes/{id}").Handler(recipeHandlerWithAudit(auditor))

	log.Event(context.Background(), "starting http server", log.Data{"bind_addr": configuration.BindAddr})
	srv := server.New(configuration.BindAddr, router)
	if err := srv.ListenAndServe(); err != nil {
		log.Event(context.Background(), "error starting http server for API", log.Error(err))
		os.Exit(1)
	}
}

func recipeListHandlerWithAudit(auditor auditing.Service) http.Handler {
	return auditing.Wrap(recipeListHandler, "list recipes", auditor, nil, 200)
}

func recipeListHandler(w http.ResponseWriter, r *http.Request) {
	list := &recipe.FullList
	c := len(list.Items)
	list.Count = c
	list.TotalCount = c
	list.ItemsPerPage = c

	b, err := json.Marshal(list)
	if err != nil {
		log.Event(r.Context(), "error returned from json marshall", log.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func recipeHandlerWithAudit(auditor auditing.Service) http.Handler {
	getAuditParams := func(r *http.Request) common.Params {
		return common.Params{"recipe_id": mux.Vars(r)["id"]}
	}

	return auditing.Wrap(recipeHandler, "get recipe", auditor, getAuditParams, 200)
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
