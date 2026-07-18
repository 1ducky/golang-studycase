package user

type User struct {
	ID       int `json:"id"`
	password string
	Username string `json:"username"`
	Role     string `json:"role"`
}

type CreateRequest struct {
	Username string
	Password string
}
type LoginRequest struct {
	Username string
	Password string
}
