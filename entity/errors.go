package entity

import "errors"

var ErrEmpty = errors.New("empty string")
var ErrNotFound = errors.New("not found")
var ErrInvalidOperation = errors.New("invalid operation")
var ErrInvalidInput = errors.New("invalid input")
