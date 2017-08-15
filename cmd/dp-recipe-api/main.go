package main

import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/go-ns/log"
	"github.com/ONSdigital/go-ns/server"

	"os"

	"github.com/ONSdigital/dp-recipe-api/config"
	"github.com/gorilla/mux"
)

func main() {
	log.Namespace = "dp-recipe-api"
	configuration, configErr := config.Get()
	if configErr != nil {
		log.Error(configErr, nil)
		os.Exit(1)
	}

	router := mux.NewRouter()
	router.Path("/recipes").HandlerFunc(recipeListHandler)
	router.Path("/recipes/{recipe_id}").HandlerFunc(recipeHandler)

	log.Debug("starting http server", log.Data{"bind_addr": configuration.BindAddr})
	srv := server.New(configuration.BindAddr, router)
	if err := srv.ListenAndServe(); err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}
}

type recipeListResponse struct {
	Count        int              `json:"count"`
	Start        int              `json:"start_index"`
	ItemsPerPage int              `json:"items_per_page"`
	Items        []recipeResponse `json:"items"`
	TotalCount   int              `json:"total_count"`
}

type recipeResponse struct {
	ID              string     `json:"id"`
	Alias           string     `json:"alias"`
	Format          string     `json:"format"`
	InputFiles      []file     `json:"files"`
	OutputInstances []instance `json:"output_instances"`
}

type instance struct {
	DatasetID string   `json:"dataset_id"`
	Editions  []string `json:"editions"`
}

type file struct {
	Description string `json:"description"`
}

var cpiRecipeList = recipeListResponse{
	Items:        []recipeResponse{cpiRecipe},
	Count:        1,
	TotalCount:   1,
	ItemsPerPage: 10,
	Start:        0,
}

var cpiRecipe = recipeResponse{
	ID:     "b944be78-f56d-409b-9ebd-ab2b77ffe187",
	Alias:  "CPI COICOP",
	Format: "v4",
	InputFiles: []file{
		{"CPI COICOP v4"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "931a8a2a-0dc8-42b6-a884-7b6054ed3b68",
			Editions:  []string{"Time-series"},
		},
	},
}

func recipeListHandler(w http.ResponseWriter, req *http.Request) {
	b, err := json.Marshal(&cpiRecipeList)
	if err != nil {
		log.ErrorR(req, err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func recipeHandler(w http.ResponseWriter, req *http.Request) {
	b, err := json.Marshal(&cpiRecipe)
	if err != nil {
		log.ErrorR(req, err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
