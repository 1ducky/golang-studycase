package middleware

import (
	"net/http"
	"restApi/internal/auth"
)

func NewAuthMiddleware(s *auth.Service) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			u, err := s.ParseJWTToken(r.Context(), token)
			if err != nil {
				next.ServeHTTP(w, r)
				return

			}
			claims := auth.NewAuthClaims(u.ID)
			ctx := auth.WithAuth(r.Context(), claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
