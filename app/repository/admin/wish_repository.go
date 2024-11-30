package repository

import (
	"context"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
)

type wishRepository struct {
	db *sqlc.Queries
}

type WishRepository interface {
	GetAll(ctx context.Context) ([]sqlc.GetAllWishRow, error)
	FindById(ctx context.Context, id int64) (sqlc.GetWishRow, error)
	Delete(ctx context.Context, id int64) error
	Datatables(ctx context.Context, arg sqlc.DatatablesWishParams) ([]sqlc.DatatablesWishRow, error)
	Count(ctx context.Context, search string) (int64, error)
}

func NewWishRepository(db *sqlc.Queries) WishRepository {
	return &wishRepository{
		db: db,
	}
}

func (r *wishRepository) GetAll(ctc context.Context) ([]sqlc.GetAllWishRow, error) {
	wishes, err := r.db.GetAllWish(ctc)
	if err != nil {
		return nil, err
	}

	return wishes, nil
}

func (r *wishRepository) FindById(ctx context.Context, id int64) (sqlc.GetWishRow, error) {
	wish, err := r.db.GetWish(ctx, id)
	if err != nil {
		return sqlc.GetWishRow{}, err
	}

	return wish, nil

}

func (r *wishRepository) Delete(ctx context.Context, id int64) error {
	return r.db.DeleteWish(ctx, id)
}

func (r *wishRepository) Datatables(ctx context.Context, arg sqlc.DatatablesWishParams) ([]sqlc.DatatablesWishRow, error) {
	wishes, err := r.db.DatatablesWish(ctx, arg)
	if err != nil {
		return nil, err
	}

	return wishes, nil
}

func (r *wishRepository) Count(ctx context.Context, search string) (int64, error) {
	count, err := r.db.CountWish(ctx, search)
	if err != nil {
		return 0, err
	}

	return count, nil
}
