package todos

import (
	"net/http"
)

func Register(h *TodoHandler, mux *http.ServeMux) {
	mux.HandleFunc("GET /todos", h.GetAll)
	mux.Handle("POST /todos", ProtectedRouteTodo(http.HandlerFunc(h.Create)))
	mux.Handle("PATCH /todos", ProtectedRouteTodo(http.HandlerFunc(h.Update)))
	mux.Handle("DELETE /todos", ProtectedRouteTodo(http.HandlerFunc(h.Delete)))
}
