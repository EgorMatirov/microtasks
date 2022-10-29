package infrastructure

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func InitPGPool(ctx context.Context, user, pswd, host, port, dbName string, maxConn int32) (*pgxpool.Pool, error) {
	pgConString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		user, pswd, host, port, dbName,
	)

	config, err := pgxpool.ParseConfig(pgConString)
	if err != nil {
		return nil, err
	}

	config.MaxConns = maxConn
	config.ConnConfig.PreferSimpleProtocol = true

	dbPool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return dbPool, nil
}
