package grpcserver

import (
	"context"
	uc "purchase/internal/usecase"
	pb "purchase/proto-codegen"

	"github.com/nikitadev-work/SyncBuyProject/common/kit/logger"
	"github.com/nikitadev-work/SyncBuyProject/common/kit/metrics"
)

type PurchaseServer struct {
	pb.UnimplementedPurchaseServiceServer
	usecase uc.PurchaseUsecase
	logger  logger.LoggerInterface
}

func New(usecase uc.PurchaseUsecase, logger logger.LoggerInterface) *PurchaseServer {
	return &PurchaseServer{
		usecase: usecase,
		logger:  logger,
	}
}

func (s *PurchaseServer) CreatePurchase(ctx context.Context, in *pb.CreatePurchaseRequest) (*pb.CreatePurchaseResponse, error) {
	defer func() {
		metrics.IncGRPCRequestsTotal("CreatePurchase")
	}()

	reqDTO, err := ConvertCreatePurchaseReqToInputDTO(in)
	if err != nil {
		s.logger.Error("convertation request data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	resp, err := s.usecase.CreatePurchase(ctx, *reqDTO)
	if err != nil {
		s.logger.Error("CreatePurchase error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	result, err := ConvertOutputToCreatePurchaseRespDTO(resp)
	if err != nil {
		s.logger.Error("convertation response data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	return result, nil
}

func (s *PurchaseServer) GetPurchase(ctx context.Context, in *pb.GetPurchaseRequest) (*pb.GetPurchaseResponse, error) {
	defer func() {
		metrics.IncGRPCRequestsTotal("GetPurchase")
	}()

	reqDTO, err := ConvertGetPurchaseReqToInputDTO(in)
	if err != nil {
		s.logger.Error("convertation request data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	resp, err := s.usecase.GetPurchase(ctx, *reqDTO)
	if err != nil {
		s.logger.Error("GetPurchase error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	result, err := ConvertOutputToGetPurchaseRespDTO(resp)
	if err != nil {
		s.logger.Error("convertation response data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	return result, nil
}

// Participant
func (s *PurchaseServer) CreateInvite(ctx context.Context, in *pb.CreateInviteRequest) (*pb.CreateInviteResponse, error) {
	defer func() {
		metrics.IncGRPCRequestsTotal("CreateInvite")
	}()

	reqDTO, err := ConvertCreateInviteReqToInputDTO(in)
	if err != nil {
		s.logger.Error("convertation request data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	resp, err := s.usecase.CreateInvite(ctx, *reqDTO)
	if err != nil {
		s.logger.Error("CreateInvite error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	result, err := ConvertOutputToCreateInviteRespDTO(resp)
	if err != nil {
		s.logger.Error("convertation response data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	return result, nil
}

func (s *PurchaseServer) JoinByInvite(ctx context.Context, in *pb.JoinByInviteRequest) (*pb.JoinByInviteResponse, error) {
	defer func() {
		metrics.IncGRPCRequestsTotal("JoinByInvite")
	}()

	reqDTO, err := ConvertJoinInviteReqToInputDTO(in)
	if err != nil {
		s.logger.Error("convertation request data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	if err := s.usecase.JoinByInvite(ctx, *reqDTO); err != nil {
		s.logger.Error("JoinByInvite error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	return &pb.JoinByInviteResponse{}, nil
}

func (s *PurchaseServer) RemoveParticipant(ctx context.Context, in *pb.RemoveParticipantRequest) (*pb.RemoveParticipantResponse, error) {
	defer func() {
		metrics.IncGRPCRequestsTotal("RemoveParticipant")
	}()

	reqDTO, err := ConvertRemoveParticipantReqToInputDTO(in)
	if err != nil {
		s.logger.Error("convertation request data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	if err := s.usecase.RemoveParticipant(ctx, *reqDTO); err != nil {
		s.logger.Error("RemoveParticipant error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	return &pb.RemoveParticipantResponse{}, nil
}

func (s *PurchaseServer) ListParticipantsByPurchaseId(ctx context.Context, in *pb.ListParticipantsByPurchaseIdRequest) (*pb.ListParticipantsByPurchaseIdResponse, error) {
	defer func() {
		metrics.IncGRPCRequestsTotal("ListParticipantsByPurchaseId")
	}()

	reqDTO, err := ConvertListParticipantsReqToInputDTO(in)
	if err != nil {
		s.logger.Error("convertation request data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	resp, err := s.usecase.ListParticipantsByPurchaseId(ctx, *reqDTO)
	if err != nil {
		s.logger.Error("ListParticipantsByPurchaseId error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	result, err := ConvertOutputToListParticipantsRespDTO(resp)
	if err != nil {
		s.logger.Error("convertation response data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	return result, nil
}

// Task
func (s *PurchaseServer) CreateTask(ctx context.Context, in *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	defer func() {
		metrics.IncGRPCRequestsTotal("CreateTask")
	}()

	reqDTO, err := ConvertCreateTaskReqToInputDTO(in)
	if err != nil {
		s.logger.Error("convertation request data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	resp, err := s.usecase.CreateTask(ctx, *reqDTO)
	if err != nil {
		s.logger.Error("CreateTask error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	result, err := ConvertOutputToCreateTaskRespDTO(resp)
	if err != nil {
		s.logger.Error("convertation response data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	return result, nil
}

func (s *PurchaseServer) TakeTask(ctx context.Context, in *pb.TakeTaskRequest) (*pb.TakeTaskResponse, error) {
	defer func() {
		metrics.IncGRPCRequestsTotal("TakeTask")
	}()

	reqDTO, err := ConvertTakeTaskReqToInputDTO(in)
	if err != nil {
		s.logger.Error("convertation request data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	if err := s.usecase.TakeTask(ctx, *reqDTO); err != nil {
		s.logger.Error("TakeTask error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	return &pb.TakeTaskResponse{}, nil
}

func (s *PurchaseServer) DeleteTask(ctx context.Context, in *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	defer func() {
		metrics.IncGRPCRequestsTotal("DeleteTask")
	}()

	reqDTO, err := ConvertDeleteTaskReqToInputDTO(in)
	if err != nil {
		s.logger.Error("convertation request data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	if err := s.usecase.DeleteTask(ctx, *reqDTO); err != nil {
		s.logger.Error("DeleteTask error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	return &pb.DeleteTaskResponse{}, nil
}

func (s *PurchaseServer) ListTasksByPurchaseId(ctx context.Context, in *pb.ListTasksByPurchaseIdRequest) (*pb.ListTasksByPurchaseIdResponse, error) {
	defer func() {
		metrics.IncGRPCRequestsTotal("ListTasksByPurchaseId")
	}()

	reqDTO, err := ConvertListTasksReqToInputDTO(in)
	if err != nil {
		s.logger.Error("convertation request data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	resp, err := s.usecase.ListTasksByPurchaseId(ctx, *reqDTO)
	if err != nil {
		s.logger.Error("ListTasksByPurchaseId error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	result, err := ConvertOutputToListTasksRespDTO(resp)
	if err != nil {
		s.logger.Error("convertation response data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	return result, nil
}

func (s *PurchaseServer) MarkTaskAsDone(ctx context.Context, in *pb.MarkTaskAsDoneRequest) (*pb.MarkTaskAsDoneResponse, error) {
	defer func() {
		metrics.IncGRPCRequestsTotal("MarkTaskAsDone")
	}()

	reqDTO, err := ConvertMarkTaskAsDoneReqToInputDTO(in)
	if err != nil {
		s.logger.Error("convertation request data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	if err := s.usecase.MarkTaskAsDone(ctx, *reqDTO); err != nil {
		s.logger.Error("MarkTaskAsDone error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	return &pb.MarkTaskAsDoneResponse{}, nil
}

// Status
func (s *PurchaseServer) LockPurchase(ctx context.Context, in *pb.LockPurchaseRequest) (*pb.LockPurchaseResponse, error) {
	defer func() {
		metrics.IncGRPCRequestsTotal("LockPurchase")
	}()

	reqDTO, err := ConvertLockPurchaseReqToInputDTO(in)
	if err != nil {
		s.logger.Error("convertation request data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	if err := s.usecase.LockPurchase(ctx, *reqDTO); err != nil {
		s.logger.Error("LockPurchase error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	return &pb.LockPurchaseResponse{}, nil
}

func (s *PurchaseServer) UnlockPurchase(ctx context.Context, in *pb.UnlockPurchaseRequest) (*pb.UnlockPurchaseResponse, error) {
	defer func() {
		metrics.IncGRPCRequestsTotal("UnlockPurchase")
	}()

	reqDTO, err := ConvertUnlockPurchaseReqToInputDTO(in)
	if err != nil {
		s.logger.Error("convertation request data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	if err := s.usecase.UnlockPurchase(ctx, *reqDTO); err != nil {
		s.logger.Error("UnlockPurchase error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	return &pb.UnlockPurchaseResponse{}, nil
}

func (s *PurchaseServer) GetSnapshot(ctx context.Context, in *pb.GetSnapshotRequest) (*pb.GetSnapshotResponse, error) {
	defer func() {
		metrics.IncGRPCRequestsTotal("GetSnapshot")
	}()

	reqDTO, err := ConvertGetSnapshotReqToInputDTO(in)
	if err != nil {
		s.logger.Error("convertation request data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	resp, err := s.usecase.GetSnapshot(ctx, *reqDTO)
	if err != nil {
		s.logger.Error("GetSnapshot error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	result, err := ConvertOutputToGetsnapshotRespDTO(resp)
	if err != nil {
		s.logger.Error("convertation response data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	return result, nil
}

func (s *PurchaseServer) MarkSettlementInitiated(ctx context.Context, in *pb.MarkSettlementInitiatedRequest) (*pb.MarkSettlementInitiatedResponse, error) {
	defer func() {
		metrics.IncGRPCRequestsTotal("MarkSettlementInitiated")
	}()

	reqDTO, err := ConvertMarkSettlementReqToInputDTO(in)
	if err != nil {
		s.logger.Error("convertation request data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	if err := s.usecase.MarkSettlementInitiated(ctx, *reqDTO); err != nil {
		s.logger.Error("MarkSettlementInitiated error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	return &pb.MarkSettlementInitiatedResponse{}, nil
}

func (s *PurchaseServer) FinishPurchase(ctx context.Context, in *pb.FinishPurchaseRequest) (*pb.FinishPurchaseResponse, error) {
	defer func() {
		metrics.IncGRPCRequestsTotal("FinishPurchase")
	}()

	reqDTO, err := ConvertFinishPurchaseReqToInputDTO(in)
	if err != nil {
		s.logger.Error("convertation request data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	if err := s.usecase.FinishPurchase(ctx, *reqDTO); err != nil {
		s.logger.Error("FinishPurchase error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	return &pb.FinishPurchaseResponse{}, nil
}

func (s *PurchaseServer) Health(ctx context.Context, _ *pb.HealthRequest) (*pb.HealthResponse, error) {
	return &pb.HealthResponse{
		Status: "OK",
	}, nil
}
