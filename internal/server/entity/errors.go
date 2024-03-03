package entity

import "errors"

var (
	// ErrUserAlreadyExists means such user is already registered
	ErrUserAlreadyExists = errors.New("user already exists")
	// ErrAuthFailed means that user failed to pass auth
	ErrAuthFailed = errors.New("authentication failure")
	// ErrNoSuchUser means user not found
	ErrNoSuchUser = errors.New("no such user")
	// ErrWrongPassword means password does not match
	ErrWrongPassword = errors.New("wrong password")
	// ErrRecordAlreadyExists when such data record is already present
	ErrRecordAlreadyExists = errors.New("record already exists")
)
