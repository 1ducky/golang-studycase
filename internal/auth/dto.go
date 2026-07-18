package auth

type LoginData struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string
}

type LoginRequest struct {
	Username string
	Password string
}

type tempUserAuth struct {
	ID       int
	Username string
	password string
	Role     string
}

type ResponseToken struct {
	Token string `json:"token"`
}
