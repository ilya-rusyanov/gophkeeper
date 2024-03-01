package entity

import "errors"

// ErrAuthAbsent means authentication data is missing
var ErrAuthAbsent = errors.New("auth data is missing")
