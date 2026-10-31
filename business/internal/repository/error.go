package repository

import(
	"errors"
)

var	ErrInvalidBody = errors.New("Invalid body for creating businesss")
var ErrNotFound = errors.New("Business with that id not found")