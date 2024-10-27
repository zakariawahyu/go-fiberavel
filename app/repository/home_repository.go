package repository

import (
	"context"
	"github.com/redis/rueidis"
	"github.com/zakariawahyu/go-fiberavel/internal/infrastructure/cache"
	"time"
)

type homeRepository struct {
	redis *cache.Storage
}

type HomeRepository interface {
	Get(ctx context.Context, key string) (string, error)
	HGet(ctx context.Context, key string, field string) (string, error)
	HGetAll(ctx context.Context, key string) (map[string]string, error)
}

func NewHomeRepository(redis *cache.Storage) *homeRepository {
	return &homeRepository{redis: redis}
}

func (r *homeRepository) Get(ctx context.Context, key string) (string, error) {
	val, err := r.redis.Client.DoCache(ctx, r.redis.Client.B().Get().Key(key).Cache(), 1*time.Hour).ToString()
	if err != nil && rueidis.IsRedisNil(err) {
		return "", nil
	}
	return val, err
}

func (r *homeRepository) HGet(ctx context.Context, key string, field string) (string, error) {
	val, err := r.redis.Client.DoCache(ctx, r.redis.Client.B().Hget().Key(key).Field(field).Cache(), 1*time.Hour).ToString()
	if err != nil && rueidis.IsRedisNil(err) {
		return "", nil
	}
	return val, err
}

func (r *homeRepository) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	val, err := r.redis.Client.DoCache(ctx, r.redis.Client.B().Hgetall().Key(key).Cache(), 1*time.Hour).AsStrMap()
	if err != nil && rueidis.IsRedisNil(err) {
		return nil, nil
	}
	return val, err
}
