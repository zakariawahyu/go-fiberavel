package config

import "github.com/spf13/viper"

type Config struct {
	App   App
	Redis Redis
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(`./env.yaml`)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
