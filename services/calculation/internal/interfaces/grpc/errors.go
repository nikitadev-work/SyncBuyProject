package grpcserver

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNegativeAmount          = errors.New("negative amount")
	ErrEmptyParticipantsList   = errors.New("list of participants must be non empty")
	ErrEmptyExpensesList       = errors.New("list of expenses must be non empty")
	ErrIncorrectCurrencyFormat = errors.New("incorrect currency format")
	ErrInvalidInputData        = errors.New("invalid input data")
)

func ErrToStatus(err error) error {
	switch {
	case errors.Is(err, ErrEmptyExpensesList):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrNegativeAmount):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrEmptyParticipantsList):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrIncorrectCurrencyFormat):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrInvalidInputData):
		return status.Error(codes.InvalidArgument, err.Error())
	default:
		return status.Error(codes.Internal, err.Error())
	}
}
