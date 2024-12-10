package config

import (
	"github.com/spf13/viper"
	"time"
)

type App struct {
	// Application Name
	// This value is the name of your application
	Name string

	// Application Port
	// This value is the server port for running
	Port string

	// Application Environment
	// This value determines the "environment" your application is currently running in.
	Environment string

	Key string

	// Application URL
	URL string

	// Application Timeout
	// This value to send timeout in context, time.Second
	Timeout time.Duration
}

func LoadApp() App {
	return App{
		Name:        viper.GetString("APP_NAME"),
		Port:        viper.GetString("APP_PORT"),
		Environment: viper.GetString("APP_ENVIRONMENT"),
		Key:         viper.GetString("APP_KEY"),
		URL:         viper.GetString("APP_URL"),
		Timeout:     viper.GetDuration("APP_TIMEOUT"),
	}
}
