package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"

	errs "github.com/ONSdigital/dp-recipe-api/apierrors"
	"github.com/ONSdigital/dp-recipe-api/recipe"
	"github.com/ONSdigital/log.go/log"
	"github.com/google/uuid"
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
		log.Event(ctx, "error returned from json marshal", log.ERROR, log.Error(err), logD)
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

//AddRecipeHandler - Add a Recipe
func (api *RecipeAPI) AddRecipeHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	// Read body
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.Event(ctx, "error in reading request body", log.ERROR, log.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Unmarshal to the shape of Response struct
	var recipe recipe.Response
	err = json.Unmarshal(b, &recipe)
	if err != nil {
		log.Event(ctx, "error returned from json unmarshal", log.ERROR, log.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Randomly generated version 4 UUID for recipe ID
	recipe.ID = uuid.UUID.String(uuid.New())

	//Validation to check if all the recipe fields are entered
	recipeValueNil := CheckRecipeValuesNil(recipe)
	if recipeValueNil {
		log.Event(ctx, "bad request error as one field value of recipe is nil in request body", log.ERROR)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Add Recipe to Mongo
	err = api.dataStore.Backend.AddRecipe(recipe)
	if err != nil {
		log.Event(ctx, "error adding recipe to mongo", log.ERROR, log.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Marshal to output the newly added recipe
	output, err := json.Marshal(recipe)
	if err != nil {
		log.Event(ctx, "error returned from json marshal", log.ERROR, log.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

//UpdateRecipeHandler - update specific recipe by ID
func (api *RecipeAPI) UpdateRecipeHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	//Get Update Recipe ID
	vars := mux.Vars(req)
	recipeID := vars["id"]
	logD := log.Data{"recipe_id": recipeID}

	// Read body
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.Event(ctx, "error in reading request body", log.ERROR, log.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Unmarshal to the shape of Response struct
	var recipeUpdate recipe.Response
	err = json.Unmarshal(b, &recipeUpdate)
	if err != nil {
		log.Event(ctx, "error returned from json unmarshal", log.ERROR, log.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	recipeUpdate.ID = recipeID

	// Use exisiting (old) recipe and fills in the fields of the update recipe with data from exisiting when particular field is nil
	var oldRecipe *recipe.Response
	oldRecipe, err = api.dataStore.Backend.GetRecipe(recipeID)
	if err == errs.ErrRecipeNotFound {
		log.Event(ctx, "recipe not found in mongo", log.ERROR, log.Error(err), logD)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		log.Event(ctx, "error getting recipe from mongo", log.ERROR, log.Error(err), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	recipeUpdate = UpdateRecipeFieldsFromPrevious(recipeUpdate, *oldRecipe)

	// Validation to check if all the recipe fields are entered
	recipeValueNil := CheckRecipeValuesNil(recipeUpdate)
	if recipeValueNil {
		log.Event(ctx, "bad request error as one field value of updated recipe is nil in request body", log.ERROR)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Update Recipe to Mongo
	err = api.dataStore.Backend.UpdateRecipe(recipeID, recipeUpdate)
	if err != nil {
		log.Event(ctx, "error updating recipe to mongo", log.ERROR, log.Error(err), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Marshal to output the newly updated recipe
	updatedRecipe, _ := api.dataStore.Backend.GetRecipe(recipeID)
	output, err := json.Marshal(updatedRecipe)
	if err != nil {
		log.Event(ctx, "error getting updated recipe from mongo", log.ERROR, log.Error(err), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

//CheckRecipeValuesNil - checks whether the values of each field in recipe are nil
func CheckRecipeValuesNil(recipe recipe.Response) (recipeValueNil bool) {
	recipeElem := reflect.ValueOf(&recipe).Elem()
	for i := 0; i < recipeElem.NumField(); i++ {
		recipeValueNil = recipeElem.Field(i).IsZero()
		if recipeValueNil {
			return recipeValueNil
		}
	}
	return recipeValueNil
}

//UpdateRecipeFieldsFromPrevious - uses the exisiting recipe and fills in the fields of the update recipe with data from exisiting when field is nil
func UpdateRecipeFieldsFromPrevious(updates recipe.Response, oldRecipe recipe.Response) (updateRecipe recipe.Response) {
	recipeElem := reflect.ValueOf(&updates).Elem()
	for i := 0; i < recipeElem.NumField(); i++ {
		recipeValueNil := recipeElem.Field(i).IsZero()
		if recipeValueNil {
			switch i {
			case 1:
				updates.Alias = oldRecipe.Alias
			case 2:
				updates.Format = oldRecipe.Format
			case 3:
				updates.InputFiles = oldRecipe.InputFiles
			case 4:
				updates.OutputInstances = oldRecipe.OutputInstances
			}
		}
	}
	return updates
}
