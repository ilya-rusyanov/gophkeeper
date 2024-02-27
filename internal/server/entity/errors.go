package entity

import "errors"

// ErrUserAlreadyExists means such user is already registered
var ErrUserAlreadyExists = errors.New("user already exists")
