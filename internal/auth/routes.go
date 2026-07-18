package auth

import "net/http"

func Register(h *Handler, mux *http.ServeMux) {
	mux.HandleFunc("POST /login", h.Login)
}
