package grpcserver

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrIncorrectFirstName  = errors.New("incorrect first name")
	ErrIncorrectLastName   = errors.New("incorrect last name")
	ErrIncorrectChatId     = errors.New("incorrect chat id")
	ErrIncorrectMeta       = errors.New("incorrect meta json")
	ErrIncorrectExternalId = errors.New("incorrect external id")
	ErrIncorrectInternalId = errors.New("incorrect internal id")
)

func ErrToStatus(err error) error {
	switch {
	case errors.Is(err, ErrIncorrectFirstName):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrIncorrectLastName):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrIncorrectChatId):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrIncorrectMeta):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrIncorrectInternalId):
		return status.Error(codes.InvalidArgument, err.Error())
	default:
		return status.Error(codes.Internal, err.Error())
	}
}
