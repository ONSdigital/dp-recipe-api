package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

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
	MongoConfig                MongoConfig
}

// MongoConfig contains the config required to connect to MongoDB.
type MongoConfig struct {
	BindAddr           string        `envconfig:"MONGODB_BIND_ADDR"   json:"-"`
	Collection         string        `envconfig:"MONGODB_COLLECTION"`
	Database           string        `envconfig:"MONGODB_DATABASE"`
	Username           string        `envconfig:"MONGODB_USERNAME"    json:"-"`
	Password           string        `envconfig:"MONGODB_PASSWORD"    json:"-"`
	IsSSL              bool          `envconfig:"MONGODB_IS_SSL"`
	EnableReadConcern  bool          `envconfig:"MONGODB_ENABLE_READ_CONCERN"`
	EnableWriteConcern bool          `envconfig:"MONGODB_ENABLE_WRITE_CONCERN"`
	QueryTimeout       time.Duration `envconfig:"MONGODB_QUERY_TIMEOUT"`
	ConnectionTimeout  time.Duration `envconfig:"MONGODB_CONNECT_TIMEOUT"`
}

var cfg *Configuration

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
			BindAddr:           "localhost:27017",
			Collection:         "recipes",
			Database:           "recipes",
			Username:           "",
			Password:           "",
			IsSSL:              false,
			QueryTimeout:       15 * time.Second,
			ConnectionTimeout:  5 * time.Second,
			EnableReadConcern:  false,
			EnableWriteConcern: true,
		},
	}

	return cfg, envconfig.Process("", cfg)
}
