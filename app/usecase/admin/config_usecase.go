package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/mashingan/smapping"
	"github.com/zakariawahyu/go-fiberavel/app/http/request"
	repository "github.com/zakariawahyu/go-fiberavel/app/repository/admin"
	"github.com/zakariawahyu/go-fiberavel/internal/infrastructure/cache"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/constants"
)

type configUsecase struct {
	configRepo repository.ConfigRepository
	redis      cache.Rueidis
}

type ConfigUsecase interface {
	FindByType(ctx context.Context, type_ string) (sqlc.GetConfigurationByTypeRow, error)
	Store(ctx context.Context, request request.ConfigRequest) error
}

func NewConfigUsecase(configRepo repository.ConfigRepository, redis cache.Rueidis) ConfigUsecase {
	return &configUsecase{
		configRepo: configRepo,
		redis:      redis,
	}
}

func (u *configUsecase) FindByType(ctx context.Context, type_ string) (sqlc.GetConfigurationByTypeRow, error) {
	configuration, err := u.configRepo.FindByType(ctx, type_)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return sqlc.GetConfigurationByTypeRow{}, err
	}

	return configuration, nil
}

func (u *configUsecase) Store(ctx context.Context, request request.ConfigRequest) error {
	var create sqlc.CreateConfigurationParams
	var update sqlc.UpdateConfigurationParams

	// Check if configuration already exists
	_, err := u.configRepo.FindByType(ctx, request.Type)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	// Insert or update configuration if not exists
	if errors.Is(err, pgx.ErrNoRows) {
		// Fill struct with create data
		if err := smapping.FillStruct(&create, smapping.MapFields(&request)); err != nil {
			return err
		}

		// Insert configuration if not exists
		if err := u.configRepo.Insert(ctx, create); err != nil {
			return err
		}
	} else {
		// Fill struct with update data
		if err := smapping.FillStruct(&update, smapping.MapFields(&request)); err != nil {
			return err
		}

		// Update configuration if exists
		if err := u.configRepo.Update(ctx, update); err != nil {
			return err
		}
	}

	// Get configuration data for redis
	configResult, err := u.configRepo.FindByType(ctx, request.Type)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	// Marshal configuration data to bytes
	configBytes, err := json.Marshal(configResult)
	if err != nil {
		return err
	}

	// Hash set configuration data to redis
	return u.redis.Hset(constants.KeyConfigs, configResult.Type, string(configBytes))
}