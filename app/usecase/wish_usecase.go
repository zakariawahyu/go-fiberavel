package usecase

import (
	"context"
	"encoding/json"
	"github.com/mashingan/smapping"
	"github.com/zakariawahyu/go-fiberavel/app/http/request"
	"github.com/zakariawahyu/go-fiberavel/app/repository"
	"github.com/zakariawahyu/go-fiberavel/internal/infrastructure/cache"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/constants"
)

type wishUsecase struct {
	wishRepo repository.WishRepository
	redis    cache.Rueidis
}

type WishUsecase interface {
	GetAll() ([]byte, error)
	Store(ctx context.Context, request request.CreateWishParams, path string) (sqlc.Wish, error)
}

func NewWishUsecase(wishRepo repository.WishRepository, redis cache.Rueidis) WishUsecase {
	return &wishUsecase{
		wishRepo: wishRepo,
		redis:    redis,
	}
}

func (u *wishUsecase) GetAll() ([]byte, error) {
	wishes, err := u.redis.Get(constants.KeyWishes)
	if err != nil {
		return nil, err
	}

	return wishes, nil
}

func (u *wishUsecase) Store(ctx context.Context, request request.CreateWishParams, path string) (sqlc.Wish, error) {
	var wish sqlc.CreateWishParams

	if err := smapping.FillStruct(&wish, smapping.MapFields(&request)); err != nil {
		return sqlc.Wish{}, err
	}

	result, err := u.wishRepo.Insert(ctx, wish)
	if err != nil {
		return sqlc.Wish{}, err
	}

	wishes, err := u.wishRepo.GetAll(ctx)
	if err != nil {
		return sqlc.Wish{}, err
	}

	wishesBytes, err := json.Marshal(wishes)
	if err != nil {
		return sqlc.Wish{}, err
	}

	if err = u.redis.Delete(path); err != nil {
		return sqlc.Wish{}, err
	}

	err = u.redis.Set(constants.KeyWishes, wishesBytes, 0)
	if err != nil {
		return sqlc.Wish{}, err
	}

	return result, nil
}
