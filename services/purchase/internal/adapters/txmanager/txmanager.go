package txmanager

import (
	"context"
	"purchase/internal/usecase"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nikitadev-work/SyncBuyProject/common/kit/logger"
)

type TxManager struct {
	pgPool   *pgxpool.Pool
	logger   logger.LoggerInterface
	txMarker string
}

var _ usecase.TxManagerInterface = (*TxManager)(nil)

func NewTxManager(pgPool *pgxpool.Pool, logger logger.LoggerInterface, txMarker string) *TxManager {
	return &TxManager{
		pgPool:   pgPool,
		logger:   logger,
		txMarker: txMarker,
	}
}

func (txM *TxManager) WithinTx(ctx context.Context, fn func(context.Context) error) error {
	tx, ok := ctx.Value(txM.txMarker).(pgx.Tx)
	var err error
	if ok {
		//Join other transaction
		err = fn(ctx)
		if err != nil {
			return err
		}
		return nil
	}

	// Begin new transaction
	tx, err = txM.pgPool.Begin(ctx)
	if err != nil {
		txM.logger.Error("failed to begin transaction", map[string]any{
			"error": err.Error(),
		})
		return err
	}

	txCtx := context.WithValue(ctx, txM.txMarker, tx)
	if err := fn(txCtx); err != nil {
		if errR := tx.Rollback(txCtx); errR != nil {
			txM.logger.Error("failed to rollback transaction", map[string]any{
				"error": errR.Error(),
			})
		}
		return err
	}

	if err := tx.Commit(txCtx); err != nil {
		txM.logger.Error("failed to commit transaction", map[string]any{
			"error": err.Error(),
		})
		return err
	}

	txM.logger.Info("successfully finished transaction", nil)
	return nil
}
