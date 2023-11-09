package psql

import (
	"context"
	"time"

	"test_service/internal/core/repository/psql/sqlc"
	"test_service/internal/pkg/logger"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type SQLStore struct {
	DB *pgxpool.Pool
	*sqlc.Queries
}

func NewStore(ctx context.Context, psqlUri string) *SQLStore {
	logger.Log.Info("connecting to psql...")
	dbConn, err := pgxpool.Connect(ctx, psqlUri)
	if err != nil {
		logger.Log.Fatal("failed to connecto to psql", zap.Error(err))
	}

	if err := dbConn.Ping(ctx); err != nil {
		logger.Log.Fatal("failed to ping psql", zap.Error(err))
		time.Sleep(5 * time.Second)
	}

	logger.Log.Info("psql connected")
	return &SQLStore{
		DB:      dbConn,
		Queries: sqlc.New(dbConn),
	}
}
