package domain

import (
	"encoding/json"
	"time"
	"unicode"

	"github.com/google/uuid"
)

type ProviderType int

const (
	Telegram ProviderType = iota
)

type Identity struct {
	ExternalId   string
	InternalId   uuid.UUID
	ProviderType ProviderType
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ChatId       string
	Meta         json.RawMessage
}

func NewIdentity(externalId string, internalId uuid.UUID, pType ProviderType,
	createdAt time.Time, updatedAt time.Time, chatId string, meta json.RawMessage) (*Identity, error) {
	err := CheckExternalId(externalId)
	if err != nil {
		return nil, err
	}
	err = CheckChatId(chatId)
	if err != nil {
		return nil, err
	}
	err = CheckMeta(meta)
	if err != nil {
		return nil, err
	}
	return &Identity{
		ExternalId:   externalId,
		InternalId:   internalId,
		ProviderType: pType,
		CreatedAt:    createdAt,
		ChatId:       chatId,
		Meta:         meta,
	}, nil
}

func CheckExternalId(externalId string) error {
	if externalId == "" {
		return ErrIncorrectExternalId
	}
	for _, char := range externalId {
		if !(char != ' ' && unicode.IsDigit(char)) {
			return ErrIncorrectExternalId
		}
	}
	return nil
}

func CheckChatId(chatId string) error {
	if chatId == "" {
		return ErrIncorrectChatId
	}
	for i, char := range chatId {
		if i == 0 && char == '-' {
			continue
		}
		if !(char != ' ' && unicode.IsDigit(char)) {
			return ErrIncorrectChatId
		}
	}
	return nil
}

func (idt *Identity) SetProviderType(newType ProviderType) error {
	switch newType {
	case Telegram:
		idt.ProviderType = newType
		return nil
	}
	return ErrIncorrectProviderType
}

func (idt *Identity) GetProviderType() (string, error) {
	switch idt.ProviderType {
	case Telegram:
		return "telegram", nil
	default:
		return "", ErrIncorrectProviderType
	}
}

func CheckMeta(meta json.RawMessage) error {
	var obj map[string]any
	err := json.Unmarshal(meta, &obj)
	if err != nil || obj == nil {
		return ErrIncorrectMeta
	}
	return nil
}
