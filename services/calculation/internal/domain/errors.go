package domain

import "errors"

var (
	ErrCurrencyMismatch        = errors.New("currency mismatch")
	ErrNegativeAmount          = errors.New("negative amount")
	ErrSelfPayment             = errors.New("payer and payee must differ")
	ErrIncorrectCurrencyFormat = errors.New("incorrect currency format")
	ErrEmptyParticipantsList   = errors.New("empty participants list")
	ErrEmptyExpensesList       = errors.New("empty expenses list")
	ErrOnlyOneParticipant      = errors.New("only one participant")
	ErrNegativeSubResult       = errors.New("subtraction: first money less then second")
	ErrIncorrectPartsAmount    = errors.New("number of parts must be positive")
	ErrCurrenciesInListDiffers = errors.New("currencies must be the same to divide money into parts")
	ErrNotPositiveIntentAmount = errors.New("amount of intent must be positive")
)
