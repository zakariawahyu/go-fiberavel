package config

import "time"

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

	// Application URL
	URL string

	// Application Image URL
	ImageURL string

	// Application Timeout
	// This value to send timeout in context, time.Second
	Timeout time.Duration
}
