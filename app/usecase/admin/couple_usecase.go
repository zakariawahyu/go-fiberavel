package usecase

import (
	"context"
	"github.com/mashingan/smapping"
	"github.com/zakariawahyu/go-fiberavel/app/http/request"
	repository "github.com/zakariawahyu/go-fiberavel/app/repository/admin"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
)

type coupleUsecase struct {
	coupleRepo repository.CoupleRepository
}

type CoupleUsecase interface {
	Store(ctx context.Context, req request.CreateCoupleRequest) (sqlc.Couple, error)
}

func NewCoupleUsecase(coupleRepo repository.CoupleRepository) CoupleUsecase {
	return &coupleUsecase{
		coupleRepo: coupleRepo,
	}
}

func (u coupleUsecase) Store(ctx context.Context, req request.CreateCoupleRequest) (sqlc.Couple, error) {
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
