package todos

import (
	"net/http"
	"restApi/internal/auth"
	"restApi/internal/http/middleware"
)

func ProtectedRouteTodo(h http.Handler) http.Handler {
	return middleware.Chain(
		h,
		auth.MustAuthenticate,
		auth.MustAuthorized(auth.RoleAdmin, auth.RoleUser),
	)
}
