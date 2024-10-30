package config

import "github.com/spf13/viper"

type Postgres struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func LoadPostgres() Postgres {
	return Postgres{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		Username: viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
		Database: viper.GetString("DB_DATABASE"),
	}
}
