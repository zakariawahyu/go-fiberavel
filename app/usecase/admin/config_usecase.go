package usecase

import (
	"context"
	"encoding/json"
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
	return u.configRepo.FindByType(ctx, type_)
}

func (u *configUsecase) StoreCover(ctx context.Context, request request.ConfigCoverRequest) error {
	var cover sqlc.UpdateConfigurationCoverParams
	cover.CustomData = []byte(request.Subtitle)
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
