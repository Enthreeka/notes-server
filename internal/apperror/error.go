package apperror

import (
	"errors"
	"fmt"
)

var (
	ErrNotes         = errors.New("notes_not_found")
	ErrInvalidNotes  = errors.New("create_error")
	ErrinvalidUpdate = errors.New("update_error")
)

var (
	ErrNoteNotFound = NewAppError(ErrNotes, "Invalid request", "US-001")
	ErrCreateNotes  = NewAppError(ErrInvalidNotes, "Failed to create notes", "US-002")
	ErrDeleteNotes  = NewAppError(ErrNotes, "Failed to delete notes", "US-003")
	ErrUpdateNotes  = NewAppError(ErrinvalidUpdate, "Failed to update notes", "US-004")
)

type AppError struct {
	Err     error  `json:"-"`
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}

func NewAppError(err error, message string, code string) *AppError {
	return &AppError{
		Err:     err,
		Message: message,
		Code:    code,
	}
}

func (a *AppError) Error() string {
	return fmt.Sprintf("%s: %s (%s)", a.Err, a.Message, a.Code)
}
