package repository

import (
	"context"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
)

type giftRepository struct {
	db *sqlc.Queries
}

type GiftRepository interface {
	GetAll(ctx context.Context) ([]sqlc.GetAllGiftRow, error)
	FindById(ctx context.Context, id int64) (sqlc.GetGiftRow, error)
	Insert(ctx context.Context, request sqlc.CreateGiftParams) (sqlc.Gift, error)
	Update(ctx context.Context, request sqlc.UpdateGiftParams) error
	Delete(ctx context.Context, id int64) error
	Datatables(ctx context.Context, arg sqlc.DatatablesGiftParams) ([]sqlc.DatatablesGiftRow, error)
	Count(ctx context.Context, search string) (int64, error)
}

func NewGiftRepository(db *sqlc.Queries) GiftRepository {
	return &giftRepository{
		db: db,
	}
}

func (r *giftRepository) GetAll(ctx context.Context) ([]sqlc.GetAllGiftRow, error) {
	gifts, err := r.db.GetAllGift(ctx)
	if err != nil {
		return nil, err
	}

	return gifts, nil
}

func (r *giftRepository) FindById(ctx context.Context, id int64) (sqlc.GetGiftRow, error) {
	gift, err := r.db.GetGift(ctx, id)
	if err != nil {
		return sqlc.GetGiftRow{}, err
	}

	return gift, nil
}

func (r *giftRepository) Insert(ctx context.Context, request sqlc.CreateGiftParams) (sqlc.Gift, error) {
	gift, err := r.db.CreateGift(ctx, request)
	if err != nil {
		return sqlc.Gift{}, err
	}

	return gift, nil
}

func (r *giftRepository) Update(ctx context.Context, request sqlc.UpdateGiftParams) error {
	return r.db.UpdateGift(ctx, request)
}

func (r *giftRepository) Delete(ctx context.Context, id int64) error {
	return r.db.DeleteGift(ctx, id)
}

func (r *giftRepository) Datatables(ctx context.Context, arg sqlc.DatatablesGiftParams) ([]sqlc.DatatablesGiftRow, error) {
	gifts, err := r.db.DatatablesGift(ctx, arg)
	if err != nil {
		return nil, err
	}

	return gifts, nil
}

func (r *giftRepository) Count(ctx context.Context, search string) (int64, error) {
	count, err := r.db.CountGift(ctx, search)
	if err != nil {
		return 0, err
	}

	return count, nil
}
