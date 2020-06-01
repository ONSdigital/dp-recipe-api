package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Configuration structure which hold information for configuring the import API
type Configuration struct {
	BindAddr                   string        `envconfig:"BIND_ADDR"`
	ZebedeeURL                 string        `envconfig:"ZEBEDEE_URL"`
	GracefulShutdownTimeout    time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT"`
	HealthCheckInterval        time.Duration `envconfig:"HEALTHCHECK_INTERVAL"`
	HealthCheckCriticalTimeout time.Duration `envconfig:"HEALTHCHECK_CRITICAL_TIMEOUT"`
	MongoConfig                MongoConfig
}

// MongoConfig contains the config required to connect to MongoDB.
type MongoConfig struct {
	BindAddr         string `envconfig:"MONGODB_BIND_ADDR"   json:"-"`
	Collection       string `envconfig:"MONGODB_COLLECTION"`
	Database         string `envconfig:"MONGODB_DATABASE"`
	EnableAuthImport bool   `envconfig:"ENABLE_AUTH_IMPORT"`
}

var cfg *Configuration

// Get - configures the application and returns the cfg
func Get() (*Configuration, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Configuration{
		BindAddr:                   ":22300",
		ZebedeeURL:                 "http://localhost:8082",
		GracefulShutdownTimeout:    5 * time.Second,
		HealthCheckInterval:        30 * time.Second,
		HealthCheckCriticalTimeout: 90 * time.Second,
		MongoConfig: MongoConfig{
			BindAddr:         "localhost:27017",
			Collection:       "recipes",
			Database:         "recipes",
			EnableAuthImport: true,
		},
	}

	return cfg, envconfig.Process("", cfg)
}
