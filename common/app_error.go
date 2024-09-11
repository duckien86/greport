package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int         `json:"status_code"`
	RootErr    error       `json:"-"`
	Message    string      `json:"message"`
	Log        string      `json:"log"`
	Key        string      `json:"error_key"`
	Details    interface{} `json:"details"`
}

func NewErrResponse(root error, msg, log, key string, details interface{}) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
		Details:    details,
	}
}

func NewFullErrResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Key:        key,
	}
}

func NewCustomError(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrResponse(root, msg, root.Error(), key, nil)
	}
	return NewErrResponse(errors.New(msg), msg, msg, key, nil)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}
	return e.RootErr
}

// Để struct AppError trở thành struct error trong GO
// Ta kế thừa hàm Error() string
func (e *AppError) Error() string {
	return e.RootError().Error()
}

func ErrDB(err error) *AppError {
	return NewFullErrResponse(
		http.StatusInternalServerError,
		err,
		"something went wrong with DB",
		err.Error(),
		"DB_ERROR",
	)
}

// return http_code : 400
func ErrInvalidRequest(err error) *AppError {
	return NewErrResponse(err, "invalid request", err.Error(), "ErrInvalidRequest", nil)
}

func ErrValidationData(err error, details interface{}) *AppError {
	return NewErrResponse(err, "invalid request", "", "ErrInvalidRequest", details)
}

func ErrInternal(err error) *AppError {
	return NewFullErrResponse(http.StatusInternalServerError, err,
		"something went wrong with serve", err.Error(), "ErrInternal")
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotList%s", entity),
	)
}
func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotGet%s", entity),
	)
}
func ErrEntityExisted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s already existed", strings.ToLower(entity)),
		fmt.Sprintf("Err%sAlreadyExisted", entity),
	)
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("Err%sNotFound", entity),
	)
}

func ErrEntityDeleted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s was deleted", strings.ToLower(entity)),
		fmt.Sprintf("Err%sDeleted", entity),
	)
}
func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete%s", entity),
	)
}
func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("cannot update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotUpdate%s", entity),
	)
}
func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("cannot create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate%s", entity),
	)
}
func ErrNoPermision(err error) *AppError {
	return NewCustomError(
		err,
		"you have no permision",
		"ErrNoPermision",
	)
}
