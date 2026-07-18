package http

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data"`
	Meta    *any   `json:"meta,omitempty"` // pointer supaya omitempty bekerja saat nil
}

func JSON[T any](w http.ResponseWriter, status int, data T, message string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(SuccessResponse[T]{Data: data, Success: true, Message: message})
}

func JSONWithMeta[T any](w http.ResponseWriter, status int, data T, meta any, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(SuccessResponse[T]{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    &meta,
	})
}

func Error(w http.ResponseWriter, status int, code string, message string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{Success: false, Code: code, Message: message})

}

func BadRequest(w http.ResponseWriter) {
	w.WriteHeader(400)
	json.NewEncoder(w).Encode(ErrorResponse{Success: false, Code: "ErrorBadRequest", Message: "Bad Request"})
}
func Unauthorized(w http.ResponseWriter) {
	w.WriteHeader(401)
	json.NewEncoder(w).Encode(ErrorResponse{Success: false, Code: "ErrorUnauthorized", Message: "Unauthorized"})
}
func ForBidden(w http.ResponseWriter) {
	w.WriteHeader(403)
	json.NewEncoder(w).Encode(ErrorResponse{Success: false, Code: "ErrorForbidden", Message: "Forbidden"})
}

type ResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (rw *ResponseWriter) WriteHeader(code int) {
	rw.StatusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
