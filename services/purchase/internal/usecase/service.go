package usecase

import (
	"context"
	"purchase/internal/domain"
	"time"
)

type PurchaseUsecaseInterface interface {
	// Purchase
	CreatePurchase(context.Context, CreatePurchaseInputDTO) (*CreatePurchaseOutputDTO, error)
	GetPurchase(context.Context, GetPurchaseInputDTO) (*GetPurchaseOutputDTO, error)

	// Participant
	CreateInvite(context.Context, CreateInviteInputDTO) (*CreateInviteOutputDTO, error)
	JoinByInvite(context.Context, JoinByInviteInputDTO) error
	RemoveParticipant(context.Context, RemoveParticipantInputDTO) error
	ListParticipantsByPurchaseId(context.Context, ListParticipantsByPurchaseIdInputDTO) (*ListParticipantsByPurchaseIdOutputDTO, error)

	// Task
	CreateTask(context.Context, CreateTaskInputDTO) (*CreateTaskOutputDTO, error)
	TakeTask(context.Context, TakeTaskInputDTO) error
	DeleteTask(context.Context, DeleteTaskInputDTO) error
	ListTasksByPurchaseId(context.Context, ListTasksByPurchaseIdInputDTO) (*ListTasksByPurchaseIdOutputDTO, error)
	MarkTaskAsDone(context.Context, MarkTaskAsDoneInputDTO) error

	// Status
	LockPurchase(context.Context, LockPurchaseInputDTO) error
	UnlockPurchase(context.Context, UnlockPurchaseInputDTO) error
	GetSnapshot(context.Context, GetSnapshotInputDTO) (*GetSnapshotOutputDTO, error)
	MarkSettlementInitiated(context.Context, MarkSettlementInitiatedInputDTO) error
	FinishPurchase(context.Context, FinishPurchaseInputDTO) error
}

type PurchaseUsecase struct {
	Repository    PurchaseRepositoryInterface
	TxManager     TxManagerInterface
	TokenProvider TokenProviderInterface
}

var _ PurchaseUsecaseInterface = (*PurchaseUsecase)(nil)

func NewPurchaseUsecase(repository PurchaseRepositoryInterface, txManager TxManagerInterface,
	tokenProvider TokenProviderInterface) *PurchaseUsecase {
	return &PurchaseUsecase{
		Repository:    repository,
		TxManager:     txManager,
		TokenProvider: tokenProvider,
	}
}

// Purchase
func (uc *PurchaseUsecase) CreatePurchase(ctx context.Context, in CreatePurchaseInputDTO) (*CreatePurchaseOutputDTO, error) {
	reqDTO := CreatePurchaseRequestDTO{
		Title:       in.Title,
		Description: in.Description,
		Currency:    in.Currency,
	}

	resp, err := uc.Repository.CreatePurchase(ctx, &reqDTO)
	if err != nil {
		return nil, err
	}

	return &CreatePurchaseOutputDTO{
		PurchaseId: resp.PurchaseId,
	}, nil
}

func (uc *PurchaseUsecase) GetPurchase(ctx context.Context, in GetPurchaseInputDTO) (*GetPurchaseOutputDTO, error) {
	reqDTO := GetPurchaseRequestDTO{
		PurchaseId: in.PurchaseId,
	}

	resp, err := uc.Repository.GetPurchase(ctx, &reqDTO)
	if err != nil {
		return nil, err
	}

	return &GetPurchaseOutputDTO{
		Purchase: resp.Purchase,
	}, nil
}

// Participant
func (uc *PurchaseUsecase) CreateInvite(ctx context.Context, in CreateInviteInputDTO) (*CreateInviteOutputDTO, error) {
	expiresAt := time.Now().Add(168 * time.Hour)
	reqGenTokenDTO := GenerateTokenRequest{
		PurchaseId: in.PurchaseId,
		ExpiresAt:  expiresAt,
	}
	res, err := uc.TokenProvider.GenerateInviteToken(&reqGenTokenDTO)
	if err != nil {
		return nil, err
	}

	reqDTO := GetPurchaseRequestDTO{
		PurchaseId: in.PurchaseId,
	}

	resp, err := uc.Repository.GetPurchase(ctx, &reqDTO)
	if err != nil {
		return nil, err
	}

	return &CreateInviteOutputDTO{
		Token:       res.Token,
		Title:       resp.Purchase.Title,
		Description: resp.Purchase.Description,
	}, nil
}

func (uc *PurchaseUsecase) JoinByInvite(ctx context.Context, in JoinByInviteInputDTO) error {
	parseTokenDTO := ParseTokenRequest{
		TokenString: in.Token,
	}
	res, err := uc.TokenProvider.ParseAndValidateInviteToken(&parseTokenDTO)
	if err != nil {
		return err
	}

	reqDTO := AddParticipantToPurchaseRequestDTO{
		UserId:     in.UserId,
		PurchaseId: res.PurchaseId,
	}

	if err := uc.Repository.AddParticipantToPurchase(ctx, &reqDTO); err != nil {
		return err
	}
	return nil
}

func (uc *PurchaseUsecase) RemoveParticipant(ctx context.Context, in RemoveParticipantInputDTO) error {
	reqDTO := RemoveParticipantRequestDTO{
		UserId:     in.UserId,
		PurchaseId: in.PurchaseId,
	}
	err := uc.Repository.RemoveParticipant(ctx, &reqDTO)
	return err
}

func (uc *PurchaseUsecase) ListParticipantsByPurchaseId(ctx context.Context, in ListParticipantsByPurchaseIdInputDTO) (*ListParticipantsByPurchaseIdOutputDTO, error) {
	reqDTO := ListParticipantsByPurchaseIdRequestDTO{
		PurchaseId: in.PurchaseId,
	}

	resp, err := uc.Repository.ListParticipantsByPurchaseId(ctx, &reqDTO)
	if err != nil {
		return nil, err
	}

	return &ListParticipantsByPurchaseIdOutputDTO{
		UserIds: resp.UserIds,
	}, nil
}

// Task
func (uc *PurchaseUsecase) CreateTask(ctx context.Context, in CreateTaskInputDTO) (*CreateTaskOutputDTO, error) {
	reqDTO := CreateTaskRequestDTO{
		Title:        in.Title,
		Description:  in.Description,
		PurchaseId:   in.PurchaseId,
		AuthorUserId: in.AuthorUserId,
		Amount:       in.Amount,
	}

	resp, err := uc.Repository.CreateTask(ctx, &reqDTO)
	if err != nil {
		return nil, err
	}

	return &CreateTaskOutputDTO{
		TaskId: resp.TaskId,
	}, nil
}

func (uc *PurchaseUsecase) TakeTask(ctx context.Context, in TakeTaskInputDTO) error {
	reqDTO := TakeTaskRequestDTO{
		TaskId: in.TaskId,
		UserId: in.ExecutorUserId,
	}
	err := uc.Repository.TakeTask(ctx, &reqDTO)
	return err
}

func (uc *PurchaseUsecase) DeleteTask(ctx context.Context, in DeleteTaskInputDTO) error {
	reqDTO := DeleteTaskRequestDTO{
		TaskId: in.TaskId,
	}
	err := uc.Repository.DeleteTask(ctx, &reqDTO)
	return err
}

func (uc *PurchaseUsecase) ListTasksByPurchaseId(ctx context.Context, in ListTasksByPurchaseIdInputDTO) (*ListTasksByPurchaseIdOutputDTO, error) {
	reqDTO := ListTasksByPurchaseIdRequestDTO{
		PurchaseId: in.PurchaseId,
	}

	resp, err := uc.Repository.ListTasksByPurchaseId(ctx, &reqDTO)
	if err != nil {
		return nil, err
	}

	return &ListTasksByPurchaseIdOutputDTO{
		Tasks: resp.Tasks,
	}, nil
}

func (uc *PurchaseUsecase) MarkTaskAsDone(ctx context.Context, in MarkTaskAsDoneInputDTO) error {
	reqDTO := MarkTaskAsDoneRequestDTO{
		TaskId: in.TaskId,
	}
	err := uc.Repository.MarkTaskAsDone(ctx, &reqDTO)
	return err
}

// Status
func (uc *PurchaseUsecase) LockPurchase(ctx context.Context, in LockPurchaseInputDTO) error {
	reqDTO := LockPurchaseRequestDTO{
		PurchaseId: in.PurchaseId,
	}
	err := uc.Repository.LockPurchase(ctx, &reqDTO)
	return err
}

func (uc *PurchaseUsecase) UnlockPurchase(ctx context.Context, in UnlockPurchaseInputDTO) error {
	reqDTO := UnlockPurchaseRequestDTO{
		PurchaseId: in.PurchaseId,
	}
	err := uc.Repository.UnlockPurchase(ctx, &reqDTO)
	return err
}

func (uc *PurchaseUsecase) GetSnapshot(ctx context.Context, in GetSnapshotInputDTO) (*GetSnapshotOutputDTO, error) {
	snapshot := domain.Snapshot{
		PurchaseId: in.PurchaseId,
	}

	err := uc.TxManager.WithinTx(ctx, func(ctx context.Context) error {
		// Purchase
		reqDTO := GetPurchaseRequestDTO{
			PurchaseId: in.PurchaseId,
		}
		resp, errIn := uc.Repository.GetPurchase(ctx, &reqDTO)
		if errIn != nil {
			return errIn
		}
		if resp.Purchase.LockedAt != nil {
			return ErrGetSnapshot
		}
		snapshot.PurchaseTitle = resp.Purchase.Title
		snapshot.Currency = resp.Purchase.Currency

		// Participants
		reqParticDTO := ListParticipantsByPurchaseIdRequestDTO{
			PurchaseId: in.PurchaseId,
		}
		respPartic, errIn := uc.Repository.ListParticipantsByPurchaseId(ctx, &reqParticDTO)
		if errIn != nil {
			return errIn
		}
		snapshot.ParticipantUserIds = respPartic.UserIds

		// Tasks
		reqTasksDTO := ListTasksByPurchaseIdRequestDTO{
			PurchaseId: in.PurchaseId,
		}
		respTasks, errIn := uc.Repository.ListTasksByPurchaseId(ctx, &reqTasksDTO)
		if errIn != nil {
			return errIn
		}
		snapshot.Tasks = respTasks.Tasks

		return nil
	})

	if err != nil {
		return nil, err
	}
	return &GetSnapshotOutputDTO{
		Snapshot: snapshot,
	}, nil
}

func (uc *PurchaseUsecase) MarkSettlementInitiated(ctx context.Context, in MarkSettlementInitiatedInputDTO) error {
	reqDTO := MarkSettlementInitiatedRequestDTO{
		PurchaseId: in.PurchaseId,
	}
	err := uc.Repository.MarkSettlementInitiated(ctx, &reqDTO)
	return err
}

func (uc *PurchaseUsecase) FinishPurchase(ctx context.Context, in FinishPurchaseInputDTO) error {
	reqDTO := FinishPurchaseRequestDTO{
		PurchaseId: in.PurchaseId,
	}
	err := uc.Repository.FinishPurchase(ctx, &reqDTO)
	return err
}
