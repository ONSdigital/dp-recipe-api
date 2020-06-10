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

//AddAllRecipeHandler - adds all the recipes from data.go to the mongo database
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

//AddRecipeHandler - add a recipe
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
	err = recipe.ValidateAddRecipe(ctx)
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

//AddInstanceHandler - add an instance to an existing recipe
func (api *RecipeAPI) AddInstanceHandler(w http.ResponseWriter, req *http.Request) {
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

	// Unmarshal to the shape of Instance struct
	var instance recipe.Instance
	err = json.Unmarshal(b, &instance)
	if err != nil {
		log.Event(ctx, "error returned from json unmarshal", log.ERROR, log.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Validation to check if all the instance fields are entered
	err = instance.ValidateAddInstance(ctx)
	if err != nil {
		log.Event(ctx, "bad request error as invalid instance given in request body", log.ERROR, log.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get currentRecipe from ID given to get instance of the recipe
	currentRecipe, err := api.dataStore.Backend.GetRecipe(recipeID)
	if err != nil {
		log.Event(ctx, "error retrieving specific recipe from mongo", log.ERROR, log.Error(err), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Append new instance into the current instance of the recipe
	currentRecipe.OutputInstances = append(currentRecipe.OutputInstances, instance)

	// Add instance to existing recipe in mongo
	err = api.dataStore.Backend.AddInstance(recipeID, currentRecipe)
	if err != nil {
		log.Event(ctx, "error adding instance by updating recipe in mongo", log.ERROR, log.Error(err), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Marshal to output the newly added recipe
	output, err := json.Marshal(currentRecipe)
	if err != nil {
		log.Event(ctx, "error returned from json marshal", log.ERROR, log.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

//AddCodelistHandler - add a codelist in the instance of an existing recipe
func (api *RecipeAPI) AddCodelistHandler(w http.ResponseWriter, req *http.Request) {
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

	// Unmarshal to the shape of Codelist struct
	var codelist recipe.CodeList
	err = json.Unmarshal(b, &codelist)
	if err != nil {
		log.Event(ctx, "error returned from json unmarshal", log.ERROR, log.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Generate the HRef of codelist if not given in request body as it follows a consistent pattern
	if codelist.ID != "" && codelist.HRef == "" {
		codelist.HRef = recipe.HRefURL + codelist.ID
	}

	// Validation to check if all the instance fields are entered
	err = codelist.ValidateAddCodelist(ctx)
	if err != nil {
		log.Event(ctx, "bad request error as invalid codelist given in request body", log.ERROR, log.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get currentRecipe from ID and retrieve specific instance for codelist to be stored
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

	// Append new codelist into the current codelist within the specific instance of the recipe
	currentRecipe.OutputInstances[instanceIndex].CodeLists = append(currentRecipe.OutputInstances[instanceIndex].CodeLists, codelist)

	// Update the current recipe in mongo with the updated codelist in the specific instance
	err = api.dataStore.Backend.AddCodelist(recipeID, instanceIndex, currentRecipe)
	if err != nil {
		log.Event(ctx, "error adding codelist by updating recipe in mongo", log.ERROR, log.Error(err), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Marshal to output the newly added recipe
	output, err := json.Marshal(currentRecipe)
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
	err = updates.ValidateUpdateRecipe(ctx)
	if err != nil {
		log.Event(ctx, "bad request error for invalid updates given in request body", log.ERROR, log.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Update Recipe to Mongo
	err = api.dataStore.Backend.UpdateRecipe(recipeID, updates)
	if err != nil {
		log.Event(ctx, "error updating recipe to mongo", log.ERROR, log.Error(err), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
}

func findInstance(instanceID string, instances []recipe.Instance) int {
	defaultIndex := -1
	for i, instance := range instances {
		if instance.DatasetID == instanceID {
			return i
		}
	}
	return defaultIndex
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

	// Validation to check fields of instance and if all the codelists of the instance are entered if update of codelist given
	err = updates.ValidateUpdateInstance(ctx)
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

	instanceIndex := findInstance(instanceID, currentRecipe.OutputInstances)
	if instanceIndex == -1 {
		log.Event(ctx, "error retrieving specific instance of recipe from mongo", log.ERROR, logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the non-updated fields to values of the currentRecipe - this needs to be done otherwise all fields in array will be overwritten
	updates.DatasetID = instanceID
	if updates.Editions == nil || len(updates.Editions) == 0 {
		updates.Editions = currentRecipe.OutputInstances[instanceIndex].Editions
	}
	if updates.Title == "" {
		updates.Title = currentRecipe.OutputInstances[instanceIndex].Title
	}
	if updates.CodeLists == nil || len(updates.CodeLists) == 0 {
		updates.CodeLists = currentRecipe.OutputInstances[instanceIndex].CodeLists
	}

	// Update Recipe to Mongo
	err = api.dataStore.Backend.UpdateInstance(recipeID, instanceIndex, updates)
	if err != nil {
		log.Event(ctx, "error updating specific instance of recipe in mongo", log.ERROR, log.Error(err), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
}

func findCodelist(codelistID string, codelists []recipe.CodeList) int {
	defaultIndex := -1
	for i, codelist := range codelists {
		if codelist.ID == codelistID {
			return i
		}
	}
	return defaultIndex
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

	// Validating fields of codelist given in request body
	err = updates.ValidateUpdateCodeList()
	if err != nil {
		log.Event(ctx, "bad request error for invalid updates given in request body", log.ERROR, log.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Find index of specific codelist and interface of codelist in output_instances of the recipe
	currentRecipe, err := api.dataStore.Backend.GetRecipe(recipeID)
	if err != nil {
		log.Event(ctx, "error retrieving specific recipe from mongo", log.ERROR, log.Error(err), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	instanceIndex := findInstance(instanceID, currentRecipe.OutputInstances)
	if instanceIndex == -1 {
		log.Event(ctx, "error retrieving specific instance of recipe from mongo", log.ERROR, logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	codelistIndex := findCodelist(codelistID, currentRecipe.OutputInstances[instanceIndex].CodeLists)
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
	err = api.dataStore.Backend.UpdateCodelist(recipeID, instanceIndex, codelistIndex, updates)
	if err != nil {
		log.Event(ctx, "error updating codelist to recipe in mongo", log.ERROR, log.Error(err), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
}
