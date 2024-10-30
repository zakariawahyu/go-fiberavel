package config

import "github.com/spf13/viper"

type Redis struct {
	Host     string
	Port     string
	Username string
	Password string
	SelectDB int
}

func LoadRedis() Redis {
	return Redis{
		Host:     viper.GetString("REDIS_HOST"),
		Port:     viper.GetString("REDIS_PORT"),
		Username: viper.GetString("REDIS_USERNAME"),
		Password: viper.GetString("REDIS_PASSWORD"),
		SelectDB: viper.GetInt("REDIS_SELECT_DB"),
	}
}
