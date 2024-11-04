package repository

import (
	"context"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
)

type rsvpRepository struct {
	db *sqlc.Queries
}

type RsvpRepository interface {
	CreateRsvp(ctx context.Context, request sqlc.CreateRsvpParams) (sqlc.Rsvp, error)
}

func NewRsvpRepository(db *sqlc.Queries) *rsvpRepository {
	return &rsvpRepository{
		db: db,
	}
}

func (r *rsvpRepository) CreateRsvp(ctx context.Context, request sqlc.CreateRsvpParams) (sqlc.Rsvp, error) {
	rsvp, err := r.db.CreateRsvp(ctx, request)
	if err != nil {
		return sqlc.Rsvp{}, err
	}

	return rsvp, nil
}
