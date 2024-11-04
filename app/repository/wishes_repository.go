package repository

import (
	"context"
	"github.com/redis/rueidis"
	"github.com/zakariawahyu/go-fiberavel/internal/infrastructure/cache"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
	"time"
)

type wishRepository struct {
	db    *sqlc.Queries
	redis *cache.Storage
}

type WishesRepository interface {
	CreateWish(ctx context.Context, request sqlc.CreateWishParams) (sqlc.Wish, error)
	GetAllWishes(ctx context.Context) ([]sqlc.GetAllWishesRow, error)
	GetRedis(ctx context.Context, key string) (string, error)
	SetRedis(ctx context.Context, key string, value string) error
}

func NewWishesRepository(db *sqlc.Queries, redis *cache.Storage) *wishRepository {
	return &wishRepository{
		db:    db,
		redis: redis,
	}
}

func (r *wishRepository) CreateWish(ctx context.Context, request sqlc.CreateWishParams) (sqlc.Wish, error) {
	wish, err := r.db.CreateWish(ctx, request)
	if err != nil {
		return sqlc.Wish{}, err
	}

	return wish, nil
}

func (r *wishRepository) GetAllWishes(ctx context.Context) ([]sqlc.GetAllWishesRow, error) {
	wishes, err := r.db.GetAllWishes(ctx)
	if err != nil {
		return nil, err
	}

	return wishes, nil
}

func (r *wishRepository) GetRedis(ctx context.Context, key string) (string, error) {
	val, err := r.redis.Client.DoCache(ctx, r.redis.Client.B().Get().Key(key).Cache(), 1*time.Hour).ToString()
	if err != nil && rueidis.IsRedisNil(err) {
		return "", nil
	}

	return val, err
}

func (r *wishRepository) SetRedis(ctx context.Context, key string, value string) error {
	return r.redis.Client.Do(ctx, r.redis.Client.B().Set().Key(key).Value(value).Build()).Error()
}
