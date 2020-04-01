package api

import (
	"context"
	"net/http"
	"os"

	"github.com/ONSdigital/dp-recipe-api/config"
	"github.com/ONSdigital/dp-recipe-api/store"
	"github.com/ONSdigital/go-ns/server"
	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/mux"
)

//RecipeAPI contains store and features for managing the recipe
type RecipeAPI struct {
	dataStore         store.DataStore
	Router            *mux.Router
	EnableMongoData   bool
	EnableMongoImport bool
}

//CreateAndInitialiseRecipeAPI create a new RecipeAPI instance based on the configuration provided and starts the HTTP server.
func CreateAndInitialiseRecipeAPI(ctx context.Context, cfg config.Configuration, dataStore store.DataStore) {
	router := mux.NewRouter()
	api := NewRecipeAPI(ctx, cfg, router, dataStore)

	log.Event(ctx, "starting http server", log.INFO, log.Data{"bind_addr": cfg.BindAddr})
	srv := server.New(cfg.BindAddr, api.Router)
	if err := srv.ListenAndServe(); err != nil {
		log.Event(ctx, "error starting http server for API", log.FATAL, log.Error(err))
		os.Exit(1)
	}
}

//NewRecipeAPI create a new Recipe API instance and register the API routes based on the application configuration.
func NewRecipeAPI(ctx context.Context, cfg config.Configuration, router *mux.Router, dataStore store.DataStore) *RecipeAPI {
	api := &RecipeAPI{
		dataStore:         dataStore,
		Router:            router,
		EnableMongoData:   cfg.MongoConfig.EnableMongoData,
		EnableMongoImport: cfg.MongoConfig.EnableMongoImport,
	}

	api.get("/health", api.HealthCheck)
	api.get("/recipes", api.RecipeListHandler)
	api.get("/recipes/{id}", api.RecipeHandler)
	return api
}

// get register a GET http.HandlerFunc.
func (api *RecipeAPI) get(path string, handler http.HandlerFunc) {
	api.Router.HandleFunc(path, handler).Methods("GET")
}
