package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/zakariawahyu/go-fiberavel/config"
)

type Postgres struct {
	Conn *pgx.Conn
}

func NewPostgres(cfg *config.Config) (*Postgres, error) {
	ctx := context.Background()

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Database,
	)

	conn, err := pgx.Connect(ctx, dataSourceName)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		Conn: conn,
	}, nil
}
