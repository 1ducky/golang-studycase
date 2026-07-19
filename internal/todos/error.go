package todos

import (
	"errors"
)

var ErrTodoNotFound = errors.New("NotFound")
var ErrTodoBadRequest = errors.New("BadRequest")
var ErrTodoTimeOut = errors.New("TimeOut")
var ErrTodoCanceled = errors.New("Canceled")
var ErrTodoUnavaliable = errors.New("Unavaliable")
var ErrUnauthorized = errors.New("Unauthorized")
var ErrPartialBulk = errors.New("PartialBulkFailed")
var ErrBulkError = errors.New("PartialBulkFailed")

func errorHandler(err error) Error {
	switch {
	case errors.Is(err, ErrTodoBadRequest):
		return Error{Code: "BadRequest", Status: 400, Message: "Bad Request"}

	case errors.Is(err, ErrTodoNotFound):
		return Error{Code: "NotFound", Status: 404, Message: "Not Found"}
	case errors.Is(err, ErrUnauthorized):
		return Error{Code: "Unauthorized", Status: 401, Message: "Unauthorized"}
	case errors.Is(err, ErrTodoTimeOut):
		return Error{Code: "TimeOut", Status: 504, Message: "Time Out"}
	case errors.Is(err, ErrTodoCanceled):
		return Error{Code: "Canceled", Status: 408, Message: "Canceled"}
	case errors.Is(err, ErrTodoUnavaliable):
		return Error{Code: "Unavaliable", Status: 503, Message: "Unavaliable"}
	case errors.Is(err, ErrPartialBulk):
		return Error{Code: "PartialBulk", Status: 400, Message: "Some of bulk failed"}
	case errors.Is(err, ErrBulkError):
		return Error{Code: "BulkError", Status: 400, Message: "All of bulk failed"}
	}

	return Error{Code: "InternalError", Status: 500, Message: "Internal Error"}
}
