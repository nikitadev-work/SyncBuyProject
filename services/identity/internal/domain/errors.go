package domain

import "errors"

var (
	ErrIncorrectStatus       = errors.New("status must be active/blocked/deleted")
	ErrIncorrectProviderType = errors.New("provider type must be telegram")
	ErrIncorrectFirstName    = errors.New("incorrect first name")
	ErrIncorrectLastName     = errors.New("incorrect last name")
	ErrIncorrectExternalId   = errors.New("incorrect external id")
	ErrIncorrectChatId       = errors.New("incorrect chat id")
	ErrIncorrectMeta         = errors.New("incorrect meta json")
	ErrChangeDeletedStatus   = errors.New("can not change deleted status")
)
