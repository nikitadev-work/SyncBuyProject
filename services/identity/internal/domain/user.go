package domain

import (
	"time"
	"unicode"

	"github.com/google/uuid"
)

type Status int

const (
	Active Status = iota
	Blocked
	Deleted
)

type User struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(id uuid.UUID, firstName string, lastName string, status Status,
	createdAt time.Time, updatedAt time.Time) (*User, error) {
	err := CheckFirstName(firstName)
	if err != nil {
		return nil, err
	}

	err = CheckLastName(lastName)
	if err != nil {
		return nil, err
	}

	return &User{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		Status:    status,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

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

func (us *User) SetStatus(newStatus Status) error {
	switch us.Status {
	case Active, Blocked:
		if !contains([]Status{Active, Blocked, Deleted}, newStatus) {
			return ErrIncorrectStatus
		}
		us.Status = newStatus
		return nil
	case Deleted:
		return ErrChangeDeletedStatus
	}
	return ErrIncorrectStatus
}

func (us *User) GetStatus() (string, error) {
	switch us.Status {
	case Active:
		return "active", nil
	case Blocked:
		return "blocked", nil
	case Deleted:
		return "deleted", nil
	default:
		return "", ErrIncorrectStatus
	}
}

func (us *User) CanUsePlatform() bool {
	switch us.Status {
	case Active:
		return true
	default:
		return false
	}
}

func contains(arr []Status, newStatus Status) bool {
	for _, item := range arr {
		if item == newStatus {
			return true
		}
	}

	return false
}
