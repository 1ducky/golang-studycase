package user

import "net/http"

func Register(h *Handler, mux *http.ServeMux) {
	mux.HandleFunc("POST /user", h.Create)
}
