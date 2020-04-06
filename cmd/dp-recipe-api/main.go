package main

import (
	"context"
	"fmt"
	"syscall"

	"github.com/ONSdigital/log.go/log"

	"os"
	"os/signal"

	health "github.com/ONSdigital/dp-healthcheck/healthcheck"
	mongolib "github.com/ONSdigital/dp-mongodb"
	mongoHealth "github.com/ONSdigital/dp-mongodb/health"
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

//health check variables - app version informaton retrieved on runtime
var (
	// BuildTime represents the time in which the service was built
	BuildTime string
	// GitCommit represents the commit (SHA-1) hash of the service that is running
	GitCommit string
	// Version represents the version of the service that is running
	Version string
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

	versionInfo, healthErr := health.NewVersionInfo(
		BuildTime,
		GitCommit,
		Version,
	)
	if healthErr != nil {
		log.Event(ctx, "failed to create service version information", log.ERROR, log.Error(healthErr))
		os.Exit(1)
	}

	//Feature Flag for Mongo Connection
	enableMongoData := cfg.MongoConfig.EnableMongoData

	mongodb := &mongo.Mongo{
		Collection: cfg.MongoConfig.Collection,
		Database:   cfg.MongoConfig.Database,
		URI:        cfg.MongoConfig.BindAddr,
	}

	session, err := mongodb.Init()
	if err != nil {
		log.Event(ctx, "failed to initialise mongo", log.FATAL, log.Error(err))
		os.Exit(1)
	} else {
		mongodb.Session = session
		log.Event(ctx, "listening to mongo db session", log.INFO, log.Data{
			"bind_address": cfg.BindAddr,
		})
	}

	healthcheck := health.New(versionInfo, cfg.HealthCheckCriticalTimeout, cfg.HealthCheckInterval)

	mongoClient := mongoHealth.NewClient(mongodb.Session)

	mongoHealth := mongoHealth.CheckMongoClient{
		Client:      *mongoClient,
		Healthcheck: mongoClient.Healthcheck,
	}
	if err = healthcheck.AddCheck("mongoDB", mongoHealth.Checker); err != nil {
		log.Event(ctx, "failed to add mongoDB checker", log.ERROR, log.Error(err))
		os.Exit(1)
	}

	if enableMongoData {
		//Create RecipeAPI instance with Mongo in datastore
		store := store.DataStore{Backend: RecipeAPIStore{mongodb}}
		api.CreateAndInitialiseRecipeAPI(ctx, *cfg, store, &healthcheck)

	} else {
		//Create RecipeAPI instance with no datastore
		api.CreateAndInitialiseRecipeAPI(ctx, *cfg, store.DataStore{Backend: nil}, &healthcheck)
	}

	healthcheck.Start(ctx)

	apiErrors := make(chan error, 1)

	// Gracefully shutdown the application closing any open resources.
	gracefulShutdown := func() {
		log.Event(ctx, (fmt.Sprintf("shutdown with timeout: %s", cfg.GracefulShutdownTimeout)), log.INFO)
		ctx, cancel := context.WithTimeout(context.Background(), cfg.GracefulShutdownTimeout)

		healthcheck.Stop()

		// stop any incoming requests before closing any outbound connections
		api.Close(ctx)

		if enableMongoData {
			if err = mongolib.Close(ctx, mongodb.Session); err != nil {
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
