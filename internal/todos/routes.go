package todos

import (
	"net/http"
)

func Register(h *TodoHandler, mux *http.ServeMux) {
	mux.HandleFunc("GET /todos", h.GetAll)
	mux.Handle("POST /todos", ProtectedRouteTodo(http.HandlerFunc(h.Create)))
	mux.Handle("POST /todos/bulk", ProtectedRouteTodo(http.HandlerFunc(h.Upload)))
	mux.Handle("PATCH /todos", ProtectedRouteTodo(http.HandlerFunc(h.Update)))
	mux.Handle("DELETE /todos", ProtectedRouteTodo(http.HandlerFunc(h.Delete)))
}
