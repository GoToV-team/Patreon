package repository

import (
	"errors"
	"patreon/internal/app"
)

var (
	ErrDefaultDB = errors.New("something wrong DB")
	NotFound     = errors.New("user not found")
)

func NewDBError(externalErr error) *app.GeneralError {
	return &app.GeneralError{
		Err:         ErrDefaultDB,
		ExternalErr: externalErr,
	}

}