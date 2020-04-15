package main

import (
	"context"
	"fmt"
	"syscall"

	"github.com/ONSdigital/log.go/log"

	"os"
	"os/signal"

	"github.com/ONSdigital/dp-recipe-api/api"
	"github.com/ONSdigital/dp-recipe-api/config"
	"github.com/ONSdigital/dp-recipe-api/mongo"
	"github.com/ONSdigital/dp-recipe-api/store"
	mongolib "github.com/ONSdigital/dp-mongodb"
)

//check that RecipeAPIStore satifies the the store.Storer interface
var _ store.Storer = (*RecipeAPIStore)(nil)

//RecipeAPIStore is a wrapper which embeds Mongo struct which between them satisfy the store.Storer interface.
type RecipeAPIStore struct {
	*mongo.Mongo
}

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
	enableMongoData := cfg.MongoConfig.EnableMongoData

	mongodb := &mongo.Mongo{
		Collection: cfg.MongoConfig.Collection,
		Database:   cfg.MongoConfig.Database,
		URI:        cfg.MongoConfig.BindAddr,
	}

	if enableMongoData {
		session, err := mongodb.Init()
		if err != nil {
			log.Event(ctx, "failed to initialise mongo", log.FATAL, log.Error(err))
			os.Exit(1)
		} else {
			mongodb.Session = session
		}

		//Create RecipeAPI instance with Mongo in datastore
		store := store.DataStore{Backend: RecipeAPIStore{mongodb}}
		api.CreateAndInitialiseRecipeAPI(ctx, *cfg, store)

	} else {
		//Create RecipeAPI instance with no datastore
		api.CreateAndInitialiseRecipeAPI(ctx, *cfg, store.DataStore{Backend: nil})
	}

	apiErrors := make(chan error, 1)

	// Gracefully shutdown the application closing any open resources.
	gracefulShutdown := func() {
		log.Event(ctx, (fmt.Sprintf("shutdown with timeout: %s", cfg.GracefulShutdownTimeout)), log.INFO)
		ctx, cancel := context.WithTimeout(context.Background(), cfg.GracefulShutdownTimeout)

		// stop any incoming requests before closing any outbound connections
		api.Close(ctx)

		if enableMongoData {

			if err := mongolib.Close(ctx, mongodb.Session); err != nil {
				log.Event(ctx, "failed to close mongo session", log.ERROR, log.Error(err))
			}
		}

		log.Event(ctx, "shutdown complete", log.INFO)

		cancel()
		os.Exit(1)
	}

	for {
		select {
		case err := <-apiErrors:
			log.Event(ctx, "api error received", log.ERROR, log.Error(err))
		case <-signals:
			log.Event(ctx, "os signal received", log.INFO)
			gracefulShutdown()
		}
	}
}
