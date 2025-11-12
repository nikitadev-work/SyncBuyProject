package repository

import (
	"context"
	"purchase/internal/domain"
	uc "purchase/internal/usecase"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pgPool   *pgxpool.Pool
	txMarker string
}

var _ uc.PurchaseRepositoryInterface = (*Repository)(nil)

func NewRepository(pgPool *pgxpool.Pool, txMarker string) *Repository {
	return &Repository{
		pgPool:   pgPool,
		txMarker: txMarker,
	}
}

// Purchase
func (repo *Repository) CreatePurchase(ctx context.Context,
	in *uc.CreatePurchaseRequestDTO) (*uc.CreatePurchaseReponseDTO, error) {
	tx, ok := ctx.Value(repo.txMarker).(pgx.Tx)
	purchaseId := uuid.New()
	reqStr := "INSERT INTO purchases (id, title, description, currency)" +
		" VALUES ($1, $2, $3, $4)"
	var err error

	if ok {
		// Command within transaction
		if _, err = tx.Exec(ctx, reqStr, purchaseId, in.Title, in.Description, in.Currency); err != nil {
			return nil, err
		}
	} else {
		// Command without transaction
		if _, err := repo.pgPool.Exec(ctx, reqStr, purchaseId, in.Title, in.Description, in.Currency); err != nil {
			return nil, err
		}
	}

	return &uc.CreatePurchaseReponseDTO{
		PurchaseId: purchaseId,
	}, nil
}

func (repo *Repository) GetPurchase(ctx context.Context,
	input *uc.GetPurchaseRequestDTO) (*uc.GetPurchaseResponseDTO, error) {
	tx, ok := ctx.Value(repo.txMarker).(pgx.Tx)
	reqStr := "SELECT id, title, description, currency," +
		" settlement_initiated_at, status, locked_at, finished_at" +
		" FROM purchases WHERE id = $1"
	var res pgx.Row
	var err error

	if ok {
		res = tx.QueryRow(ctx, reqStr, input.PurchaseId)
	} else {
		res = repo.pgPool.QueryRow(ctx, reqStr, input.PurchaseId)
	}

	var prch domain.Purchase
	err = res.Scan(&prch.Id, &prch.Title, &prch.Description,
		&prch.Currency, &prch.SettlementInitiatedAt,
		&prch.Status, &prch.LockedAt, &prch.FinishedAt)
	if err != nil {
		return nil, ErrParsingResponse
	}

	return &uc.GetPurchaseResponseDTO{
		Purchase: prch,
	}, nil
}

// Participant
func (repo *Repository) AddParticipantToPurchase(ctx context.Context,
	input *uc.AddParticipantToPurchaseRequestDTO) error {
	tx, ok := ctx.Value(repo.txMarker).(pgx.Tx)
	reqStr := "INSERT INTO participants (user_id, purchase_id) VALUES ($1, $2)"
	reqCheckStr := "SELECT locked_at FROM purchases WHERE id = $1"
	var lockedAt *time.Time

	if ok {
		// Check if the purchase is unlocked
		res := tx.QueryRow(ctx, reqCheckStr, input.PurchaseId)
		if err := res.Scan(&lockedAt); err != nil {
			return err
		}
		if lockedAt != nil {
			return ErrEditLockedPurchase
		}

		// Edit purchase
		if _, err := tx.Exec(ctx, reqStr, input.UserId, input.PurchaseId); err != nil {
			return err
		}
	} else {
		// Check if the purchase is unlocked
		res := repo.pgPool.QueryRow(ctx, reqCheckStr, input.PurchaseId)
		if err := res.Scan(&lockedAt); err != nil {
			return err
		}
		if lockedAt != nil {
			return ErrEditLockedPurchase
		}

		// Edit purchase
		if _, err := repo.pgPool.Exec(ctx, reqStr, input.UserId, input.PurchaseId); err != nil {
			return err
		}
	}

	return nil
}

func (repo *Repository) RemoveParticipant(ctx context.Context,
	input *uc.RemoveParticipantRequestDTO) error {
	tx, ok := ctx.Value(repo.txMarker).(pgx.Tx)
	reqCheckStr := "SELECT locked_at FROM purchases WHERE id = $1"
	reqStr := "DELETE FROM participants WHERE user_id = $1 AND purchase_id = $2"
	var lockedAt *time.Time

	if ok {
		res := tx.QueryRow(ctx, reqCheckStr, input.PurchaseId)
		if err := res.Scan(&lockedAt); err != nil {
			return err
		}
		if lockedAt != nil {
			return ErrEditLockedPurchase
		}
		if _, err := tx.Exec(ctx, reqStr, input.UserId, input.PurchaseId); err != nil {
			return nil
		}
	} else {
		res := repo.pgPool.QueryRow(ctx, reqCheckStr, input.PurchaseId)
		if err := res.Scan(&lockedAt); err != nil {
			return err
		}
		if lockedAt != nil {
			return ErrEditLockedPurchase
		}

		if _, err := repo.pgPool.Exec(ctx, reqStr, input.UserId, input.PurchaseId); err != nil {
			return nil
		}
	}

	return nil
}

func (repo *Repository) ListParticipantsByPurchaseId(ctx context.Context,
	input *uc.ListParticipantsByPurchaseIdRequestDTO) (*uc.ListParticipantsByPurchaseIdResponseDTO, error) {
	tx, ok := ctx.Value(repo.txMarker).(pgx.Tx)
	reqStr := "SELECT id FROM participants WHERE purchase_id = $1"
	userIds := make([]uuid.UUID, 0, 10)
	var rows pgx.Rows
	var err error

	if ok {
		rows, err = tx.Query(ctx, reqStr, input.PurchaseId)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var userId uuid.UUID
			if err := rows.Scan(&userId); err != nil {
				return nil, err
			}
			userIds = append(userIds, userId)
		}

		if err := rows.Err(); err != nil {
			return nil, err
		}
	} else {
		rows, err = repo.pgPool.Query(ctx, reqStr, input.PurchaseId)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var userId uuid.UUID
			if err := rows.Scan(&userId); err != nil {
				return nil, err
			}
			userIds = append(userIds, userId)
		}

		if err := rows.Err(); err != nil {
			return nil, err
		}
	}

	return &uc.ListParticipantsByPurchaseIdResponseDTO{
		UserIds: userIds,
	}, nil
}

// Task
func (repo *Repository) CreateTask(ctx context.Context,
	input *uc.CreateTaskRequestDTO) (*uc.CreateTaskResponseDTO, error) {
	tx, ok := ctx.Value(repo.txMarker).(pgx.Tx)
	reqCheckStr := "SELECT locked_at FROM purchases WHERE id = $1"
	reqStr := "INSERT INTO tasks (id, title, description, purchaseId, authorUserId, amount)" +
		" VALUES ($1, $2, $3, $4, $5, $6)"
	taskId := uuid.New()
	var lockedAt *time.Time

	if ok {
		res := tx.QueryRow(ctx, reqCheckStr, input.PurchaseId)
		if err := res.Scan(&lockedAt); err != nil {
			return nil, err
		}
		if lockedAt != nil {
			return nil, ErrEditLockedPurchase
		}

		_, err := tx.Exec(ctx, reqStr, taskId, input.Title, input.Description, input.PurchaseId,
			input.AuthorUserId, input.Amount)
		if err != nil {
			return nil, err
		}
	} else {
		res := repo.pgPool.QueryRow(ctx, reqCheckStr, input.PurchaseId)
		if err := res.Scan(&lockedAt); err != nil {
			return nil, err
		}
		if lockedAt != nil {
			return nil, ErrEditLockedPurchase
		}

		_, err := repo.pgPool.Exec(ctx, reqStr, taskId, input.Title, input.Description, input.PurchaseId,
			input.AuthorUserId, input.Amount)
		if err != nil {
			return nil, err
		}
	}

	return &uc.CreateTaskResponseDTO{
		TaskId: taskId,
	}, nil
}

func (repo *Repository) TakeTask(ctx context.Context, input *uc.TakeTaskRequestDTO) error {
	tx, ok := ctx.Value(repo.txMarker).(pgx.Tx)
	reqStr := "UPDATE tasks SET executor_user_id = $1 WHERE id = $2"
	reqTaskStr := "SELECT purchase_id FROM tasks WHERE id = $1"
	reqCheckStr := "SELECT locked_at FROM purchases WHERE id = $1"
	var lockedAt *time.Time
	var purchaseId uuid.UUID

	if ok {
		// Get PurchaseId by TaskId
		row := tx.QueryRow(ctx, reqTaskStr, input.TaskId)
		if err := row.Scan(&purchaseId); err != nil {
			return err
		}

		res := tx.QueryRow(ctx, reqCheckStr, purchaseId)
		if err := res.Scan(&lockedAt); err != nil {
			return err
		}
		if lockedAt != nil {
			return ErrEditLockedPurchase
		}

		if _, err := tx.Exec(ctx, reqStr, input.UserId, input.TaskId); err != nil {
			return err
		}
	} else {
		// Get PurchaseId by TaskId
		row := repo.pgPool.QueryRow(ctx, reqTaskStr, input.TaskId)
		if err := row.Scan(&purchaseId); err != nil {
			return err
		}

		res := repo.pgPool.QueryRow(ctx, reqCheckStr, purchaseId)
		if err := res.Scan(&lockedAt); err != nil {
			return err
		}
		if lockedAt != nil {
			return ErrEditLockedPurchase
		}

		if _, err := repo.pgPool.Exec(ctx, reqStr, input.UserId, input.TaskId); err != nil {
			return err
		}
	}

	return nil
}

func (repo *Repository) DeleteTask(ctx context.Context, input *uc.DeleteTaskRequestDTO) error {
	tx, ok := ctx.Value(repo.txMarker).(pgx.Tx)
	reqStr := "DELETE FROM tasks WHERE id = $1"
	reqTaskStr := "SELECT purchase_id FROM tasks WHERE id = $1"
	reqCheckStr := "SELECT locked_at FROM purchases WHERE id = $1"
	var lockedAt *time.Time
	var purchaseId uuid.UUID

	if ok {
		// Get PurchaseId by TaskId
		row := tx.QueryRow(ctx, reqTaskStr, input.TaskId)
		if err := row.Scan(&purchaseId); err != nil {
			return err
		}

		res := tx.QueryRow(ctx, reqCheckStr, purchaseId)
		if err := res.Scan(&lockedAt); err != nil {
			return err
		}
		if lockedAt != nil {
			return ErrEditLockedPurchase
		}

		if _, err := tx.Exec(ctx, reqStr, input.TaskId); err != nil {
			return nil
		}
	} else {
		// Get PurchaseId by TaskId
		row := repo.pgPool.QueryRow(ctx, reqTaskStr, input.TaskId)
		if err := row.Scan(&purchaseId); err != nil {
			return err
		}

		res := repo.pgPool.QueryRow(ctx, reqCheckStr, purchaseId)
		if err := res.Scan(&lockedAt); err != nil {
			return err
		}
		if lockedAt != nil {
			return ErrEditLockedPurchase
		}

		if _, err := repo.pgPool.Exec(ctx, reqStr, input.TaskId); err != nil {
			return nil
		}
	}

	return nil
}

func (repo *Repository) ListTasksByPurchaseId(ctx context.Context,
	input *uc.ListTasksByPurchaseIdRequestDTO) (*uc.ListTasksByPurchaseIdResponseDTO, error) {
	tx, ok := ctx.Value(repo.txMarker).(pgx.Tx)
	reqStr := "SELECT id FROM tasks WHERE purchase_id = $1"
	tasks := make([]domain.Task, 0, 10)
	var rows pgx.Rows
	var err error

	if ok {
		rows, err = tx.Query(ctx, reqStr, input.PurchaseId)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var task domain.Task
			err := rows.Scan(&task.Id, &task.Title, &task.Description,
				&task.AuthorUserId, &task.ExecutorUserId, &task.Done, &task.Amount)
			if err != nil {
				return nil, err
			}
			tasks = append(tasks, task)
		}

		if err := rows.Err(); err != nil {
			return nil, err
		}
	} else {
		rows, err = repo.pgPool.Query(ctx, reqStr, input.PurchaseId)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var task domain.Task
			err := rows.Scan(&task.Id, &task.Title, &task.Description,
				&task.AuthorUserId, &task.ExecutorUserId, &task.Done, &task.Amount)
			if err != nil {
				return nil, err
			}
			tasks = append(tasks, task)
		}

		if err := rows.Err(); err != nil {
			return nil, err
		}
	}

	return &uc.ListTasksByPurchaseIdResponseDTO{
		Tasks: tasks,
	}, nil
}

func (repo *Repository) MarkTaskAsDone(ctx context.Context, input *uc.MarkTaskAsDoneRequestDTO) error {
	tx, ok := ctx.Value(repo.txMarker).(pgx.Tx)
	reqStr := "UPDATE tasks SET done = true WHERE id = $1"
	reqTaskStr := "SELECT purchase_id FROM tasks WHERE id = $1"
	reqCheckStr := "SELECT locked_at FROM purchases WHERE id = $1"
	var lockedAt *time.Time
	var purchaseId uuid.UUID

	if ok {
		// Get PurchaseId by TaskId
		row := tx.QueryRow(ctx, reqTaskStr, input.TaskId)
		if err := row.Scan(&purchaseId); err != nil {
			return err
		}

		res := tx.QueryRow(ctx, reqCheckStr, purchaseId)
		if err := res.Scan(&lockedAt); err != nil {
			return err
		}
		if lockedAt != nil {
			return ErrEditLockedPurchase
		}

		if _, err := tx.Exec(ctx, reqStr, input.TaskId); err != nil {
			return err
		}
	} else {
		// Get PurchaseId by TaskId
		row := repo.pgPool.QueryRow(ctx, reqTaskStr, input.TaskId)
		if err := row.Scan(&purchaseId); err != nil {
			return err
		}

		res := repo.pgPool.QueryRow(ctx, reqCheckStr, purchaseId)
		if err := res.Scan(&lockedAt); err != nil {
			return err
		}
		if lockedAt != nil {
			return ErrEditLockedPurchase
		}

		if _, err := repo.pgPool.Exec(ctx, reqStr, input.TaskId); err != nil {
			return err
		}
	}

	return nil
}

// Status
func (repo *Repository) LockPurchase(ctx context.Context, input *uc.LockPurchaseRequestDTO) error {
	tx, ok := ctx.Value(repo.txMarker).(pgx.Tx)
	reqStr := "UPDATE purchases SET locked_at = CURRENT_TIMESTAMP WHERE id = $1"
	reqCheckStr := "SELECT locked_at FROM purchases WHERE id = $1"
	var lockedAt *time.Time

	if ok {
		res := tx.QueryRow(ctx, reqCheckStr, input.PurchaseId)
		if err := res.Scan(&lockedAt); err != nil {
			return err
		}
		if lockedAt != nil {
			return ErrLockPurchase
		}

		if _, err := tx.Exec(ctx, reqStr, input.PurchaseId); err != nil {
			return err
		}
	} else {
		res := repo.pgPool.QueryRow(ctx, reqCheckStr, input.PurchaseId)
		if err := res.Scan(&lockedAt); err != nil {
			return err
		}
		if lockedAt != nil {
			return ErrLockPurchase
		}

		if _, err := repo.pgPool.Exec(ctx, reqStr, input.PurchaseId); err != nil {
			return err
		}
	}

	return nil
}

func (repo *Repository) UnlockPurchase(ctx context.Context, input *uc.UnlockPurchaseRequestDTO) error {
	tx, ok := ctx.Value(repo.txMarker).(pgx.Tx)
	reqCheckStr := "SELECT settlement_initiated_at FROM purchases WHERE id = $1"
	reqStr := "UPDATE purchases SET locked_at = NULL WHERE id = $1"
	var settlementInitiatedAt *time.Time

	if ok {
		res := tx.QueryRow(ctx, reqCheckStr, input.PurchaseId)
		if err := res.Scan(&settlementInitiatedAt); err != nil {
			return err
		}
		if settlementInitiatedAt != nil {
			return ErrUnlockPurchase
		}
		if _, err := tx.Exec(ctx, reqStr, input.PurchaseId); err != nil {
			return err
		}
	} else {
		res := repo.pgPool.QueryRow(ctx, reqStr, input.PurchaseId)
		if err := res.Scan(&settlementInitiatedAt); err != nil {
			return err
		}
		if settlementInitiatedAt == nil {
			return ErrUnlockPurchase
		}
		if _, err := repo.pgPool.Exec(ctx, reqStr, input.PurchaseId); err != nil {
			return err
		}
	}

	return nil
}

func (repo *Repository) MarkSettlementInitiated(ctx context.Context,
	input *uc.MarkSettlementInitiatedRequestDTO) error {
	tx, ok := ctx.Value(repo.txMarker).(pgx.Tx)
	reqStr := "UPDATE purchases SET settlement_initiated_at = CURRENT_TIMESTAMP WHERE id = $1"

	if ok {
		if _, err := tx.Exec(ctx, reqStr, input.PurchaseId); err != nil {
			return err
		}
	} else {
		if _, err := repo.pgPool.Exec(ctx, reqStr, input.PurchaseId); err != nil {
			return err
		}
	}

	return nil
}


func (repo *Repository) FinishPurchase(ctx context.Context, input *uc.FinishPurchaseRequestDTO) error {
	tx, ok := ctx.Value(repo.txMarker).(pgx.Tx)
	reqCheckStr := "SELECT locked_at FROM purchases WHERE id = $1"
	reqStr := "UPDATE purchases SET finished_at = CURRENT_TIMESTAMP WHERE id = $1"
	var lockedAt *time.Time

	if ok {
		res := tx.QueryRow(ctx, reqCheckStr, input.PurchaseId)
		if err := res.Scan(&lockedAt); err != nil {
			return err
		}
		if lockedAt == nil {
			return ErrUnlockPurchase
		}
		if _, err := tx.Exec(ctx, reqStr, input.PurchaseId); err != nil {
			return err
		}
	} else {
		if _, err := repo.pgPool.Exec(ctx, reqStr, input.PurchaseId); err != nil {
			return err
		}
	}

	return nil
}
