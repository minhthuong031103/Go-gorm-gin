package common

import (
	"errors"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"statusCode" example:"400"`
	RootErr    error  `json:"-"`
	Message    string `json:"message" example:"error message"`
	Log        string `json:"log" example:"error log"`
	Key        string `json:"errorKey" example:"error_key"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func (a *AppError) RootError() error {
	var err *AppError
	if errors.As(a.RootErr, &err) {
		return err.RootError()

	}
	return a.RootErr
}

func (a *AppError) Error() string {
	return a.RootError().Error()
}

func InternalError(err error) *AppError {
	return NewErrorResponse(
		err,
		"Đã có lỗi xảy ra, vui lòng thử lại sau",
		err.Error(),
		"INTERNAL_ERROR",
	)
}
func NewCustomError(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}

	return NewErrorResponse(errors.New(msg), msg, msg, key)
}
