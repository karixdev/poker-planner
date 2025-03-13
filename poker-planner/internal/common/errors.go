package common

import "errors"

type Error struct {
	AppError    error
	SourceError error
}

func NewError(appError, sourceError error) error {
	return &Error{
		AppError:    appError,
		SourceError: sourceError,
	}
}

func (e *Error) Error() string {
	return errors.Join(e.AppError, e.SourceError).Error()
}
