package service

import (
	"context"
	"net/http"

	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	dphttp "github.com/ONSdigital/dp-net/http"
	"github.com/ONSdigital/dp-recipe-api/config"
	"github.com/ONSdigital/dp-recipe-api/mongo"
	"github.com/ONSdigital/dp-recipe-api/store"
	"github.com/ONSdigital/log.go/v2/log"
)

// ExternalServiceList holds the initialiser and initialisation state of external services.
type ExternalServiceList struct {
	HealthCheck bool
	MongoDB     bool
	Init        Initialiser
}

// Init implements the Initialiser interface to initialise dependencies
type Init struct{}

// NewServiceList creates a new service list with the provided initialiser
func NewServiceList(initialiser Initialiser) *ExternalServiceList {
	return &ExternalServiceList{
		Init: initialiser,
	}
}

// GetHTTPServer creates an http server
func (e *ExternalServiceList) GetHTTPServer(bindAddr string, router http.Handler) HTTPServer {
	s := e.Init.DoGetHTTPServer(bindAddr, router)
	return s
}

// GetHealthCheck creates a healthcheck with versionInfo and sets the HealthCheck flag to true
func (e *ExternalServiceList) GetHealthCheck(cfg *config.Configuration, buildTime, gitCommit, version string) (HealthChecker, error) {
	hc, err := e.Init.DoGetHealthCheck(cfg, buildTime, gitCommit, version)
	if err != nil {
		return nil, err
	}
	e.HealthCheck = true
	return hc, nil
}

// GetMongoDB returns a mongodb health client and dataset mongo object
func (e *ExternalServiceList) GetMongoDB(ctx context.Context, cfg *config.Configuration) (store.MongoDB, error) {
	mongodb, err := e.Init.DoGetMongoDB(ctx, cfg)
	if err != nil {
		log.Error(ctx, "failed to initialise mongo", err)
		return nil, err
	}
	e.MongoDB = true
	return mongodb, nil
}

// DoGetHTTPServer creates an HTTP Server with the provided bind address and router
func (e *Init) DoGetHTTPServer(bindAddr string, router http.Handler) HTTPServer {
	s := dphttp.NewServer(bindAddr, router)
	s.HandleOSSignals = false
	return s
}

// DoGetHealthCheck creates a healthcheck with versionInfo
func (e *Init) DoGetHealthCheck(cfg *config.Configuration, buildTime, gitCommit, version string) (HealthChecker, error) {
	versionInfo, err := healthcheck.NewVersionInfo(buildTime, gitCommit, version)
	if err != nil {
		return nil, err
	}
	hc := healthcheck.New(versionInfo, cfg.HealthCheckCriticalTimeout, cfg.HealthCheckInterval)
	return &hc, nil
}

// DoGetMongoDB returns a MongoDB
func (e *Init) DoGetMongoDB(ctx context.Context, cfg *config.Configuration) (store.MongoDB, error) {
	mongodb := &mongo.Mongo{
		Collection:                 cfg.MongoConfig.Collection,
		Database:                   cfg.MongoConfig.Database,
		URI:                        cfg.MongoConfig.BindAddr,
		Username:                   cfg.MongoConfig.Username,
		Password:                   cfg.MongoConfig.Password,
		IsSSL:                      cfg.MongoConfig.IsSSL,
		QueryTimeoutInSeconds:      cfg.MongoConfig.QueryTimeout,
		ConnectionTimeoutInSeconds: cfg.MongoConfig.ConnectionTimeout,
	}
	if err := mongodb.Init(ctx, cfg.MongoConfig.EnableReadConcern, cfg.MongoConfig.EnableWriteConcern); err != nil {
		return nil, err
	}
	log.Info(ctx, "listening to mongo db session", log.Data{"URI": mongodb.URI})
	return mongodb, nil
}
