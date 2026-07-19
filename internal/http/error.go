package http

type ErrorResponse struct {
	Success bool   `json:"success"`
	Code    string `json:"code"` //frontend code translate
	Message string `json:"message"`
}

type ErrorResponseWithMeta[T any] struct {
	Success bool   `json:"success"`
	Code    string `json:"code"` //frontend code translate
	Message string `json:"message"`
	Meta    T      `json:"meta"`
}
