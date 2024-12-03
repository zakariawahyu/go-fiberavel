package repository

import (
	"context"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
)

type configRepository struct {
	db *sqlc.Queries
}

type ConfigRepository interface {
	FindByType(ctx context.Context, type_ string) (sqlc.GetConfigurationByTypeRow, error)
	UpdateCover(ctx context.Context, request sqlc.UpdateConfigurationCoverParams) error
	UpdateVenue(ctx context.Context, request sqlc.UpdateConfigurationVenueParams) error
	UpdateGift(ctx context.Context, request sqlc.UpdateConfigurationGiftParams) error
	UpdateWish(ctx context.Context, request sqlc.UpdateConfigurationWishParams) error
}

func NewConfigRepository(db *sqlc.Queries) *configRepository {
	return &configRepository{
		db: db,
	}
}

func (r *configRepository) FindByType(ctx context.Context, type_ string) (sqlc.GetConfigurationByTypeRow, error) {
	config, err := r.db.GetConfigurationByType(ctx, type_)
	if err != nil {
		return sqlc.GetConfigurationByTypeRow{}, err
	}

	return config, nil
}

func (r *configRepository) UpdateCover(ctx context.Context, request sqlc.UpdateConfigurationCoverParams) error {
	return r.db.UpdateConfigurationCover(ctx, request)
}

func (r *configRepository) UpdateVenue(ctx context.Context, request sqlc.UpdateConfigurationVenueParams) error {
	return r.db.UpdateConfigurationVenue(ctx, request)
}

func (r *configRepository) UpdateGift(ctx context.Context, request sqlc.UpdateConfigurationGiftParams) error {
	return r.db.UpdateConfigurationGift(ctx, request)
}

func (r *configRepository) UpdateWish(ctx context.Context, request sqlc.UpdateConfigurationWishParams) error {
	return r.db.UpdateConfigurationWish(ctx, request)
}
