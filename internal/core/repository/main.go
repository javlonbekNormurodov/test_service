package repository

import (
	"context"

	"test_service/internal/core/repository/psql"
	"test_service/internal/core/repository/psql/sqlc"
)

type Store interface {
	sqlc.Querier
}

func New(ctx context.Context, dsn string) Store {
	return psql.NewStore(ctx, dsn)
}
