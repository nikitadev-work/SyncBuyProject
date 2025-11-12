package repository

import "errors"

var (
	ErrParsingResponse    = errors.New("error while parsing response from database")
	ErrUnlockPurchase     = errors.New("cannot unlock purchase: settlement already initiated")
	ErrFinishPurchase     = errors.New("cannot finish purchase: purchase is unlocked")
	ErrEditLockedPurchase = errors.New("cannot modify purchase: purchase is locked")
	ErrLockPurchase       = errors.New("cannot lock purchase: purchase is already locked")
)
