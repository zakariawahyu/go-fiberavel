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
	Insert(ctx context.Context, request sqlc.CreateConfigurationParams) error
	Update(ctx context.Context, request sqlc.UpdateConfigurationParams) error
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

func (r *configRepository) Insert(ctx context.Context, request sqlc.CreateConfigurationParams) error {
	return r.db.CreateConfiguration(ctx, request)
}

func (r *configRepository) Update(ctx context.Context, request sqlc.UpdateConfigurationParams) error {
	return r.db.UpdateConfiguration(ctx, request)
}
