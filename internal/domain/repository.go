package domain

import (
	"go.uber.org/zap"
)

type Repository struct {
	//store        pgxutil.Store
	//txBeginner   pgxutil.TxBeginner
	logger *zap.Logger
}

func NewRepo(
	//pool *pgxpool.Pool,
	logger *zap.Logger) Repository {
	return Repository{
		//store:        pool,
		//txBeginner:   pool,
		logger: logger,
	}
}
