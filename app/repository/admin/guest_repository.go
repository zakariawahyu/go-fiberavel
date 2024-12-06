package repository

import (
	"context"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
)

type guestRepository struct {
	db *sqlc.Queries
}

type GuestRepository interface {
	GetAll(ctx context.Context) ([]sqlc.GetAllGuestRow, error)
	FindById(ctx context.Context, id int64) (sqlc.GetGuestRow, error)
	Insert(ctx context.Context, request sqlc.CreateGuestParams) error
	Update(ctx context.Context, request sqlc.UpdateGuestParams) error
	Delete(ctx context.Context, id int64) error
	Datatables(ctx context.Context, arg sqlc.DatatablesGuestParams) ([]sqlc.DatatablesGuestRow, error)
	Count(ctx context.Context, search string) (int64, error)
}

func NewGuestRepository(db *sqlc.Queries) GuestRepository {
	return &guestRepository{
		db: db,
	}
}

func (r *guestRepository) GetAll(ctx context.Context) ([]sqlc.GetAllGuestRow, error) {
	guests, err := r.db.GetAllGuest(ctx)
	if err != nil {
		return nil, err
	}

	return guests, nil
}

func (r *guestRepository) FindById(ctx context.Context, id int64) (sqlc.GetGuestRow, error) {
	guest, err := r.db.GetGuest(ctx, id)
	if err != nil {
		return sqlc.GetGuestRow{}, err
	}

	return guest, nil
}

func (r *guestRepository) Insert(ctx context.Context, request sqlc.CreateGuestParams) error {
	return r.db.CreateGuest(ctx, request)
}

func (r *guestRepository) Update(ctx context.Context, request sqlc.UpdateGuestParams) error {
	return r.db.UpdateGuest(ctx, request)
}

func (r *guestRepository) Delete(ctx context.Context, id int64) error {
	return r.db.DeleteGuest(ctx, id)
}

func (r *guestRepository) Datatables(ctx context.Context, arg sqlc.DatatablesGuestParams) ([]sqlc.DatatablesGuestRow, error) {
	guests, err := r.db.DatatablesGuest(ctx, arg)
	if err != nil {
		return nil, err
	}

	return guests, nil
}

func (r *guestRepository) Count(ctx context.Context, search string) (int64, error) {
	count, err := r.db.CountGuest(ctx, search)
	if err != nil {
		return 0, err
	}

	return count, nil
}
