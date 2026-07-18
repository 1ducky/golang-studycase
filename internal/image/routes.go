package image

import "net/http"

func Register(h *Handler, mux *http.ServeMux) {
	mux.Handle("POST /upload", http.HandlerFunc(h.Upload))
}
