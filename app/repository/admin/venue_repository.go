package repository

import (
	"context"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
)

type venueRepository struct {
	db *sqlc.Queries
}

type VenueRepository interface {
	GetAll(ctx context.Context) ([]sqlc.GetAllVenueRow, error)
	FindById(ctx context.Context, id int64) (sqlc.GetVenueRow, error)
	Insert(ctx context.Context, request sqlc.CreateVenueParams) (sqlc.Venue, error)
	Update(ctx context.Context, request sqlc.UpdateVenueParams) error
	Delete(ctx context.Context, id int64) error
	Datatables(ctx context.Context, arg sqlc.DatatablesVenueParams) ([]sqlc.DatatablesVenueRow, error)
	Count(ctx context.Context, search string) (int64, error)
}

func NewVenueRepository(db *sqlc.Queries) VenueRepository {
	return &venueRepository{
		db: db,
	}
}

func (r *venueRepository) GetAll(ctx context.Context) ([]sqlc.GetAllVenueRow, error) {
	venues, err := r.db.GetAllVenue(ctx)
	if err != nil {
		return nil, err
	}

	return venues, nil
}

func (r *venueRepository) FindById(ctx context.Context, id int64) (sqlc.GetVenueRow, error) {
	venue, err := r.db.GetVenue(ctx, id)
	if err != nil {
		return sqlc.GetVenueRow{}, err
	}

	return venue, nil
}

func (r *venueRepository) Insert(ctx context.Context, request sqlc.CreateVenueParams) (sqlc.Venue, error) {
	venue, err := r.db.CreateVenue(ctx, request)
	if err != nil {
		return sqlc.Venue{}, err
	}

	return venue, nil
}

func (r *venueRepository) Update(ctx context.Context, request sqlc.UpdateVenueParams) error {
	return r.db.UpdateVenue(ctx, request)
}

func (r *venueRepository) Delete(ctx context.Context, id int64) error {
	return r.db.DeleteVenue(ctx, id)
}

func (r *venueRepository) Datatables(ctx context.Context, arg sqlc.DatatablesVenueParams) ([]sqlc.DatatablesVenueRow, error) {
	venues, err := r.db.DatatablesVenue(ctx, arg)
	if err != nil {
		return nil, err
	}

	return venues, nil
}

func (r *venueRepository) Count(ctx context.Context, search string) (int64, error) {
	count, err := r.db.CountVenue(ctx, search)
	if err != nil {
		return 0, err
	}

	return count, nil
}
