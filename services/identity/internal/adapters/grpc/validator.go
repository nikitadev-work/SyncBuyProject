package grpcserver

import (
	"encoding/json"
	pb "identity/proto-codegen"
	"unicode"

	"github.com/google/uuid"
)

func CheckFirstName(firstName string) error {
	if firstName == "" {
		return ErrIncorrectFirstName
	}
	for i, char := range firstName {
		if (i == 0 && !unicode.IsLetter(char)) || (i == len(firstName)-1 && char == ' ') {
			return ErrIncorrectFirstName
		}
		if !((unicode.IsLetter(char) || unicode.IsDigit(char) || unicode.IsSpace(char)) && char != ' ') {
			return ErrIncorrectFirstName
		}
	}
	return nil
}

func CheckLastName(lastName string) error {
	if lastName == "" {
		return ErrIncorrectLastName
	}
	for i, char := range lastName {
		if (i == 0 && !unicode.IsLetter(char)) || (i == len(lastName)-1 && char == ' ') {
			return ErrIncorrectLastName
		}
		if !((unicode.IsLetter(char) || unicode.IsDigit(char) || unicode.IsSpace(char)) && char != ' ') {
			return ErrIncorrectLastName
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

func CheckMeta(meta json.RawMessage) error {
	var obj map[string]any
	err := json.Unmarshal(meta, &obj)
	if err != nil || obj == nil {
		return ErrIncorrectMeta
	}
	return nil
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

func ValidateRegisterOrGetTelegramUser(in *pb.RegisterOrGetTelegramUserRequest) error {
	err := CheckFirstName(in.FirstName)
	if err != nil {
		return ErrIncorrectFirstName
	}

	err = CheckLastName(in.LastName)
	if err != nil {
		return ErrIncorrectLastName
	}

	err = CheckChatId(in.ChatId)
	if err != nil {
		return ErrIncorrectChatId
	}

	return nil
}

func CheckInternalId(in string) error {
	_, err := uuid.Parse(in)
	if err != nil {
		return ErrIncorrectInternalId
	}

	return nil
}
