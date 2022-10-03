package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	errs "github.com/ONSdigital/dp-recipe-api/apierrors"
	"github.com/ONSdigital/dp-recipe-api/models"
	"github.com/ONSdigital/dp-recipe-api/utils"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// RecipeListHandler - get all recipes
func (api *RecipeAPI) RecipeListHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	logData := log.Data{}

	offsetParameter := req.URL.Query().Get("offset")
	limitParameter := req.URL.Query().Get("limit")

	limit := api.defaultLimit
	offset := api.defaultOffset

	var err error

	if offsetParameter != "" {
		logData["offset"] = offsetParameter
		offset, err = utils.ValidatePositiveInt(offsetParameter)
		if err != nil {
			log.Error(ctx, "invalid query parameter: offset", err, logData)
			handleErr(ctx, w, err, nil)
			return
		}
	}

	if limitParameter != "" {
		logData["limit"] = limitParameter
		limit, err = utils.ValidatePositiveInt(limitParameter)
		if err != nil {
			log.Error(ctx, "invalid query parameter: limit", err, logData)
			handleErr(ctx, w, err, nil)
			return
		}
	}

	if limit > api.maxLimit {
		logData["max_limit"] = api.maxLimit
		log.Error(ctx, "limit is greater than the maximum allowed", errors.New("limit is greater than the maximum allowed"), logData)
		handleCustomErr(ctx, w, errs.ErrorMaximumLimitReached(api.maxLimit), logData, http.StatusBadRequest)
		return
	}

	var recipeResults *models.RecipeResults

	recipeResults, err = api.dataStore.Backend.GetRecipes(ctx, offset, limit)
	if err != nil {
		log.Error(ctx, "getRecipes endpoint: failed to retrieve a list of recipes from mongo", err, logData)
		handleErr(ctx, w, err, nil)
		return
	}

	b, err := json.Marshal(recipeResults)
	if err != nil {
		log.Error(ctx, "getRecipes endpoint: failed to marshal recipes resource into bytes", err, logData)
		handleErr(ctx, w, err, nil)
		return
	}

	writeResponse(ctx, w, http.StatusOK, b, "getRecipes", logData)
	log.Info(ctx, "getRecipes endpoint: request successful", logData)
}

// RecipeHandler - get recipe by ID
func (api *RecipeAPI) RecipeHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)
	recipeID := vars["id"]
	logD := log.Data{"recipe_id": recipeID}

	var r models.Recipe
	recipe, err := api.dataStore.Backend.GetRecipe(ctx, recipeID)
	if err == errs.ErrRecipeNotFound {
		log.Error(ctx, "recipe not found in mongo", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		log.Error(ctx, "error getting recipe from mongo", err, logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	r = *recipe

	b, err := json.Marshal(r)
	if err != nil {
		log.Error(ctx, "error returned from json marshal", err, logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

// HealthCheck - health check endpoint
func (api *RecipeAPI) HealthCheck(w http.ResponseWriter, req *http.Request) {
	// Set status to 200 OK
	w.WriteHeader(200)
}

// AddRecipeHandler - add a recipe
func (api *RecipeAPI) AddRecipeHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	// Read body
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.Error(ctx, "error in reading request body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Unmarshal to the shape of Recipe struct
	var recipe models.Recipe
	err = json.Unmarshal(b, &recipe)
	if err != nil {
		log.Error(ctx, "error returned from json unmarshal", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Randomly generated version 4 UUID for recipe ID
	if recipe.ID == "" {
		recipe.ID = uuid.UUID.String(uuid.New())
	}

	// Validation to check if all the recipe fields are entered
	err = recipe.ValidateAddRecipe(ctx)
	if err != nil {
		log.Error(ctx, "bad request error as incomplete recipe given in request body", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add Recipe to Mongo
	err = api.dataStore.Backend.AddRecipe(ctx, recipe)
	if err != nil {
		log.Error(ctx, "error adding recipe to mongo", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Marshal to output the newly added recipe
	output, err := json.Marshal(recipe)
	if err != nil {
		log.Error(ctx, "error returned from json marshal", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

// AddInstanceHandler - add an instance to an existing recipe
func (api *RecipeAPI) AddInstanceHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	// Get Update Recipe ID
	vars := mux.Vars(req)
	recipeID := vars["id"]
	logD := log.Data{"recipe_id": recipeID}

	// Read body
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.Error(ctx, "error in reading request body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Unmarshal to the shape of Instance struct
	var instance models.Instance
	err = json.Unmarshal(b, &instance)
	if err != nil {
		log.Error(ctx, "error returned from json unmarshal", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get currentRecipe from ID given to get instance of the recipe
	currentRecipe, err := api.dataStore.Backend.GetRecipe(ctx, recipeID)
	if err != nil {
		log.Error(ctx, "error retrieving specific recipe from mongo", err, logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Validation to check if all the instance fields are entered
	err = instance.ValidateAddInstance(ctx, currentRecipe)
	if err != nil {
		log.Error(ctx, "bad request error as invalid instance given in request body", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Append new instance into the current instance of the recipe
	currentRecipe.OutputInstances = append(currentRecipe.OutputInstances, instance)

	// Add instance to existing recipe in mongo
	err = api.dataStore.Backend.UpdateRecipe(ctx, recipeID, *currentRecipe)
	if err != nil {
		log.Error(ctx, "error adding instance by updating recipe in mongo", err, logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Marshal to output the newly added recipe
	output, err := json.Marshal(currentRecipe)
	if err != nil {
		log.Error(ctx, "error returned from json marshal", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

// AddCodelistHandler - add a codelist in the instance of an existing recipe
func (api *RecipeAPI) AddCodelistHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	// Get Update Recipe ID
	vars := mux.Vars(req)
	recipeID := vars["id"]
	instanceID := vars["dataset_id"]
	logD := log.Data{"recipe_id": recipeID, "dataset_id": instanceID}

	// Read body
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.Error(ctx, "error in reading request body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Unmarshal to the shape of Codelist struct
	var codelist models.CodeList
	err = json.Unmarshal(b, &codelist)
	if err != nil {
		log.Error(ctx, "error returned from json unmarshal", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Generate the HRef of codelist if not given in request body as it follows a consistent pattern
	if codelist.ID != "" && codelist.HRef == "" {
		log.Info(ctx, "href automatically updated with new id")
		codelist.HRef = models.HRefURL + codelist.ID
	}

	// Get currentRecipe from ID and retrieve specific instance for codelist to be stored
	currentRecipe, err := api.dataStore.Backend.GetRecipe(ctx, recipeID)
	if err != nil {
		log.Error(ctx, "error retrieving specific recipe from mongo", err, logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Validation to check if all the instance fields are entered
	err = codelist.ValidateAddCodelist(ctx, currentRecipe)
	if err != nil {
		log.Error(ctx, "bad request error as invalid codelist given in request body", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	instanceIndex := findInstance(instanceID, currentRecipe.OutputInstances)
	if instanceIndex == -1 {
		log.Error(ctx, "error retrieving specific instance of recipe from mongo", errors.New("error retrieving specific instance of recipe from mongo"), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Append new codelist into the current codelist within the specific instance of the recipe
	currentRecipe.OutputInstances[instanceIndex].CodeLists = append(currentRecipe.OutputInstances[instanceIndex].CodeLists, codelist)

	// Update the current recipe in mongo with the updated codelist in the specific instance
	err = api.dataStore.Backend.AddCodelist(ctx, recipeID, instanceIndex, currentRecipe)
	if err != nil {
		log.Error(ctx, "error adding codelist by updating recipe in mongo", err, logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Marshal to output the newly added recipe
	output, err := json.Marshal(currentRecipe)
	if err != nil {
		log.Error(ctx, "error returned from json marshal", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

// UpdateRecipeHandler - update specific recipe by ID
func (api *RecipeAPI) UpdateRecipeHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	// Get Update Recipe ID
	vars := mux.Vars(req)
	recipeID := vars["id"]
	logD := log.Data{"recipe_id": recipeID}

	// Read body
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.Error(ctx, "error in reading request body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Unmarshal to the shape of Recipe struct
	var updates models.Recipe
	err = json.Unmarshal(b, &updates)
	if err != nil {
		log.Error(ctx, "error returned from json unmarshal", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Validation to check if all the recipe fields are entered
	err = updates.ValidateUpdateRecipe(ctx)
	if err != nil {
		log.Error(ctx, "bad request error for invalid updates given in request body", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update Recipe to Mongo
	err = api.dataStore.Backend.UpdateRecipe(ctx, recipeID, updates)
	if err != nil {
		log.Error(ctx, "error updating recipe to mongo", err, logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
}

// UpdateInstanceHandler - update specific recipe by ID
func (api *RecipeAPI) UpdateInstanceHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	// Get Update Recipe ID
	vars := mux.Vars(req)
	recipeID := vars["id"]
	instanceID := vars["dataset_id"]
	logD := log.Data{"recipe_id": recipeID, "dataset_id": instanceID}

	// Read body
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.Error(ctx, "error in reading request body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Unmarshal to the shape of Instance struct
	var updates models.Instance
	err = json.Unmarshal(b, &updates)
	if err != nil {
		log.Error(ctx, "error returned from json unmarshal", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Find index of specific interface in output_instances of the recipe
	currentRecipe, err := api.dataStore.Backend.GetRecipe(ctx, recipeID)
	if err != nil {
		log.Error(ctx, "error retrieving specific recipe from mongo", err, logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Validation to check fields of instance and if all the code lists of the instance are entered if update of codelist given
	err = updates.ValidateUpdateInstance(ctx, currentRecipe)
	if err != nil {
		log.Error(ctx, "bad request error for invalid updates given in request body", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	instanceIndex := findInstance(instanceID, currentRecipe.OutputInstances)
	if instanceIndex == -1 {
		log.Error(ctx, "error retrieving specific instance of recipe from mongo", errors.New("error retrieving specific instance of recipe from mongo"), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the non-updated fields to values of the currentRecipe - this needs to be done otherwise all fields in array will be overwritten
	currentInstance := currentRecipe.OutputInstances[instanceIndex]
	updates = setInstance(instanceID, currentInstance, updates)

	// Update Recipe to Mongo
	err = api.dataStore.Backend.UpdateInstance(ctx, recipeID, instanceIndex, updates)
	if err != nil {
		log.Error(ctx, "error updating specific instance of recipe in mongo", err, logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
}

// UpdateCodelistHandler - update specific recipe by ID
func (api *RecipeAPI) UpdateCodelistHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	// Get Update Recipe ID
	vars := mux.Vars(req)
	recipeID := vars["id"]
	instanceID := vars["dataset_id"]
	codelistID := vars["code_list_id"]
	logD := log.Data{"recipe_id": recipeID, "dataset_id": instanceID, "code_list_id": codelistID}

	// Read body
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.Error(ctx, "error in reading request body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Unmarshal to the shape of Recipe struct
	var updates models.CodeList
	err = json.Unmarshal(b, &updates)
	if err != nil {
		log.Error(ctx, "error returned from json unmarshal", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Generate the HRef of codelist if not given in request body as it follows a consistent pattern
	if updates.ID != "" && updates.HRef == "" {
		log.Info(ctx, "href automatically updated with new id")
		updates.HRef = models.HRefURL + updates.ID
	}

	// Validating fields of codelist given in request body
	err = updates.ValidateUpdateCodeList(ctx)
	if err != nil {
		log.Error(ctx, "bad request error for invalid updates given in request body", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Find index of specific codelist and interface of codelist in output_instances of the recipe
	currentRecipe, err := api.dataStore.Backend.GetRecipe(ctx, recipeID)
	if err != nil {
		log.Error(ctx, "error retrieving specific recipe from mongo", err, logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	instanceIndex := findInstance(instanceID, currentRecipe.OutputInstances)
	if instanceIndex == -1 {
		log.Error(ctx, "error retrieving specific instance of recipe from mongo", errors.New("error retrieving specific instance of recipe from mongo"), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	codelistIndex := findCodelist(codelistID, currentRecipe.OutputInstances[instanceIndex].CodeLists)
	if codelistIndex == -1 {
		log.Error(ctx, "error retrieving specific codelist of instance of recipe from mongo", errors.New("error retrieving specific codelist of instance of recipe from mongo"), logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the non-updated fields to values of the currentRecipe
	currentCodelist := currentRecipe.OutputInstances[instanceIndex].CodeLists[codelistIndex]
	updates = setCodelist(codelistID, currentCodelist, updates)

	// Update Recipe to Mongo
	err = api.dataStore.Backend.UpdateCodelist(ctx, recipeID, instanceIndex, codelistIndex, updates)
	if err != nil {
		log.Error(ctx, "error updating codelist to recipe in mongo", err, logD)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
}

func findInstance(instanceID string, instances []models.Instance) int {
	defaultIndex := -1
	for i, instance := range instances {
		if instance.DatasetID == instanceID {
			return i
		}
	}
	return defaultIndex
}

func setInstance(instanceID string, currentInstance models.Instance, updates models.Instance) models.Instance {
	if updates.DatasetID == "" {
		updates.DatasetID = instanceID
	}
	if updates.Editions == nil || len(updates.Editions) == 0 {
		updates.Editions = currentInstance.Editions
	}
	if updates.Title == "" {
		updates.Title = currentInstance.Title
	}
	if updates.CodeLists == nil || len(updates.CodeLists) == 0 {
		updates.CodeLists = currentInstance.CodeLists
	}
	return updates
}

func findCodelist(codelistID string, codelists []models.CodeList) int {
	defaultIndex := -1
	for i, codelist := range codelists {
		if codelist.ID == codelistID {
			return i
		}
	}
	return defaultIndex
}

func setCodelist(codelistID string, currentCodelist models.CodeList, updates models.CodeList) models.CodeList {
	if updates.ID == "" {
		updates.ID = codelistID
	}
	if updates.HRef == "" {
		updates.HRef = currentCodelist.HRef
	}
	if updates.Name == "" {
		updates.Name = currentCodelist.Name
	}
	if updates.IsHierarchy == nil {
		updates.IsHierarchy = currentCodelist.IsHierarchy
	}
	return updates
}
