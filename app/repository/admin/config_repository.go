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
