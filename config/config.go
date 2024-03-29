package config

import (
	"time"

	mongodriver "github.com/ONSdigital/dp-mongodb/v3/mongodb"

	"github.com/kelseyhightower/envconfig"
)

type MongoConfig = mongodriver.MongoDriverConfig

// Configuration structure which hold information for configuring the import API
type Configuration struct {
	BindAddr                   string        `envconfig:"BIND_ADDR"`
	DefaultLimit               int           `envconfig:"DEFAULT_LIMIT"`
	DefaultMaxLimit            int           `envconfig:"DEFAULT_MAXIMUM_LIMIT"`
	DefaultOffset              int           `envconfig:"DEFAULT_OFFSET"`
	GracefulShutdownTimeout    time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT"`
	HealthCheckInterval        time.Duration `envconfig:"HEALTHCHECK_INTERVAL"`
	HealthCheckCriticalTimeout time.Duration `envconfig:"HEALTHCHECK_CRITICAL_TIMEOUT"`
	ZebedeeURL                 string        `envconfig:"ZEBEDEE_URL"`
	MongoConfig
}

var cfg *Configuration

const (
	RecipesCollection = "RecipesCollection"
)

// Get - configures the application and returns the cfg
func Get() (*Configuration, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Configuration{
		BindAddr:                   ":22300",
		DefaultLimit:               20,
		DefaultMaxLimit:            1000,
		DefaultOffset:              0,
		GracefulShutdownTimeout:    5 * time.Second,
		HealthCheckInterval:        30 * time.Second,
		HealthCheckCriticalTimeout: 90 * time.Second,
		ZebedeeURL:                 "http://localhost:8082",
		MongoConfig: MongoConfig{
			ClusterEndpoint:               "localhost:27017",
			Username:                      "",
			Password:                      "",
			Database:                      "recipes",
			Collections:                   map[string]string{RecipesCollection: "recipes"},
			ReplicaSet:                    "",
			IsStrongReadConcernEnabled:    false,
			IsWriteConcernMajorityEnabled: true,
			ConnectTimeout:                5 * time.Second,
			QueryTimeout:                  15 * time.Second,
			TLSConnectionConfig: mongodriver.TLSConnectionConfig{
				IsSSL:      false,
				VerifyCert: false,
			},
		},
	}

	return cfg, envconfig.Process("", cfg)
}
