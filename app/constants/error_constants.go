package constants

import (
	"errors"
)

var (
	// ErrTodoNotFound ...
	ErrTodoNotFound error = errors.New("Requested todo entity not found")
)
