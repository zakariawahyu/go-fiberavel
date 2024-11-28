package usecase

import (
	"context"
	"github.com/mashingan/smapping"
	"github.com/zakariawahyu/go-fiberavel/app/http/request"
	repository "github.com/zakariawahyu/go-fiberavel/app/repository/admin"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/datatables"
)

type coupleUsecase struct {
	coupleRepo repository.CoupleRepository
}

type CoupleUsecase interface {
	Store(ctx context.Context, req request.CreateCoupleRequest) (sqlc.Couple, error)
	Datatables(ctx context.Context, params *datatables.DataTableParams) (*datatables.DataTableResponse, error)
}

func NewCoupleUsecase(coupleRepo repository.CoupleRepository) CoupleUsecase {
	return &coupleUsecase{
		coupleRepo: coupleRepo,
	}
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

func (u *coupleUsecase) Datatables(ctx context.Context, params *datatables.DataTableParams) (*datatables.DataTableResponse, error) {
	var total, filtered int64

	search := params.Search
	orderColumn := map[string]string{
		"0": "couple_type",
		"1": "name",
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
