package repository

import (
	"context"
	"github.com/zakariawahyu/go-fiberavel/internal/infrastructure/cache"
)

type homeRepository struct {
	redis *cache.Storage
}

type HomeRepository interface {
	Get(ctx context.Context, key string) (map[string]string, error)
}

func NewHomeRepository(redis *cache.Storage) *homeRepository {
	return &homeRepository{redis: redis}
}

func (r *homeRepository) Get(ctx context.Context, key string) (map[string]string, error) {
	return r.redis.Client.Do(ctx, r.redis.Client.B().Hgetall().Key(key).Build()).AsStrMap()
}
