package api

//go:generate moq -out mock/mockauth.go -pkg mock . AuthHandler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ONSdigital/dp-authorisation/auth"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/ONSdigital/dp-recipe-api/config"
	"github.com/ONSdigital/dp-recipe-api/store"
	"github.com/ONSdigital/go-ns/server"
	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	errs "github.com/ONSdigital/dp-recipe-api/apierrors"
)

var (
	srv *server.Server

	create = auth.Permissions{Create: true}
	update = auth.Permissions{Update: true}
)

// AuthHandler provides authorisation checks on requests
type AuthHandler interface {
	Require(required auth.Permissions, handler http.HandlerFunc) http.HandlerFunc
}

//RecipeAPI contains store and features for managing the recipe
type RecipeAPI struct {
	dataStore     store.DataStore
	Router        *mux.Router
	permissions   AuthHandler
	defaultLimit  int
	defaultOffset int
	maxLimit      int
}

//CreateAndInitialiseRecipeAPI create a new RecipeAPI instance based on the configuration provided and starts the HTTP server.
func CreateAndInitialiseRecipeAPI(ctx context.Context, cfg config.Configuration, dataStore store.DataStore, hc *healthcheck.HealthCheck, errorChan chan error, permissions AuthHandler) {
	router := mux.NewRouter()
	api := NewRecipeAPI(ctx, cfg, router, dataStore, permissions)

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
func NewRecipeAPI(ctx context.Context, cfg config.Configuration, router *mux.Router, dataStore store.DataStore, permissions AuthHandler) *RecipeAPI {
	api := &RecipeAPI{
		dataStore:     dataStore,
		Router:        router,
		permissions:   permissions,
		defaultLimit:  cfg.DefaultLimit,
		defaultOffset: cfg.DefaultOffset,
		maxLimit:      cfg.DefaultMaxLimit,
	}

	api.get("/health", api.HealthCheck)
	api.get("/recipes", api.RecipeListHandler)
	api.get("/recipes/{id}", api.RecipeHandler)
	api.post("/recipes", permissions.Require(create, api.AddRecipeHandler))
	api.post("/recipes/{id}/instances", permissions.Require(create, api.AddInstanceHandler))
	api.post("/recipes/{id}/instances/{dataset_id}/code-lists", permissions.Require(create, api.AddCodelistHandler))
	api.put("/recipes/{id}", permissions.Require(update, api.UpdateRecipeHandler))
	api.put("/recipes/{id}/instances/{dataset_id}", permissions.Require(update, api.UpdateInstanceHandler))
	api.put("/recipes/{id}/instances/{dataset_id}/code-lists/{code_list_id}", permissions.Require(update, api.UpdateCodelistHandler))
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


func writeResponse(ctx context.Context, w http.ResponseWriter, statusCode int, b []byte, action string, logData log.Data) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if _, err := w.Write(b); err != nil {
		log.Event(ctx, fmt.Sprintf("%s endpoint: failed to write response body", action), log.ERROR, log.Error(err), logData)
	}
}

func handleErr(ctx context.Context, w http.ResponseWriter, err error, logData log.Data) {
	if logData == nil {
		logData = log.Data{}
	}

	var status int
	response := err

	switch {
	case errs.NotFoundMap[err]:
		status = http.StatusNotFound
	case errs.BadRequestMap[err]:
		status = http.StatusBadRequest
	default:
		status = http.StatusInternalServerError
		response = errs.ErrInternalServer
	}

	logResponseStatus(ctx, logData, status, err)
	http.Error(w, response.Error(), status)
}

func handleCustomErr(ctx context.Context, w http.ResponseWriter, err error, logData log.Data, status int) {
	if logData == nil {
		logData = log.Data{}
	}
	logResponseStatus(ctx, logData, status, err)
	http.Error(w, err.Error(), status)
}

func logResponseStatus(ctx context.Context, logData log.Data, status int, err error) {
	logData["responseStatus"] = status
	log.Event(ctx, "request unsuccessful", log.ERROR, log.Error(err), logData)
}