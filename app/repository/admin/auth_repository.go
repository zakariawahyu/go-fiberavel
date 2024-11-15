package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/zakariawahyu/go-fiberavel/app/http/middleware"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
)

type authRepository struct {
	db *sqlc.Queries
}

type AuthRepository interface {
	Login(ctx context.Context, username string) (sqlc.LoginRow, error)
}

func NewAuthRepository(db *sqlc.Queries) *authRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) Login(ctx context.Context, username string) (sqlc.LoginRow, error) {
	auth, err := r.db.Login(ctx, username)
	if err != nil {
		if err == pgx.ErrNoRows {
			return sqlc.LoginRow{}, middleware.ErrLogin
		}
		return sqlc.LoginRow{}, err
	}

	return auth, nil
}
