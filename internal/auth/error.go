package auth

import "errors"

var ErrorNotFound = errors.New("NotFound")
var ErrorBadRequest = errors.New("BadRequest")

type Error struct {
	Code    string `json:"code"` //frontend code translate
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func errorHandler(err error) Error {
	switch {
	case errors.Is(err, ErrorBadRequest):
		return Error{Code: "BadRequest", Status: 400, Message: "Bad Request"}

	case errors.Is(err, ErrorNotFound):
		return Error{Code: "NotFound", Status: 404, Message: "Not Found"}
	}
	return Error{Code: "InternalError", Status: 500, Message: "Internal Error"}
}
