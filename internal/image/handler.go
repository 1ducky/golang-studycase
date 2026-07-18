package image

import (
	"io"
	"net/http"
	"restApi/internal/auth"
	httpHelper "restApi/internal/http"
	"restApi/internal/reader"
	"strconv"
)

type Handler struct {
	Service Service
}

func NewHandler(ser *Service) *Handler {
	return &Handler{
		Service: *ser,
	}
}

func (h *Handler) Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpHelper.Error(w, http.StatusMethodNotAllowed, "InvalidMethod", "Method is not allowed")
		return
	}
	user, ok := auth.GetAuht(r.Context())
	if !ok {
		httpHelper.Unauthorized(w)
		return
	}

	mr, err := reader.HttpRequestToMultipartReader(r, 10<<20)
	if err != nil {
		httpHelper.BadRequest(w)
		return
	}

	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			httpHelper.BadRequest(w)
			return
		}
		defer part.Close()

		if part.FormName() != FileFieldName {
			part.Close()
			continue
		}
		mapper, err := PartToImageRequest(part, strconv.Itoa(user.ID))
		if err != nil {
			er := ErrorHandler(err)
			httpHelper.Error(w, er.Status, er.Code, er.Message)
			return
		}

		_, err = h.Service.Upload(r.Context(), mapper)
		if err != nil {
			er := ErrorHandler(err)
			httpHelper.Error(w, er.Status, er.Code, er.Message)
			return
		}
	}

	httpHelper.JSON(w, http.StatusOK, "FileUploaded", "File Successfuly Uploaded")
}
