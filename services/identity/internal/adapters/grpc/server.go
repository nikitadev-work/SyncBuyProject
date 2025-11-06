package grpcserver

import (
	"context"
	uc "identity/internal/usecase"
	pb "identity/proto-codegen"

	"github.com/nikitadev-work/SyncBuyProject/common/kit/logger"
	"github.com/nikitadev-work/SyncBuyProject/common/kit/metrics"
)

type IdentityServer struct {
	pb.UnimplementedIdentityServiceServer
	usecase uc.IdentityUsecase
	logger  logger.LoggerInterface
}

func New(usecase uc.IdentityUsecase, logger logger.LoggerInterface) *IdentityServer {
	return &IdentityServer{
		usecase: usecase,
		logger:  logger,
	}
}

func (s *IdentityServer) RegisterOrGetTelegramUser(ctx context.Context, in *pb.RegisterOrGetTelegramUserRequest) (*pb.RegisterOrGetTelegramUserResponse, error) {
	defer func() {
		metrics.IncGRPCRequestsTotal("RegisterOrGetTelegramUser")
	}()

	err := ValidateRegisterOrGetTelegramUser(in)
	if err != nil {
		s.logger.Error("validation error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	meta, err := ConvertMetaToJson(in.Meta)
	if err != nil {
		s.logger.Error("meta convertation error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	reqDTO, err := ConvertRegOrGetUserRequestToInputDTO(in, meta)
	if err != nil {
		s.logger.Error("request data convertation error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	resp, err := s.usecase.RegisterOrGetUserByTelegram(ctx, *reqDTO)
	if err != nil {
		s.logger.Error("RegOrGetUserByTelegram error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	result, err := ConvertOutputDTOToRegOrGetUserResponse(&resp)
	if err != nil {
		s.logger.Error("response data convertation error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	return result, nil
}

func (s *IdentityServer) GetUserByTelegramId(ctx context.Context, in *pb.GetUserByTelegramIdRequest) (*pb.GetUserByTelegramIdResponse, error) {
	defer func() {
		metrics.IncGRPCRequestsTotal("GetUserByTelegramId")
	}()

	err := CheckExternalId(in.TelegramId)
	if err != nil {
		s.logger.Error("validation request error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	reqDTO, err := ConvertGetUserTelegramRequestToInputDTO(in)
	if err != nil {
		s.logger.Error("convertation request data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	resp, err := s.usecase.GetUserByTelegram(ctx, *reqDTO)
	if err != nil {
		s.logger.Error("GetUserByTelegram error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	result, err := ConvertOutputDTOToGetUserTelegramResonse(&resp)
	if err != nil {
		s.logger.Error("convertation response data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	return result, nil
}

func (s *IdentityServer) GetUserByUserId(ctx context.Context, in *pb.GetUserByUserIdRequest) (*pb.GetUserByUserIdResponse, error) {
	defer func() {
		metrics.IncGRPCRequestsTotal("GetUserByUserId")
	}()

	err := CheckInternalId(in.UserId)
	if err != nil {
		s.logger.Error("validation request error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	reqDTO, err := ConvertGetUserByUserIdToInputDTO(in)
	if err != nil {
		s.logger.Error("convertation request data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	resp, err := s.usecase.GetUserByUserId(ctx, *reqDTO)
	if err != nil {
		s.logger.Error("GetUserByUserId error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	result, err := ConvertOutputDTOToGetUserByUserIdResponse(&resp)
	if err != nil {
		s.logger.Error("convertation request data error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	return result, nil
}

func (s *IdentityServer) Health(ctx context.Context, _ *pb.HealthRequest) (*pb.HealthResponse, error) {
	return &pb.HealthResponse{
		Status: "OK",
	}, nil
}
