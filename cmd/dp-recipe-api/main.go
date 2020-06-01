package main

import (
	"context"
	"errors"
	"fmt"
	"syscall"

	"github.com/ONSdigital/dp-authorisation/auth"
	rchttp "github.com/ONSdigital/dp-rchttp"
	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/mux"

	"os"
	"os/signal"

	"github.com/ONSdigital/dp-api-clients-go/zebedee"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
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
	datastore := &store.DataStore{Backend: nil}

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
	} else {
		log.Event(ctx, "listening to mongo db session", log.INFO, log.Data{
			"mongo_bind_address": mongodb.URI,
		})
	}

	mongoClient := mongoHealth.NewClient(mongodb.Session)
	zebedeeClient := zebedee.New(cfg.ZebedeeURL)

	// Add dataset API and graph checks
	registerCheckers(ctx, &hc, mongoClient, zebedeeClient, cfg.MongoConfig.EnableAuthImport)

	//Create RecipeAPI instance with Mongo in datastore
	datastore.Backend = RecipeAPIStore{mongodb}

	hc.Start(ctx)

	apiErrors := make(chan error, 1)
	recipePermissions, permissions := getAuthorisationHandlers(ctx, cfg)
	api.CreateAndInitialiseRecipeAPI(ctx, *cfg, *datastore, &hc, apiErrors, recipePermissions, permissions)

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

		if err = mongolib.Close(shutdownContext, mongodb.Session); err != nil {
			log.Event(shutdownContext, "failed to close mongo session", log.ERROR, log.Error(err))
			hasShutdownError = true
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

func getAuthorisationHandlers(ctx context.Context, cfg *config.Configuration) (api.AuthHandler, api.AuthHandler) {
	if !cfg.MongoConfig.EnableAuthImport {
		log.Event(ctx, "feature flag not enabled defaulting to nop auth impl", log.INFO, log.Data{"feature": "ENABLE_AUTH_IMPORT"})
		return &auth.NopHandler{}, &auth.NopHandler{}
	}

	log.Event(ctx, "feature flag enabled", log.INFO, log.Data{"feature": "ENABLE_AUTH_IMPORT"})
	auth.LoggerNamespace("dp-recipe-api-auth")

	authClient := auth.NewPermissionsClient(rchttp.NewClient())
	authVerifier := auth.DefaultPermissionsVerifier()

	// for checking caller permissions when we have a recipeID, collection ID and user/service token
	recipePermissions := auth.NewHandler(
		auth.NewDatasetPermissionsRequestBuilder(cfg.ZebedeeURL, "id", mux.Vars),
		authClient,
		authVerifier,
	)

	// for checking caller permissions when we only have a user/service token
	permissions := auth.NewHandler(
		auth.NewPermissionsRequestBuilder(cfg.ZebedeeURL),
		authClient,
		authVerifier,
	)

	return recipePermissions, permissions
}

// registerCheckers adds the checkers for the provided clients to the health check object
func registerCheckers(ctx context.Context, hc *healthcheck.HealthCheck, mongoClient *mongoHealth.Client, zebedeeClient *zebedee.Client, EnableAuthImport bool) {
	var hasErrors bool
	if EnableAuthImport {
		if err := hc.AddCheck("Zebedee", zebedeeClient.Checker); err != nil {
			hasErrors = true
			log.Event(ctx, "error adding check for zebedeee", log.ERROR, log.Error(err))
		}
	}

	mongoHealth := mongoHealth.CheckMongoClient{
		Client:      *mongoClient,
		Healthcheck: mongoClient.Healthcheck,
	}
	if err := hc.AddCheck("mongoDB", mongoHealth.Checker); err != nil {
		hasErrors = true
		log.Event(ctx, "error adding mongoDB checker", log.FATAL, log.Error(err))
		os.Exit(1)
	}

	if hasErrors {
		log.Event(ctx, "error registering checkers for healthcheck", log.ERROR)
	}
}
