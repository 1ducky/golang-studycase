package auth

import (
	"net/http"
	httpHelper "restApi/internal/http"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	payload, err := httpHelper.DecodeJSON[LoginRequest](r)
	if err != nil {
		httpHelper.BadRequest(w)
		return
	}
	token, err := h.service.Login(r.Context(), payload)
	if err != nil {
		er := errorHandler(err)
		httpHelper.Error(w, er.Status, er.Code, er.Message)
		return
	}

	httpHelper.JSON(w, 200, ResponseToken{Token: token}, "Created")
}
