package config

type App struct {
	// Application Name
	// This value is the name of your application
	Name string `mapstructure:"APP_NAME"`

	// Application Port
	// This value is the server port for running
	Port string `mapstructure:"APP_PORT"`

	// Application Environment
	// This value determines the "environment" your application is currently running in.
	Environment string `mapstructure:"APP_ENVIRONMENT"`

	// Application URL
	URL string `mapstructure:"APP_URL"`

	// Application Image URL
	ImageURL string `mapstructure:"APP_IMAGE_URL"`

	// Application Timeout
	// This value to send timeout in context, time.Second
	Timeout int `mapstructure:"APP_TIMEOUT"`
}
