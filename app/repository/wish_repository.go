package repository

import (
	"context"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
)

type wishRepository struct {
	db *sqlc.Queries
}

type WishRepository interface {
	Insert(ctx context.Context, request sqlc.CreateWishParams) (sqlc.Wish, error)
	GetAll(ctx context.Context) ([]sqlc.GetAllWishRow, error)
}

func NewWishesRepository(db *sqlc.Queries) WishRepository {
	return &wishRepository{
		db: db,
	}
}

func (r *wishRepository) Insert(ctx context.Context, request sqlc.CreateWishParams) (sqlc.Wish, error) {
	wish, err := r.db.CreateWish(ctx, request)
	if err != nil {
		return sqlc.Wish{}, err
	}

	return wish, nil
}

func (r *wishRepository) GetAll(ctx context.Context) ([]sqlc.GetAllWishRow, error) {
	wishes, err := r.db.GetAllWish(ctx)
	if err != nil {
		return nil, err
	}

	return wishes, nil
}
