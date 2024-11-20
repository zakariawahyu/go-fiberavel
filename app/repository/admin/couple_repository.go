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
