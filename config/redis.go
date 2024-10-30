package config

type Redis struct {
	Host     string `mapstructure:"REDIS_HOST"`
	Port     string `mapstructure:"REDIS_PORT"`
	Username string `mapstructure:"REDIS_USERNAME"`
	Password string `mapstructure:"REDIS_PASSWORD"`
	SelectDB int    `mapstructure:"REDIS_SELECT_DB"`
}
