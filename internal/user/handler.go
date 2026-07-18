package user

import (
	"log"
	"net/http"
	httpHelper "restApi/internal/http"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	payload, err := httpHelper.DecodeJSON[CreateRequest](r)
	if err != nil {
		httpHelper.BadRequest(w)
		return
	}
	log.Print(payload)
	success, err := h.service.Create(r.Context(), payload)
	if err != nil {
		er := errorHandler(err)
		httpHelper.Error(w, er.Status, er.Code, er.Message)
		return
	}

	httpHelper.JSON(w, 200, Response{ok: success}, "Created")
}
