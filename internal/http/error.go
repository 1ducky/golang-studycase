package http

type ErrorResponse struct {
	Success bool   `json:"success"`
	Code    string `json:"code"` //frontend code translate
	Message string `json:"message"`
}
