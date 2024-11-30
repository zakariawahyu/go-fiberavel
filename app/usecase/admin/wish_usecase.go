package usecase

import (
	"context"
	"encoding/json"
	"github.com/zakariawahyu/go-fiberavel/app/repository/admin"
	"github.com/zakariawahyu/go-fiberavel/internal/infrastructure/cache"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/constants"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/datatables"
)

type wishUsecase struct {
	wishRepo repository.WishRepository
	redis    cache.Rueidis
}

type WishUsecase interface {
	Destroy(ctx context.Context, id int64) error
	Datatables(ctx context.Context, params *datatables.DataTableParams) (*datatables.DataTableResponse, error)
	Publish(ctx context.Context) error
}

func NewWishUsecase(wishRepo repository.WishRepository, redis cache.Rueidis) WishUsecase {
	return &wishUsecase{
		wishRepo: wishRepo,
		redis:    redis,
	}
}

func (u *wishUsecase) Destroy(ctx context.Context, id int64) error {
	_, err := u.wishRepo.FindById(ctx, id)
	if err != nil {
		return err
	}

	return u.wishRepo.Delete(ctx, id)
}

func (u *wishUsecase) Datatables(ctx context.Context, params *datatables.DataTableParams) (*datatables.DataTableResponse, error) {
	var total, filtered int64
	var search = params.Search

	orderColumn := map[string]string{
		"1": "name",
		"2": "wish_description",
	}

	arg := sqlc.DatatablesWishParams{
		Column1: search,
		Column2: orderColumn[params.OrderColumn],
		Column3: params.OrderDirection,
		Limit:   int32(params.Length),
		Offset:  int32(params.Start),
	}

	wishes, err := u.wishRepo.Datatables(ctx, arg)
	if err != nil {
		return nil, err
	}

	if wishes == nil {
		wishes = []sqlc.DatatablesWishRow{}
	}

	total, err = u.wishRepo.Count(ctx, "")
	if err != nil {
		return nil, err
	}

	if search != "" {
		filtered, err = u.wishRepo.Count(ctx, search)
		if err != nil {
			return nil, err
		}
	} else {
		filtered = total
	}

	return datatables.NewDataTableResponse(params.Draw, total, filtered, wishes), nil
}

func (u *wishUsecase) Publish(ctx context.Context) error {
	wishes, err := u.wishRepo.GetAll(ctx)
	if err != nil {
		return err
	}

	venueBytes, err := json.Marshal(wishes)
	if err != nil {
		return err
	}

	return u.redis.Set(constants.KeyWishes, venueBytes, 0)
}
