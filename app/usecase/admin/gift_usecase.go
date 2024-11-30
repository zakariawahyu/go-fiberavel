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
	"github.com/zakariawahyu/go-fiberavel/internal/utils/datatables"
)

type giftUsecase struct {
	giftRepo repository.GiftRepository
	redis    cache.Rueidis
}

type GiftUsecase interface {
	FindById(ctx context.Context, id int64) (sqlc.GetGiftRow, error)
	Store(ctx context.Context, request request.CreateGiftRequest) (sqlc.Gift, error)
	Update(ctx context.Context, request request.UpdateGiftRequest) error
	Destroy(ctx context.Context, id int64) error
	Datatables(ctx context.Context, params *datatables.DataTableParams) (*datatables.DataTableResponse, error)
	Publish(ctx context.Context) error
}

func NewGiftUsecase(giftRepo repository.GiftRepository, redis cache.Rueidis) GiftUsecase {
	return &giftUsecase{
		giftRepo: giftRepo,
		redis:    redis,
	}
}

func (u *giftUsecase) FindById(ctx context.Context, id int64) (sqlc.GetGiftRow, error) {
	return u.giftRepo.FindById(ctx, id)
}

func (u *giftUsecase) Store(ctx context.Context, request request.CreateGiftRequest) (sqlc.Gift, error) {
	var gift sqlc.CreateGiftParams

	if err := smapping.FillStruct(&gift, smapping.MapFields(&request)); err != nil {
		return sqlc.Gift{}, err
	}

	result, err := u.giftRepo.Insert(ctx, gift)
	if err != nil {
		return sqlc.Gift{}, err
	}

	return result, nil
}

func (u *giftUsecase) Update(ctx context.Context, request request.UpdateGiftRequest) error {
	var gift sqlc.UpdateGiftParams

	if err := smapping.FillStruct(&gift, smapping.MapFields(&request)); err != nil {
		return err
	}

	return u.giftRepo.Update(ctx, gift)
}

func (u *giftUsecase) Destroy(ctx context.Context, id int64) error {
	_, err := u.giftRepo.FindById(ctx, id)
	if err != nil {
		return err
	}

	return u.giftRepo.Delete(ctx, id)
}

func (u *giftUsecase) Datatables(ctx context.Context, params *datatables.DataTableParams) (*datatables.DataTableResponse, error) {
	var total, filtered int64
	var search = params.Search

	orderColumn := map[string]string{
		"1": "name",
		"2": "account_name",
	}

	arg := sqlc.DatatablesGiftParams{
		Column1: search,
		Column2: orderColumn[params.OrderColumn],
		Column3: params.OrderDirection,
		Limit:   int32(params.Length),
		Offset:  int32(params.Start),
	}

	gifts, err := u.giftRepo.Datatables(ctx, arg)
	if err != nil {
		return nil, err
	}

	if gifts == nil {
		gifts = []sqlc.DatatablesGiftRow{}
	}

	total, err = u.giftRepo.Count(ctx, "")
	if err != nil {
		return nil, err
	}

	if search != "" {
		filtered, err = u.giftRepo.Count(ctx, search)
		if err != nil {
			return nil, err
		}
	} else {
		filtered = total
	}

	return datatables.NewDataTableResponse(params.Draw, total, filtered, gifts), nil
}

func (u *giftUsecase) Publish(ctx context.Context) error {
	gifts, err := u.giftRepo.GetAll(ctx)
	if err != nil {
		return err
	}

	giftBytes, err := json.Marshal(gifts)
	if err != nil {
		return err
	}

	return u.redis.Set(constants.KeyGift, giftBytes, 0)
}
