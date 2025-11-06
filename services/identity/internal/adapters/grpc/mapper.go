package grpcserver

import (
	"encoding/json"
	uc "identity/internal/usecase"
	pb "identity/proto-codegen"

	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

func ConvertRegOrGetUserRequestToInputDTO(input *pb.RegisterOrGetTelegramUserRequest, meta json.RawMessage) (*uc.RegisterOrGetUserByTelegramInputDTO, error) {
	return &uc.RegisterOrGetUserByTelegramInputDTO{
		TelegramId: input.TelegramId,
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		ChatId:     input.ChatId,
		Meta:       meta,
	}, nil
}

func ConvertOutputDTOToRegOrGetUserResponse(input *uc.RegisterOrGetUserByTelegramOutputDTO) (*pb.RegisterOrGetTelegramUserResponse, error) {
	return &pb.RegisterOrGetTelegramUserResponse{
		UserId: input.UserId.String(),
	}, nil
}

func ConvertGetUserTelegramRequestToInputDTO(input *pb.GetUserByTelegramIdRequest) (*uc.GetUserByTelegramInputDTO, error) {
	return &uc.GetUserByTelegramInputDTO{
		TelegramId: input.TelegramId,
	}, nil
}

func ConvertOutputDTOToGetUserTelegramResonse(input *uc.GetUserByTelegramOutputDTO) (*pb.GetUserByTelegramIdResponse, error) {
	profile := pb.Profile{
		UserId:    input.User.Id.String(),
		FirstName: input.User.FirstName,
		LastName:  input.User.LastName,
	}

	return &pb.GetUserByTelegramIdResponse{
		Profile: &profile,
	}, nil
}

func ConvertGetUserByUserIdToInputDTO(input *pb.GetUserByUserIdRequest) (*uc.GetUserByUserIdInputDTO, error) {
	id, err := ConvertStringToInternalId(input.UserId)
	if err != nil {
		return nil, ErrIncorrectInternalId
	}

	return &uc.GetUserByUserIdInputDTO{
		UserId: id,
	}, nil
}

func ConvertOutputDTOToGetUserByUserIdResponse(input *uc.GetUserByUserIdOutputDTO) (*pb.GetUserByUserIdResponse, error) {
	profile := pb.Profile{
		UserId:    input.User.Id.String(),
		FirstName: input.User.FirstName,
		LastName:  input.User.LastName,
	}

	return &pb.GetUserByUserIdResponse{
		Profile: &profile,
	}, nil
}

func ConvertMetaToJson(input *structpb.Struct) (json.RawMessage, error) {
	meta, err := protojson.Marshal(input)
	if err != nil {
		return nil, ErrIncorrectMeta
	}

	var rawMessage json.RawMessage = meta

	return rawMessage, nil
}

func ConvertStringToInternalId(in string) (uuid.UUID, error) {
	id, err := uuid.Parse(in)
	if err != nil {
		return uuid.Nil, ErrIncorrectInternalId
	}

	return id, nil
}
