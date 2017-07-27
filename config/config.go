package config

import "github.com/ian-kent/gofigure"

type appConfiguration struct {
	BindAddr string `env:"BIND_ADDR" flag:"bind-addr" flagDesc:"The port to bind to"`
}

var configuration *appConfiguration

// Get - configures the application and returns the configuration
func Get() (*appConfiguration, error) {
	if configuration != nil {
		return configuration, nil
	}

	configuration = &appConfiguration{
		BindAddr: ":22300",
	}

	if err := gofigure.Gofigure(configuration); err != nil {
		return configuration, err
	}

	return configuration, nil
}
