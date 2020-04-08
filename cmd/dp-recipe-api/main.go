package main

import (
	"context"
	"errors"
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

	hc := health.New(versionInfo, cfg.HealthCheckCriticalTimeout, cfg.HealthCheckInterval)

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

	mongoClient := mongoHealth.NewClient(session)

	mongoHealth := mongoHealth.CheckMongoClient{
		Client:      *mongoClient,
		Healthcheck: mongoClient.Healthcheck,
	}
  
	datastore := &store.DataStore{Backend: nil}
	if enableMongoData {
		//Create RecipeAPI instance with Mongo in datastore
		datastore.Backend = RecipeAPIStore{mongodb}
	}

	if err = hc.AddCheck("mongoDB", mongoHealth.Checker); err != nil {
		log.Event(ctx, "failed to add mongoDB checker", log.ERROR, log.Error(err))
		os.Exit(1)
	}
	hc.Start(ctx)

	apiErrors := make(chan error, 1)
	api.CreateAndInitialiseRecipeAPI(ctx, *cfg, *datastore, &hc, apiErrors)

	// block until a fatal error occurs
	select {
	case err := <-apiErrors:
		log.Event(ctx, "api error received", log.ERROR, log.Error(err))
	case <-signals:
		log.Event(ctx, "os signal received", log.INFO)
	}

	log.Event(ctx, fmt.Sprintf("shutdown with timeout: %s", cfg.GracefulShutdownTimeout), log.INFO)
	shutdownContext, cancel := context.WithTimeout(context.Background(), cfg.GracefulShutdownTimeout)

	// track shutdown gracefully closes app
	var gracefulShutdown bool

	// Gracefully shutdown the application closing any open resources.
	go func() {
		defer cancel()
		var hasShutdownError bool

		hc.Stop()

		// stop any incoming requests before closing any outbound connections
		if err = api.Close(shutdownContext); err != nil {
			log.Event(shutdownContext, "failed to close http server", log.ERROR, log.Error(err))
			hasShutdownError = true
		}

		if enableMongoData {
			if err = mongolib.Close(shutdownContext, mongodb.Session); err != nil {
				log.Event(shutdownContext, "failed to close mongo session", log.ERROR, log.Error(err))
				hasShutdownError = true
			}
		}

		log.Event(shutdownContext, "shutdown complete", log.INFO)

		if !hasShutdownError {
			gracefulShutdown = true
		}
	}()

	// wait for shutdown success (via cancel) or failure (timeout)
	<-shutdownContext.Done()

	if !gracefulShutdown {
		err = errors.New("failed to shutdown gracefully")
		log.Event(shutdownContext, "failed to shutdown gracefully ", log.ERROR, log.Error(err))
		os.Exit(1)
	}

	log.Event(shutdownContext, "graceful shutdown was successful", log.INFO)

	os.Exit(0)
}
