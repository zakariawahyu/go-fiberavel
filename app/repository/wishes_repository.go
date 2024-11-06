package repository

import (
	"context"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
)

type wishRepository struct {
	db *sqlc.Queries
}

type WishesRepository interface {
	CreateWish(ctx context.Context, request sqlc.CreateWishParams) (sqlc.Wish, error)
	GetAllWishes(ctx context.Context) ([]sqlc.GetAllWishesRow, error)
}

func NewWishesRepository(db *sqlc.Queries) *wishRepository {
	return &wishRepository{
		db: db,
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
