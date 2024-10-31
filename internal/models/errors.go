package models

import "errors"

var (
	ErrNoRecored = errors.New("models: no matching records found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail = errors.New("models: duplicate email")
)
