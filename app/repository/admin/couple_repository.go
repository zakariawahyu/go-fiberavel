package repository

import (
	"context"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
)

type coupleRepository struct {
	db *sqlc.Queries
}

type CoupleRepository interface {
	Insert(ctx context.Context, request sqlc.CreateCoupleParams) (sqlc.Couple, error)
	Datatables(ctx context.Context, arg sqlc.DatatablesCoupleParams) ([]sqlc.DatatablesCoupleRow, error)
	Count(ctx context.Context, search string) (int64, error)
}

func NewCoupleRepository(db *sqlc.Queries) *coupleRepository {
	return &coupleRepository{
		db: db,
	}
}

func (r *coupleRepository) Insert(ctx context.Context, request sqlc.CreateCoupleParams) (sqlc.Couple, error) {
	couple, err := r.db.CreateCouple(ctx, request)
	if err != nil {
		return sqlc.Couple{}, err
	}

	return couple, nil
}

func (r *coupleRepository) Datatables(ctx context.Context, arg sqlc.DatatablesCoupleParams) ([]sqlc.DatatablesCoupleRow, error) {
	couples, err := r.db.DatatablesCouple(ctx, arg)
	if err != nil {
		return nil, err
	}

	return couples, nil
}
func (r *coupleRepository) Count(ctx context.Context, search string) (int64, error) {
	count, err := r.db.CountCouple(ctx, search)
	if err != nil {
		return 0, err
	}

	return count, nil
}
