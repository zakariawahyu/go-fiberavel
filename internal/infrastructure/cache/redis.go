package cache

import (
	"context"
	"fmt"
	"github.com/redis/rueidis"
	"github.com/zakariawahyu/go-fiberavel/config"
)

type Storage struct {
	Client rueidis.Client
}

func NewRedis(cfg *config.Config) (*Storage, error) {
	client, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress:  []string{fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)},
		Username:     cfg.Redis.Username,
		Password:     cfg.Redis.Password,
		SelectDB:     cfg.Redis.SelectDB,
		DisableCache: true,
	})

	if err != nil {
		return nil, err
	}

	// Test connection
	if err := client.Do(context.Background(), client.B().Ping().Build()).Error(); err != nil {
		return nil, err
	}

	// Create new storage
	return &Storage{
		Client: client,
	}, nil
}
