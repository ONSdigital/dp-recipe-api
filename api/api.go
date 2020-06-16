package api

import (
	"context"
	"net/http"

	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/ONSdigital/dp-recipe-api/config"
	"github.com/ONSdigital/dp-recipe-api/store"
	"github.com/ONSdigital/go-ns/server"
	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

var srv *server.Server

//RecipeAPI contains store and features for managing the recipe
type RecipeAPI struct {
	dataStore         store.DataStore
	Router            *mux.Router
	EnableMongoImport bool
	EnableAuthImport  bool
}

//CreateAndInitialiseRecipeAPI create a new RecipeAPI instance based on the configuration provided and starts the HTTP server.
func CreateAndInitialiseRecipeAPI(ctx context.Context, cfg config.Configuration, dataStore store.DataStore, hc *healthcheck.HealthCheck, errorChan chan error) {
	router := mux.NewRouter()
	api := NewRecipeAPI(ctx, cfg, router, dataStore)

	healthcheckHandler := newMiddleware(hc.Handler)
	middleware := alice.New(healthcheckHandler)

	srv = server.New(cfg.BindAddr, middleware.Then(api.Router))

	// Disable this here to allow main to manage graceful shutdown of the entire app.
	srv.HandleOSSignals = false

	go func() {
		log.Event(ctx, "starting http server", log.INFO, log.Data{"bind_addr": cfg.BindAddr})
		if err := srv.ListenAndServe(); err != nil {
			log.Event(ctx, "error starting http server for API", log.FATAL, log.Error(err))
			errorChan <- err
		}
	}()
}

//NewRecipeAPI create a new Recipe API instance and register the API routes based on the application configuration.
func NewRecipeAPI(ctx context.Context, cfg config.Configuration, router *mux.Router, dataStore store.DataStore) *RecipeAPI {
	api := &RecipeAPI{
		dataStore:         dataStore,
		Router:            router,
		EnableMongoImport: cfg.MongoConfig.EnableMongoImport,
		EnableAuthImport:  cfg.MongoConfig.EnableAuthImport,
	}

	api.get("/health", api.HealthCheck)
	api.get("/recipes", api.RecipeListHandler)
	api.get("/recipes/{id}", api.RecipeHandler)
	if api.EnableMongoImport {
		api.post("/allrecipes", api.AddAllRecipeHandler)
	}
	if api.EnableAuthImport {
		api.post("/recipes", api.AddRecipeHandler)
		api.post("/recipes/{id}/instances", api.AddInstanceHandler)
		api.post("/recipes/{id}/instances/{instance_id}/codelists", api.AddCodelistHandler)
		api.put("/recipes/{id}", api.UpdateRecipeHandler)
		api.put("/recipes/{id}/instances/{instance_id}", api.UpdateInstanceHandler)
		api.put("/recipes/{id}/instances/{instance_id}/codelists/{codelist_id}", api.UpdateCodelistHandler)
	}
	return api
}

//get - register a GET http.HandlerFunc.
func (api *RecipeAPI) get(path string, handler http.HandlerFunc) {
	api.Router.HandleFunc(path, handler).Methods("GET")
}

//post - register a POST http.HandlerFunc.
func (api *RecipeAPI) post(path string, handler http.HandlerFunc) {
	api.Router.HandleFunc(path, handler).Methods("POST")
}

//put - register a PUT http.HandlerFunc.
func (api *RecipeAPI) put(path string, handler http.HandlerFunc) {
	api.Router.HandleFunc(path, handler).Methods("PUT")
}

//Close represents the graceful shutting down of the http server
func Close(ctx context.Context) error {
	if err := srv.Shutdown(ctx); err != nil {
		return err
	}
	log.Event(ctx, "graceful shutdown of http server complete", log.INFO)
	return nil
}

//newMiddleware creates a new http.Handler to intercept /health requests.
func newMiddleware(healthcheckHandler func(http.ResponseWriter, *http.Request)) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			if req.Method == "GET" && req.URL.Path == "/health" {
				healthcheckHandler(w, req)
				return
			}
			h.ServeHTTP(w, req)
		})
	}
}
