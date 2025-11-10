package repository

import "errors"

var (
	ErrParsingResponse = errors.New("error while parsing response from database")
)
