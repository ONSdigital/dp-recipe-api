package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Configuration structure which hold information for configuring the import API
type Configuration struct {
	BindAddr    string `envconfig:"BIND_ADDR"`
	MongoConfig MongoConfig
}

// MongoConfig contains the config required to connect to MongoDB.
type MongoConfig struct {
	BindAddr          string `envconfig:"MONGODB_BIND_ADDR"   json:"-"`
	Collection        string `envconfig:"MONGODB_COLLECTION"`
	Database          string `envconfig:"MONGODB_DATABASE"`
	EnableMongoData   bool   `envconfig:"ENABLE_MONGO_DATA"`
	EnableMongoImport bool   `envconfig:"ENABLE_MONGO_IMPORT"`
}

var cfg *Configuration

// Get - configures the application and returns the cfg
func Get() (*Configuration, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Configuration{
		BindAddr: ":22300",
		MongoConfig: MongoConfig{
			BindAddr:          "localhost:27017",
			Collection:        "recipes",
			Database:          "recipes",
			EnableMongoData:   false,
			EnableMongoImport: false,
		},
	}

	return cfg, envconfig.Process("", cfg)
}
