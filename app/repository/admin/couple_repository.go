package repository

import (
	"context"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
)

type coupleRepository struct {
	db *sqlc.Queries
}

type CoupleRepository interface {
	GetAll(ctx context.Context) ([]sqlc.GetAllCoupleRow, error)
	FindByID(ctx context.Context, id int64) (sqlc.GetCoupleRow, error)
	Insert(ctx context.Context, request sqlc.CreateCoupleParams) (sqlc.Couple, error)
	Update(ctx context.Context, request sqlc.UpdateCoupleParams) error
	Delete(ctx context.Context, id int64) error
	Datatables(ctx context.Context, arg sqlc.DatatablesCoupleParams) ([]sqlc.DatatablesCoupleRow, error)
	Count(ctx context.Context, search string) (int64, error)
}

func NewCoupleRepository(db *sqlc.Queries) *coupleRepository {
	return &coupleRepository{
		db: db,
	}
}
func (r *coupleRepository) GetAll(ctx context.Context) ([]sqlc.GetAllCoupleRow, error) {
	couples, err := r.db.GetAllCouple(ctx)
	if err != nil {
		return nil, err
	}

	return couples, nil
}

func (r *coupleRepository) FindByID(ctx context.Context, id int64) (sqlc.GetCoupleRow, error) {
	couple, err := r.db.GetCouple(ctx, id)
	if err != nil {
		return sqlc.GetCoupleRow{}, err
	}

	return couple, nil
}

func (r *coupleRepository) Insert(ctx context.Context, request sqlc.CreateCoupleParams) (sqlc.Couple, error) {
	couple, err := r.db.CreateCouple(ctx, request)
	if err != nil {
		return sqlc.Couple{}, err
	}

	return couple, nil
}

func (r *coupleRepository) Update(ctx context.Context, request sqlc.UpdateCoupleParams) error {
	return r.db.UpdateCouple(ctx, request)
}

func (r *coupleRepository) Delete(ctx context.Context, id int64) error {
	return r.db.DeleteCouple(ctx, id)
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
