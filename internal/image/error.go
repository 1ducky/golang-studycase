package image

import "errors"

var ErrInvalidContentType = errors.New("file type is not allowed")
var ErrInternalError = errors.New("failed to save file")
var ErrBadRequest = errors.New("invalid multipart request")
var ErrFormName = errors.New("form name is not allowed")
var ErrFailedWrite = errors.New("failed to save file to storage")

func ErrorHandler(err error) *Error {
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidContentType):
			return &Error{
				Code: "InvalidContentType", Status: 400, Message: "Invalid Content Type",
			}
		case errors.Is(err, ErrBadRequest):
			return &Error{
				Code: "BadRequest", Status: 400, Message: "Bad Request",
			}
		case errors.Is(err, ErrFormName):
			return &Error{
				Code: "FormName", Status: 400, Message: "Form Name Error",
			}
		case errors.Is(err, ErrFailedWrite):
			return &Error{
				Code: "FailedWrite", Status: 500, Message: "Failed to write file",
			}
		default:
			return &Error{
				Code: "InternalError", Status: 500, Message: "Internal Error",
			}
		}
	}
	return nil
}
