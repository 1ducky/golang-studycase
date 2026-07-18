package todos

import (
	"context"
	"log"
	"net/http"
	"restApi/internal/auth"
	httpHelper "restApi/internal/http"
	"strconv"
	"time"
)

type TodoHandler struct {
	service *Service
}

func NewTodoHandler(s *Service) *TodoHandler {
	return &TodoHandler{service: s}
}

func (h *TodoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	query := httpHelper.ParseGenericRequest(r)
	if err != nil || id == 0 {
		todos, meta, err := h.service.GetAll(r.Context(), query)
		if err != nil {
			er := errorHandler(err)
			httpHelper.Error(w, er.Status, er.Code, er.Message)
			return
		}
		httpHelper.JSONWithMeta(w, 200, todos, meta, "Geting")
		return
	}
	h.GetById(w, r)

}

func (h *TodoHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		httpHelper.BadRequest(w)

		return
	}
	todos, err := h.service.GetById(r.Context(), id)
	if err != nil {
		er := errorHandler(err)
		httpHelper.Error(w, er.Status, er.Code, er.Message)
		return
	}
	httpHelper.JSON(w, 200, todos, "Getting")
}

func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	auth, ok := auth.GetAuht(r.Context())
	ctxT, cancel := context.WithTimeout(r.Context(), 1000*time.Millisecond)
	defer cancel()
	log.Print(auth, ok)
	if !ok {
		httpHelper.Unauthorized(w)
		return
	}
	payload, err := httpHelper.DecodeJSON[CreateRequest](r)

	if err != nil {
		httpHelper.BadRequest(w)
		return
	}
	todo, err := h.service.Create(ctxT, payload)
	if err != nil {
		er := errorHandler(err)
		httpHelper.Error(w, er.Status, er.Code, er.Message)
		return
	}

	httpHelper.JSON(w, 200, todo, "Created")
}
func (h *TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
	_, ok := auth.GetAuht(r.Context())
	if !ok {
		httpHelper.Unauthorized(w)
		return
	}

	payload, err := httpHelper.DecodeJSON[UpdateRequest](r)
	if err != nil {
		httpHelper.BadRequest(w)
		return
	}
	todo, err := h.service.Update(r.Context(), payload)
	if err != nil {
		er := errorHandler(err)
		httpHelper.Error(w, er.Status, er.Code, er.Message)
		return
	}

	httpHelper.JSON(w, 200, todo, "Updated")
}

func (h *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	_, ok := auth.GetAuht(r.Context())
	if !ok {
		httpHelper.Unauthorized(w)
		return
	}

	payload, err := httpHelper.DecodeJSON[DeleteRequest](r)
	if err != nil {
		httpHelper.BadRequest(w)
		return
	}
	err = h.service.Delete(r.Context(), payload)
	if err != nil {
		er := errorHandler(err)
		httpHelper.Error(w, er.Status, er.Code, er.Message)
		return
	}

	httpHelper.JSON(w, 200, true, "Deleted")
}
