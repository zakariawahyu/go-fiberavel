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
	"github.com/zakariawahyu/go-fiberavel/internal/utils/helper"
)

type coupleUsecase struct {
	coupleRepo repository.CoupleRepository
	redis      cache.Rueidis
}

type CoupleUsecase interface {
	FindByID(ctx context.Context, id int64) (sqlc.GetCoupleRow, error)
	Store(ctx context.Context, req request.CreateCoupleRequest) (sqlc.Couple, error)
	Update(ctx context.Context, req request.UpdateCoupleRequest) error
	Destroy(ctx context.Context, id int64) error
	Datatables(ctx context.Context, params *datatables.DataTableParams) (*datatables.DataTableResponse, error)
	Publish(ctx context.Context) error
}

func NewCoupleUsecase(coupleRepo repository.CoupleRepository, redis cache.Rueidis) CoupleUsecase {
	return &coupleUsecase{
		coupleRepo: coupleRepo,
		redis:      redis,
	}
}

func (u *coupleUsecase) FindByID(ctx context.Context, id int64) (sqlc.GetCoupleRow, error) {
	return u.coupleRepo.FindByID(ctx, id)
}

func (u *coupleUsecase) Store(ctx context.Context, req request.CreateCoupleRequest) (sqlc.Couple, error) {
	var couple sqlc.CreateCoupleParams

	if err := smapping.FillStruct(&couple, smapping.MapFields(&req)); err != nil {
		return sqlc.Couple{}, err
	}

	result, err := u.coupleRepo.Insert(ctx, couple)
	if err != nil {
		return sqlc.Couple{}, err
	}

	return result, nil
}

func (u *coupleUsecase) Update(ctx context.Context, req request.UpdateCoupleRequest) error {
	var couple sqlc.UpdateCoupleParams

	if err := smapping.FillStruct(&couple, smapping.MapFields(&req)); err != nil {
		return err
	}

	couple.Image = helper.StringToPointer(*couple.Image)

	return u.coupleRepo.Update(ctx, couple)
}

func (u *coupleUsecase) Destroy(ctx context.Context, id int64) error {
	_, err := u.coupleRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return u.coupleRepo.Delete(ctx, id)
}

func (u *coupleUsecase) Datatables(ctx context.Context, params *datatables.DataTableParams) (*datatables.DataTableResponse, error) {
	var total, filtered int64

	search := params.Search
	orderColumn := map[string]string{
		"1": "couple_type",
		"2": "name",
	}

	arg := sqlc.DatatablesCoupleParams{
		Column1: search,
		Column2: orderColumn[params.OrderColumn],
		Column3: params.OrderDirection,
		Limit:   int32(params.Length),
		Offset:  int32(params.Start),
	}

	couples, err := u.coupleRepo.Datatables(ctx, arg)
	if err != nil {
		return nil, err
	}

	if couples == nil {
		couples = []sqlc.DatatablesCoupleRow{}
	}

	total, err = u.coupleRepo.Count(ctx, "")
	if err != nil {
		return nil, err
	}

	if search != "" {
		filtered, err = u.coupleRepo.Count(ctx, search)
		if err != nil {
			return nil, err
		}
	} else {
		filtered = total
	}

	return datatables.NewDataTableResponse(params.Draw, total, filtered, couples), nil
}

func (u *coupleUsecase) Publish(ctx context.Context) error {
	couples, err := u.coupleRepo.GetAll(ctx)
	if err != nil {
		return err
	}

	coupleBytes, err := json.Marshal(couples)
	if err != nil {
		return err
	}

	return u.redis.Set(constants.KeyCouples, coupleBytes, 0)
}
