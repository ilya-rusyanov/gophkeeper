package entity

import "errors"

var (
	// ErrUserAlreadyExists means such user is already registered
	ErrUserAlreadyExists = errors.New("user already exists")
	// ErrAuthFailed means that user failed to pass auth
	ErrAuthFailed = errors.New("authentication failure")
)
