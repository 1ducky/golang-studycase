package auth

type Role string

const (
	RoleUser  Role = "User"
	RoleAdmin Role = "Admin"
)

type AuthClaims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     `json:"role"`
}

func NewAuthClaims(id int) *AuthClaims {
	return &AuthClaims{
		ID:   id,
		Role: RoleUser,
	}
}
