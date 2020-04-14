package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ONSdigital/dp-recipe-api/recipe"
	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/mux"
)

//RecipeListHandler - get all recipes
// USAGE: curl -X GET http://localhost:22300/recipes
func (api *RecipeAPI) RecipeListHandler(w http.ResponseWriter, req *http.Request) {
	var list recipe.List
	if api.EnableMongoData {
		var err error
		list.Items, err = api.dataStore.Backend.GetRecipes()
		if err != nil {
			log.Event(req.Context(), "error getting recipes from mongo", log.ERROR, log.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		list = recipe.FullList
	}

	c := len(list.Items)
	list.Count = c
	list.TotalCount = c
	list.Limit = c

	b, err := json.Marshal(list)
	if err != nil {
		log.Event(req.Context(), "error returned from json marshall", log.ERROR, log.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

//RecipeHandler as existed in the hardcoded-only version of this service.
func (api *RecipeAPI) RecipeHandler(w http.ResponseWriter, req *http.Request) {
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

//HealthCheck - health check endpoint
func (api *RecipeAPI) HealthCheck(w http.ResponseWriter, req *http.Request) {
	// Set status to 200 OK
	w.WriteHeader(200)
}
