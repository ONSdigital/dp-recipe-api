package api

import (
	"encoding/json"
	"net/http"

	errs "github.com/ONSdigital/dp-recipe-api/apierrors"
	"github.com/ONSdigital/dp-recipe-api/recipe"
	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/mux"
)

//RecipeListHandler - get all recipes
func (api *RecipeAPI) RecipeListHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	var list recipe.List
	if api.EnableMongoData {
		var err error
		list.Items, err = api.dataStore.Backend.GetRecipes(ctx)
		if err != nil && err != errs.ErrRecipesNotFound {
			log.Event(ctx, "error getting recipes from mongo", log.ERROR, log.Error(err))
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
		log.Event(ctx, "error returned from json marshal", log.ERROR, log.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

//RecipeHandler - get recipe by ID
func (api *RecipeAPI) RecipeHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)
	recipeID := vars["id"]
	logD := log.Data{"recipe_id": recipeID}

	var r recipe.Response
	if api.EnableMongoData {

		recipe, err := api.dataStore.Backend.GetRecipe(recipeID)
		if err == errs.ErrRecipeNotFound {
			log.Event(ctx, "recipe not found in mongo", log.ERROR, log.Error(err))
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err != nil {
			log.Event(ctx, "error getting recipe from mongo", log.ERROR, log.Error(err), logD)
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
			log.Event(ctx, "recipe not found", log.ERROR, logD)
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}

	b, err := json.Marshal(r)
	if err != nil {
		log.Event(ctx, "error returned from json marshall", log.ERROR, log.Error(err), logD)
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
func (api *RecipeAPI) AddAllRecipeHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	for _, item := range recipe.FullList.Items {
		err := api.dataStore.Backend.AddRecipe(item)
		if err != nil {
			log.Event(ctx, "error in adding all recipes to mongo from data.go", log.ERROR, log.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
