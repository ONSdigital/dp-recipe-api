package main

import (
	"context"
	"syscall"

	"github.com/ONSdigital/go-ns/server"
	"github.com/ONSdigital/log.go/log"

	"os"
	"os/signal"

	"github.com/ONSdigital/dp-recipe-api/api"
	"github.com/ONSdigital/dp-recipe-api/config"
	"github.com/ONSdigital/dp-recipe-api/mongo"
	"github.com/gorilla/mux"
)

func main() {
	log.Namespace = "recipe-api"
	ctx := context.Background()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	cfg, configErr := config.Get()
	if configErr != nil {
		log.Event(ctx, "error loading app config", log.FATAL, log.Error(configErr))
		os.Exit(1)
	}

	log.Event(ctx, "config on startup", log.INFO, log.Data{"config": cfg})

	//Feature Flag for Mongo Connection
	ENABLE_MONGO_DATA := cfg.MongoConfig.EnableMongoData

	if ENABLE_MONGO_DATA {
		mongodb := &mongo.Mongo{
			Collection: cfg.MongoConfig.Collection,
			Database:   cfg.MongoConfig.Database,
			URI:        cfg.MongoConfig.BindAddr,
		}

		var err error
		mongodb.Session, err = mongodb.Init()
		if err != nil {
			log.Event(ctx, "failed to initialise mongo", log.FATAL, log.Error(err))
			os.Exit(1)
		}
	}

	router := mux.NewRouter()
	router.Methods("GET").Path("/recipes").HandlerFunc(api.RecipeListHandler)
	router.Methods("GET").Path("/recipes/{id}").HandlerFunc(api.RecipeHandler)

	log.Event(ctx, "starting http server", log.INFO, log.Data{"bind_addr": cfg.BindAddr})
	srv := server.New(cfg.BindAddr, router)
	if err := srv.ListenAndServe(); err != nil {
		log.Event(ctx, "error starting http server for API", log.FATAL, log.Error(err))
		os.Exit(1)
	}
}
