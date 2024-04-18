package domain

import "errors"

var (
	ErrInternalServer     = errors.New("internal server error")
	ErrInvalidRequest     = errors.New("invalid request")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrLoginAlreadyExists = errors.New("login already exists")
	ErrTokenExpired       = errors.New("token expired")
	BookNotFound          = errors.New("book not found")
)
