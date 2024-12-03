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
	StoreCover(ctx context.Context, request request.ConfigCoverRequest) error
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

func (u *configUsecase) StoreCover(ctx context.Context, request request.ConfigCoverRequest) error {
	var cover sqlc.UpdateConfigurationCoverParams

	customData := map[string]any{
		"custom_data": map[string]string{
			"subtitle": request.Subtitle,
		},
	}

	customDataBytes, err := json.Marshal(customData)
	if err != nil {
		return err
	}

	cover.CustomData = customDataBytes
	if err := smapping.FillStruct(&cover, smapping.MapFields(&request)); err != nil {
		return nil
	}

	if err := u.configRepo.UpdateCover(ctx, cover); err != nil {
		return err
	}

	coverResult, err := u.configRepo.FindByType(ctx, cover.Type)
	if err != nil {
		return err
	}

	coverBytes, err := json.Marshal(coverResult)
	if err != nil {
		return err
	}

	return u.redis.Hset(constants.KeyConfigs, coverResult.Type, string(coverBytes))
}
