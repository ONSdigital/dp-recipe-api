package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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
	if recipe.ID == "" {
		recipe.ID = uuid.UUID.String(uuid.New())
	}

	//Validation to check if all the recipe fields are entered
	err = recipe.ValidateAddRecipe()
	if err != nil {
		log.Event(ctx, "bad request error as incomplete recipe given in request body", log.ERROR, log.Error(err))
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
	var updates recipe.Response
	err = json.Unmarshal(b, &updates)
	if err != nil {
		log.Event(ctx, "error returned from json unmarshal", log.ERROR, log.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Validation to check if all the recipe fields are entered
	err = updates.ValidateUpdateRecipe("recipe")
	if err != nil {
		log.Event(ctx, "bad request error for invalid updates given in request body", log.ERROR, log.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Update Recipe to Mongo
	err = api.dataStore.Backend.UpdateRecipe(recipeID, updates, 0, 0)
	if err != nil {
		log.Event(ctx, "error updating recipe to mongo", log.ERROR, log.Error(err), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Marshal to output the newly updated recipe
	updatedRecipe, _ := api.dataStore.Backend.GetRecipe(recipeID)
	output, err := json.Marshal(updatedRecipe)
	if err != nil {
		log.Event(ctx, "error marshaling updated recipe from mongo", log.ERROR, log.Error(err), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

//UpdateInstanceHandler - update specific recipe by ID
func (api *RecipeAPI) UpdateInstanceHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	//Get Update Recipe ID
	vars := mux.Vars(req)
	recipeID := vars["id"]
	instanceID := vars["instance_id"]
	logD := log.Data{"recipe_id": recipeID, "instance_id": instanceID}

	// Read body
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.Event(ctx, "error in reading request body", log.ERROR, log.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Unmarshal to the shape of Instance struct
	var updates recipe.Instance
	err = json.Unmarshal(b, &updates)
	if err != nil {
		log.Event(ctx, "error returned from json unmarshal", log.ERROR, log.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Validation to check if all the codelists of the instance are entered if update of codelist given
	var missingFields []string
	var invalidFields []string
	err = updates.ValidateInstance("instance", missingFields, invalidFields, 0)
	if err != nil {
		log.Event(ctx, "bad request error for invalid updates given in request body", log.ERROR, log.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Find index of specific interface in output_instances of the recipe
	currentRecipe, err := api.dataStore.Backend.GetRecipe(recipeID)
	if err != nil {
		log.Event(ctx, "error retrieving specific recipe from mongo", log.ERROR, log.Error(err), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	instanceIndex := -1
	for i, instance := range currentRecipe.OutputInstances {
		if instance.DatasetID == instanceID {
			instanceIndex = i
			break
		}
	}
	if instanceIndex == -1 {
		log.Event(ctx, "error retrieving specific instance of recipe from mongo", log.ERROR, logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the non-updated fields to values of the currentRecipe
	updates.DatasetID = instanceID
	if updates.Editions == nil {
		updates.Editions = currentRecipe.OutputInstances[instanceIndex].Editions
	}
	if updates.Title == "" {
		updates.Title = currentRecipe.OutputInstances[instanceIndex].Title
	}
	if updates.CodeLists == nil {
		updates.CodeLists = currentRecipe.OutputInstances[instanceIndex].CodeLists
	}

	// Update Recipe to Mongo
	err = api.dataStore.Backend.UpdateRecipe(recipeID, updates, instanceIndex, 0)
	if err != nil {
		log.Event(ctx, "error updating recipe to mongo", log.ERROR, log.Error(err), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Marshal to output the newly updated recipe
	updatedRecipe, _ := api.dataStore.Backend.GetRecipe(recipeID)
	output, err := json.Marshal(updatedRecipe)
	if err != nil {
		log.Event(ctx, "error marshaling updated recipe from mongo", log.ERROR, log.Error(err), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

//UpdateCodelistHandler - update specific recipe by ID
func (api *RecipeAPI) UpdateCodelistHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	//Get Update Recipe ID
	vars := mux.Vars(req)
	recipeID := vars["id"]
	instanceID := vars["instance_id"]
	codelistID := vars["codelist_id"]
	logD := log.Data{"recipe_id": recipeID, "instance_id": instanceID, "codelist_id": codelistID}

	// Read body
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.Event(ctx, "error in reading request body", log.ERROR, log.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Unmarshal to the shape of Response struct
	var updates recipe.CodeList
	err = json.Unmarshal(b, &updates)
	if err != nil {
		log.Event(ctx, "error returned from json unmarshal", log.ERROR, log.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Find index of specific codelist and interface of codelist in output_instances of the recipe
	currentRecipe, err := api.dataStore.Backend.GetRecipe(recipeID)
	if err != nil {
		log.Event(ctx, "error retrieving specific recipe from mongo", log.ERROR, log.Error(err), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	instanceIndex := -1
	codelistIndex := -1
	for i, instance := range currentRecipe.OutputInstances {
		if instance.DatasetID == instanceID {
			instanceIndex = i

			for j, codelist := range instance.CodeLists {
				if codelist.ID == codelistID {
					codelistIndex = j
					break
				}
			}

			break
		}
	}

	if instanceIndex == -1 {
		log.Event(ctx, "error retrieving specific instance of recipe from mongo", log.ERROR, logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if codelistIndex == -1 {
		log.Event(ctx, "error retrieving specific codelist of instance of recipe from mongo", log.ERROR, logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the non-updated fields to values of the currentRecipe
	updates.ID = codelistID
	if updates.HRef == "" {
		updates.HRef = currentRecipe.OutputInstances[instanceIndex].CodeLists[codelistIndex].HRef
	}
	if updates.Name == "" {
		updates.Name = currentRecipe.OutputInstances[instanceIndex].CodeLists[codelistIndex].Name
	}
	if updates.IsHierarchy == nil {
		updates.IsHierarchy = currentRecipe.OutputInstances[instanceIndex].CodeLists[codelistIndex].IsHierarchy
	}

	// Update Recipe to Mongo
	err = api.dataStore.Backend.UpdateRecipe(recipeID, updates, instanceIndex, codelistIndex)
	if err != nil {
		log.Event(ctx, "error updating recipe to mongo", log.ERROR, log.Error(err), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Marshal to output the newly updated recipe
	updatedRecipe, _ := api.dataStore.Backend.GetRecipe(recipeID)
	output, err := json.Marshal(updatedRecipe)
	if err != nil {
		log.Event(ctx, "error marshaling updated recipe from mongo", log.ERROR, log.Error(err), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
