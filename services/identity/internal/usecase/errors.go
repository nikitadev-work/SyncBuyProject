package usecase

import "errors"

var (
	ErrUserDoesNotExist     = errors.New("user does not exist")
	ErrIdentityDoesNotExist = errors.New("identity for this user does not exist")
	ErrCreateUser           = errors.New("error while creating new user")
	ErrCreateIdentity       = errors.New("error while creating new identity")
)
