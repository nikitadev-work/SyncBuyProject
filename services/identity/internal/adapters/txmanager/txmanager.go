package txmanager

import (
	"context"
	"identity/internal/usecase"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nikitadev-work/SyncBuyProject/common/kit/logger"
)

type TxManager struct {
	pgPool *pgxpool.Pool
	logger logger.LoggerInterface
}

var _ usecase.TxManagerInterface = (*TxManager)(nil)

func NewTxManager(pgPool *pgxpool.Pool, logger logger.LoggerInterface) *TxManager {
	return &TxManager{
		pgPool: pgPool,
		logger: logger,
	}
}

func (txM *TxManager) WithinTx(ctx context.Context, fn func(context.Context) error) error {
	tx, err := txM.pgPool.Begin(ctx)
	if err != nil {
		txM.logger.Error("failed to begin transaction", map[string]any{
			"error": err.Error(),
		})
	}

	txCtx := context.WithValue(ctx, "txMarker", tx)

	err = fn(txCtx)
	if err != nil {
		tx.Rollback(ctx)
	}

	tx.Commit(ctx)

	txM.logger.Info("successfully finished transaction", nil)

	return nil
}
