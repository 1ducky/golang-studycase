package application

import "net/http"

var ErrTodoNotFound = &AppError{
	Code:       "TODO_NOT_FOUND",
	Message:    "Todo not found",
	StatusCode: http.StatusNotFound,
}

var ErrTodoEmpty = &AppError{
	Code:       "TODO_EMPTY",
	Message:    "Task cannot be empty",
	StatusCode: http.StatusBadRequest,
}
