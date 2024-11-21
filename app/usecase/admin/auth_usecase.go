package usecase

import (
	"context"
	"github.com/mashingan/smapping"
	"github.com/zakariawahyu/go-fiberavel/app/http/middleware"
	"github.com/zakariawahyu/go-fiberavel/app/http/request"
	repository "github.com/zakariawahyu/go-fiberavel/app/repository/admin"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/helper"
)

type authUsecase struct {
	authRepo repository.AuthRepository
}

type AuthUsecase interface {
	Login(ctx context.Context, req request.LoginRequest) (sqlc.LoginRow, error)
}

func NewAuthhUsecase(authRepo repository.AuthRepository) AuthUsecase {
	return &authUsecase{
		authRepo: authRepo,
	}
}

func (u authUsecase) Login(ctx context.Context, req request.LoginRequest) (sqlc.LoginRow, error) {
	var auth sqlc.LoginRow

	if err := smapping.FillStruct(&auth, smapping.MapFields(req)); err != nil {
		return sqlc.LoginRow{}, err
	}

	result, err := u.authRepo.Login(ctx, auth.Username)
	if err != nil {
		return sqlc.LoginRow{}, err
	}

	if err := helper.ComparePassword(result.Password, auth.Password); err != nil {
		return sqlc.LoginRow{}, middleware.ErrPasswordNotMatch
	}

	return result, nil
}
