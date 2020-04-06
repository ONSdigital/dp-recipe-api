package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	errs "github.com/ONSdigital/dp-recipe-api/apierrors"
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

//RecipeHandler - get recipe by ID
// USAGE: curl -X GET http://localhost:22300/recipes/{id}
func (api *RecipeAPI) RecipeHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	recipeID := vars["id"]
	logD := log.Data{"recipe_id": recipeID}

	var r recipe.Response
	if api.EnableMongoData {

		recipe, err := api.dataStore.Backend.GetRecipe(recipeID)
		if err != nil {
			log.Event(req.Context(), "recipe not found", log.ERROR, logD)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		r = *recipe

	} else {

		found := false

		for _, item := range recipe.FullList.Items {
			if item.ID == recipeID {
				r = item
				found = true
				break
			}
		}

		if !found {
			log.Event(req.Context(), "recipe not found", log.ERROR, logD)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
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

//AddAllRecipeHandler - Adds all the recipes from data.go to the mongo database
//USAGE: curl -X POST http://localhost:22300/allrecipes
func (api *RecipeAPI) AddAllRecipeHandler(w http.ResponseWriter, req *http.Request) {
	for _, item := range recipe.FullList.Items {
		err := api.dataStore.Backend.AddRecipe(item)
		if err != nil {
			log.Event(req.Context(), "error in adding all recipes to mongo from data.go", log.ERROR, log.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

//AddRecipeHandler - adds recipes
// USAGE: curl -X POST http://localhost:22300/recipes -d '{"alias":"Hello"}'
// USAGE: curl -X POST http://localhost:22300/recipes -d "@data.json"
func (api *RecipeAPI) AddRecipeHandler(w http.ResponseWriter, req *http.Request) {

	// Read body
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal to the shape of Response struct
	var recipe recipe.Response
	err = json.Unmarshal(b, &recipe)
	if err != nil {
		http.Error(w, errs.ErrAddUpdateRecipeBadRequest.Error(), 500)
		return
	}

	// Add Recipe to Mongo
	err = api.dataStore.Backend.AddRecipe(recipe)
	if err != nil {
		fmt.Println(err)
	}

	// Marshal to output the newly added recipe
	output, err := json.Marshal(recipe)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
