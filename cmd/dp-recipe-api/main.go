package main

import (
	"context"
	"syscall"

	"github.com/ONSdigital/log.go/log"

	"os"
	"os/signal"

	"github.com/ONSdigital/dp-recipe-api/api"
	"github.com/ONSdigital/dp-recipe-api/config"
	"github.com/ONSdigital/dp-recipe-api/mongo"
	"github.com/ONSdigital/dp-recipe-api/store"
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

		//Create RecipeAPI instance with Mongo in datastore
		store := store.DataStore{Backend: RecipeAPIStore{mongodb}}
		api.CreateAndInitialiseRecipeAPI(ctx, *cfg, store)

	} else {
		//Create RecipeAPI instance with no datastore
		api.CreateAndInitialiseRecipeAPI(ctx, *cfg, store.DataStore{Backend: nil})
	}

}
