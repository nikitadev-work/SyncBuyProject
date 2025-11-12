package grpcserver

import (
	uc "purchase/internal/usecase"
	pb "purchase/proto-codegen"
	purchasepb "purchase/proto-codegen"

	"github.com/google/uuid"
)

func ConvertCreatePurchaseReqToInputDTO(in *pb.CreatePurchaseRequest) (*uc.CreatePurchaseInputDTO, error) {
	if err := CheckTitle(in.Title); err != nil {
		return nil, ErrIncorrectTitle
	}
	if err := CheckDescription(in.Description); err != nil {
		return nil, ErrIncorrectDescription
	}
	if err := CheckCurrency(in.Currency); err != nil {
		return nil, ErrIncorrectCurrency
	}
	return &uc.CreatePurchaseInputDTO{
		Title:       in.Title,
		Description: in.Description,
		Currency:    in.Currency,
	}, nil
}

func ConvertOutputToCreatePurchaseRespDTO(in *uc.CreatePurchaseOutputDTO) (*pb.CreatePurchaseResponse, error) {
	return &pb.CreatePurchaseResponse{
		PurchaseId: in.PurchaseId.String(),
	}, nil
}

func ConvertGetPurchaseReqToInputDTO(in *pb.GetPurchaseRequest) (*uc.GetPurchaseInputDTO, error) {
	id, err := uuid.Parse(in.PurchaseId)
	if err != nil {
		return nil, err
	}
	return &uc.GetPurchaseInputDTO{
		PurchaseId: id,
	}, nil
}

func ConvertOutputToGetPurchaseRespDTO(in *uc.GetPurchaseOutputDTO) (*pb.GetPurchaseResponse, error) {
	purchase := pb.Purchase{
		Title:       in.Purchase.Title,
		Description: in.Purchase.Description,
		Currency:    in.Purchase.Currency,
	}
	return &pb.GetPurchaseResponse{
		Purchase: &purchase,
	}, nil
}

func ConvertCreateInviteReqToInputDTO(in *pb.CreateInviteRequest) (*uc.CreateInviteInputDTO, error) {
	id, err := uuid.Parse(in.PurchaseId)
	if err != nil {
		return nil, ErrIncorrectInternalId
	}
	return &uc.CreateInviteInputDTO{
		PurchaseId: id,
	}, nil
}

func ConvertOutputToCreateInviteRespDTO(in *uc.CreateInviteOutputDTO) (*pb.CreateInviteResponse, error) {
	return &pb.CreateInviteResponse{
		Token:       in.Token,
		Title:       in.Title,
		Description: in.Description,
	}, nil
}

func ConvertJoinInviteReqToInputDTO(in *pb.JoinByInviteRequest) (*uc.JoinByInviteInputDTO, error) {
	if err := CheckToken(in.Token); err != nil {
		return nil, ErrIncorrectToken
	}
	userId, err := uuid.Parse(in.UserId)
	if err != nil {
		return nil, ErrIncorrectInternalId
	}
	return &uc.JoinByInviteInputDTO{
		UserId: userId,
		Token:  in.Token,
	}, nil
}

func ConvertRemoveParticipantReqToInputDTO(in *pb.RemoveParticipantRequest) (*uc.RemoveParticipantInputDTO, error) {
	userId, err := uuid.Parse(in.UserId)
	if err != nil {
		return nil, ErrIncorrectInternalId
	}
	purchaseId, err := uuid.Parse(in.PurchaseId)
	if err != nil {
		return nil, ErrIncorrectInternalId
	}
	return &uc.RemoveParticipantInputDTO{
		UserId:     userId,
		PurchaseId: purchaseId,
	}, nil
}

func ConvertListParticipantsReqToInputDTO(in *pb.ListParticipantsByPurchaseIdRequest) (*uc.ListParticipantsByPurchaseIdInputDTO, error) {
	purchaseId, err := uuid.Parse(in.PurchaseId)
	if err != nil {
		return nil, err
	}
	return &uc.ListParticipantsByPurchaseIdInputDTO{
		PurchaseId: purchaseId,
	}, nil
}

func ConvertOutputToListParticipantsRespDTO(in *uc.ListParticipantsByPurchaseIdOutputDTO) (*pb.ListParticipantsByPurchaseIdResponse, error) {
	purchaseIds := make([]string, 0, len(in.UserIds))
	for _, v := range in.UserIds {
		purchaseIds = append(purchaseIds, v.String())
	}
	return &pb.ListParticipantsByPurchaseIdResponse{
		UserIds: purchaseIds,
	}, nil
}

func ConvertCreateTaskReqToInputDTO(in *pb.CreateTaskRequest) (*uc.CreateTaskInputDTO, error) {
	purchaseId, err := uuid.Parse(in.PurchaseId)
	if err != nil {
		return nil, err
	}
	authorUserId, err := uuid.Parse(in.AuthorUserId)
	if err != nil {
		return nil, err
	}
	if err := CheckTitle(in.Title); err != nil {
		return nil, ErrIncorrectTitle
	}
	if err := CheckDescription(in.Description); err != nil {
		return nil, ErrIncorrectDescription
	}
	if err := CheckAmount(in.Amount); err != nil {
		return nil, ErrIncorrectAmount
	}
	return &uc.CreateTaskInputDTO{
		Title:        in.Title,
		Description:  in.Description,
		PurchaseId:   purchaseId,
		AuthorUserId: authorUserId,
		Amount:       in.Amount,
	}, nil
}

func ConvertOutputToCreateTaskRespDTO(in *uc.CreateTaskOutputDTO) (*pb.CreateTaskResponse, error) {
	return &pb.CreateTaskResponse{
		TaskId: in.TaskId.String(),
	}, nil
}

func ConvertTakeTaskReqToInputDTO(in *pb.TakeTaskRequest) (*uc.TakeTaskInputDTO, error) {
	taskId, err := uuid.Parse(in.TaskId)
	if err != nil {
		return nil, err
	}
	executorId, err := uuid.Parse(in.ExecutorUserId)
	if err != nil {
		return nil, err
	}
	return &uc.TakeTaskInputDTO{
		TaskId:         taskId,
		ExecutorUserId: executorId,
	}, nil
}

func ConvertDeleteTaskReqToInputDTO(in *pb.DeleteTaskRequest) (*uc.DeleteTaskInputDTO, error) {
	taskId, err := uuid.Parse(in.TaskId)
	if err != nil {
		return nil, err
	}
	return &uc.DeleteTaskInputDTO{
		TaskId: taskId,
	}, nil
}

func ConvertListTasksReqToInputDTO(in *pb.ListTasksByPurchaseIdRequest) (*uc.ListTasksByPurchaseIdInputDTO, error) {
	purchaseId, err := uuid.Parse(in.PurchaseId)
	if err != nil {
		return nil, err
	}
	return &uc.ListTasksByPurchaseIdInputDTO{
		PurchaseId: purchaseId,
	}, nil
}

func ConvertOutputToListTasksRespDTO(in *uc.ListTasksByPurchaseIdOutputDTO) (*pb.ListTasksByPurchaseIdResponse, error) {
	tasks := make([]*purchasepb.Task, 0, len(in.Tasks))
	for _, v := range in.Tasks {
		newTask := purchasepb.Task{
			Id:             v.Id.String(),
			Title:          v.Title,
			Description:    v.Description,
			PurchaseId:     v.PurchaseId.String(),
			AuthorUserId:   v.AuthorUserId.String(),
			ExecutorUserId: v.ExecutorUserId.String(),
			Done:           v.Done,
			Amount:         v.Amount,
		}
		tasks = append(tasks, &newTask)
	}
	return &pb.ListTasksByPurchaseIdResponse{
		Tasks: tasks,
	}, nil
}

func ConvertMarkTaskAsDoneReqToInputDTO(in *pb.MarkTaskAsDoneRequest) (*uc.MarkTaskAsDoneInputDTO, error) {
	taskId, err := uuid.Parse(in.TaskId)
	if err != nil {
		return nil, err
	}
	return &uc.MarkTaskAsDoneInputDTO{
		TaskId: taskId,
	}, nil
}

func ConvertLockPurchaseReqToInputDTO(in *pb.LockPurchaseRequest) (*uc.LockPurchaseInputDTO, error) {
	purchaseId, err := uuid.Parse(in.PurchaseId)
	if err != nil {
		return nil, err
	}
	return &uc.LockPurchaseInputDTO{
		PurchaseId: purchaseId,
	}, nil
}

func ConvertUnlockPurchaseReqToInputDTO(in *pb.UnlockPurchaseRequest) (*uc.UnlockPurchaseInputDTO, error) {
	purchaseId, err := uuid.Parse(in.PurchaseId)
	if err != nil {
		return nil, err
	}
	return &uc.UnlockPurchaseInputDTO{
		PurchaseId: purchaseId,
	}, nil
}

func ConvertGetSnapshotReqToInputDTO(in *pb.GetSnapshotRequest) (*uc.GetSnapshotInputDTO, error) {
	purchaseId, err := uuid.Parse(in.PurchaseId)
	if err != nil {
		return nil, err
	}
	return &uc.GetSnapshotInputDTO{
		PurchaseId: purchaseId,
	}, nil
}

func ConvertOutputToGetsnapshotRespDTO(in *uc.GetSnapshotOutputDTO) (*pb.GetSnapshotResponse, error) {
	userIds := make([]string, 0, len(in.Snapshot.ParticipantUserIds))
	for _, v := range in.Snapshot.ParticipantUserIds {
		userIds = append(userIds, v.String())
	}

	tasks := make([]*purchasepb.Task, 0, len(in.Snapshot.Tasks))
	for _, v := range in.Snapshot.Tasks {
		newTask := purchasepb.Task{
			Id:             v.Id.String(),
			Title:          v.Title,
			Description:    v.Description,
			PurchaseId:     v.PurchaseId.String(),
			AuthorUserId:   v.AuthorUserId.String(),
			ExecutorUserId: v.ExecutorUserId.String(),
			Done:           v.Done,
			Amount:         v.Amount,
		}
		tasks = append(tasks, &newTask)
	}
	snapshot := purchasepb.Snapshot{
		PurchaseId:         in.Snapshot.PurchaseId.String(),
		PurchaseTitle:      in.Snapshot.PurchaseTitle,
		ParticipantUserIds: userIds,
		Tasks:              tasks,
		Currency:           in.Snapshot.Currency,
	}

	return &pb.GetSnapshotResponse{
		Snapshot: &snapshot,
	}, nil
}

func ConvertMarkSettlementReqToInputDTO(in *pb.MarkSettlementInitiatedRequest) (*uc.MarkSettlementInitiatedInputDTO, error) {
	purchaseId, err := uuid.Parse(in.PurchaseId)
	if err != nil {
		return nil, err
	}
	return &uc.MarkSettlementInitiatedInputDTO{
		PurchaseId: purchaseId,
	}, nil
}

func ConvertFinishPurchaseReqToInputDTO(in *pb.FinishPurchaseRequest) (*uc.FinishPurchaseInputDTO, error) {
	purchaseId, err := uuid.Parse(in.PurchaseId)
	if err != nil {
		return nil, err
	}
	return &uc.FinishPurchaseInputDTO{
		PurchaseId: purchaseId,
	}, nil
}
