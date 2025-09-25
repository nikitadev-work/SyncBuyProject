package domain

import (
	"strings"
)

type Currency struct {
	Code string
}

func NewCurrency(code string) (Currency, error) {
	if len(code) != 3 {
		return Currency{}, ErrIncorrectCurrencyFormat
	}

	return Currency{Code: strings.ToUpper(code)}, nil
}

func IsSupportedCurrencyFormat(code string) bool {
	code = strings.ToUpper(code)
	supportedCur := []string{"RUB"}

	for _, v := range supportedCur {
		if v == code {
			return true
		}
	}

	return false
}
