package usecase

import "errors"

var (
	ErrGetSnapshot = errors.New("cannot get snapshot: purchase must be locked")
)
