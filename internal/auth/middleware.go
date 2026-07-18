package auth

import (
	"net/http"
	httpHelper "restApi/internal/http"
)

func MustAuthorized(role ...Role) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, ok := CheckRole(r.Context(), role...)
			if !ok {
				httpHelper.ForBidden(w)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func MustAuthenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, ok := GetAuht(r.Context())
		if !ok {
			httpHelper.Unauthorized(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}
