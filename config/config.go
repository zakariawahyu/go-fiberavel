package config

import "github.com/spf13/viper"

type Config struct {
	App      App
	Postgres Postgres
	Redis    Redis
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(`./.env`)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		App:      LoadApp(),
		Postgres: LoadPostgres(),
		Redis:    LoadRedis(),
	}, nil
}
