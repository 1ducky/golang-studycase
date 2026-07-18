package application

type AppError struct {
	Code       string
	Message    string
	StatusCode int
}
