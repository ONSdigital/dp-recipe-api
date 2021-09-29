package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ONSdigital/dp-api-clients-go/zebedee"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	mongoHealth "github.com/ONSdigital/dp-mongodb/v3/health"
	"github.com/ONSdigital/dp-recipe-api/config"
	"github.com/ONSdigital/dp-recipe-api/mongo"
	"github.com/ONSdigital/dp-recipe-api/service"
	"github.com/ONSdigital/dp-recipe-api/store"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/pkg/errors"
)

// check that RecipeAPIStore satifies the the store.Storer interface
var _ store.Storer = (*RecipeAPIStore)(nil)

// RecipeAPIStore is a wrapper which embeds Mongo struct which between them satisfy the store.Storer interface.
type RecipeAPIStore struct {
	*mongo.Mongo
}

// health check variables - app version informaton retrieved on runtime
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

	if err := run(ctx); err != nil {
		log.Error(ctx, "application unexpectedly failed", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Create the service, providing an error channel for fatal errors
	svcErrors := make(chan error, 1)
	svcList := service.NewServiceList(&service.Init{})

	// Read config
	cfg, err := config.Get()
	if err != nil {
		log.Fatal(ctx, "failed to retrieve configuration", err)
		return err
	}
	log.Info(ctx, "config on startup", log.Data{"config": cfg, "build_time": BuildTime, "git-commit": GitCommit})

	// Run the service
	svc := service.New(cfg, svcList)
	if err := svc.Run(ctx, BuildTime, GitCommit, Version, svcErrors); err != nil {
		return errors.Wrap(err, "running service failed")
	}

	// Blocks until an os interrupt or a fatal error occurs
	select {
	case err := <-svcErrors:
		log.Error(ctx, "service error received", err)
	case sig := <-signals:
		log.Info(ctx, "os signal received", log.Data{"signal": sig})
	}
	return svc.Close(ctx)
}

// registerCheckers adds the checkers for the provided clients to the health check object
func registerCheckers(ctx context.Context, hc *healthcheck.HealthCheck, mongoClient *mongoHealth.Client, zebedeeClient *zebedee.Client) {
	var hasErrors bool
	if err := hc.AddCheck("Zebedee", zebedeeClient.Checker); err != nil {
		hasErrors = true
		log.Error(ctx, "error adding check for zebedeee", err)
	}

	mongoHealth := mongoHealth.CheckMongoClient{
		Client:      *mongoClient,
		Healthcheck: mongoClient.Healthcheck,
	}
	if err := hc.AddCheck("mongoDB", mongoHealth.Checker); err != nil {
		hasErrors = true
		log.Fatal(ctx, "error adding mongoDB checker", err)
		os.Exit(1)
	}

	if hasErrors {
		log.Error(ctx, "error registering checkers for healthcheck", errors.New("error registering checkers for healthcheck"))
	}
}
