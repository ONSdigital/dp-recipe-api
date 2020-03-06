package api

import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/dp-recipe-api/recipe"
	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

//RecipeListHandler - get all recipes
// USAGE: curl -X GET http://localhost:22300/recipes
func RecipeListHandler(w http.ResponseWriter, req *http.Request) {
	list := &recipe.FullList
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

//RecipeHandler - get recipe by ID
// USAGE: curl -X GET http://localhost:22300/recipes/{id}
func RecipeHandler(w http.ResponseWriter, req *http.Request) {
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
