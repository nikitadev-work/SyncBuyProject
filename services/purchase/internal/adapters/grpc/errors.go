package grpcserver

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrIncorrectInternalId  = errors.New("incorrect internal id")
	ErrIncorrectTitle       = errors.New("incorrect title")
	ErrIncorrectDescription = errors.New("incorrect description")
	ErrIncorrectCurrency    = errors.New("incorrect currency")
	ErrIncorrectToken       = errors.New("incorrect token")
	ErrIncorrectAmount      = errors.New("incorrect amount")
)

func ErrToStatus(err error) error {
	switch {
	case errors.Is(err, ErrIncorrectTitle):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrIncorrectDescription):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrIncorrectCurrency):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrIncorrectToken):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrIncorrectInternalId):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrIncorrectAmount):
		return status.Error(codes.InvalidArgument, err.Error())
	default:
		return status.Error(codes.Internal, err.Error())
	}
}
