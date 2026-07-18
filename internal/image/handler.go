package image

import (
	"crypto/rand"
	"io"
	"log"
	"net/http"
	"os"
	"restApi/internal/auth"
	httpHelper "restApi/internal/http"
	"strconv"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
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

	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	mr, err := r.MultipartReader()
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
			continue
		}
		buf := make([]byte, 512)
		n, err := io.ReadFull(part, buf)
		if err != nil && err != io.ErrUnexpectedEOF {
			httpHelper.Error(w, http.StatusInternalServerError, "InternalError", "failed to read file")
			return
		}

		contentType := http.DetectContentType(buf[:n])
		log.Println("Get First Byte To Validate Type")
		log.Println(contentType)

		ext, ok := validationImagetype(contentType)
		if !ok {
			httpHelper.Error(w, http.StatusBadRequest, "InvalidContentType", "file type is not allowed")
			return
		}

		fileName := user.Username + strconv.Itoa(user.ID) + "_" + rand.Text() + string(ext)

		dst, err := os.Create("./storage/" + fileName)
		if err != nil {
			httpHelper.Error(w, http.StatusInternalServerError, "InternalError", "failed to create file")
			return
		}
		defer dst.Close()

		if _, err := dst.Write(buf[:n]); err != nil {
			httpHelper.Error(w, http.StatusInternalServerError, "InternalError", "failed to write file")
			return
		}

		if _, err := io.Copy(dst, part); err != nil {
			httpHelper.Error(w, http.StatusInternalServerError, "InternalError", "failed to write file")
			return
		}

	}
	httpHelper.JSON(w, http.StatusOK, "FileUploaded", "File Successfuly Uploaded")
}

// func (h *Handler) Upload(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		httpHelper.Error(w, http.StatusMethodNotAllowed, "InvalidMethod", "Method is not allowed")
// 		return
// 	}

// 	r.Body = http.MaxBytesReader(w, r.Body, MaxuploadImage)

// 	mr, err := r.MultipartReader()
// 	if err != nil {
// 		httpHelper.Error(w, http.StatusBadRequest, "BadRequest", "invalid multipart request")
// 		return
// 	}

// 	absStoragePath, err := filepath.Abs(storagePath)
// 	if err != nil {
// 		httpHelper.Error(w, http.StatusInternalServerError, "InternalError", "failed to resolve storage path")
// 		return
// 	}

// 	found := false
// 	for {
// 		part, err := mr.NextPart()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			httpHelper.Error(w, http.StatusBadRequest, "BadRequest", "failed to read part")
// 			return
// 		}

// 		if part.FormName() != "image" {
// 			continue
// 		}
// 		found = true

// 		buff := make([]byte, 512)
// 		n, err := io.ReadFull(part, buff)
// 		if err != nil && err != io.ErrUnexpectedEOF {
// 			httpHelper.Error(w, http.StatusInternalServerError, "InternalError", "failed to read file")
// 			return
// 		}

// 		contentType := http.DetectContentType(buff[:n])
// 		if !validationImagetype(contentType) {
// 			httpHelper.Error(w, http.StatusBadRequest, "InvalidContentType", "file type is not allowed")
// 			return
// 		}

// 		safeFilename := filepath.Base(part.FileName())
// 		fullPath := filepath.Join(absStoragePath, safeFilename)

// 		dst, err := os.Create(fullPath)
// 		if err != nil {
// 			httpHelper.Error(w, http.StatusInternalServerError, "InternalError", "failed to create file")
// 			return
// 		}

// 		reader := io.MultiReader(bytes.NewReader(buff[:n]), part)
// 		_, err = io.Copy(dst, reader)
// 		dst.Close()
// 		if err != nil {
// 			httpHelper.Error(w, http.StatusInternalServerError, "InternalError", "failed to save file")
// 			return
// 		}
// 	}

// 	if !found {
// 		httpHelper.Error(w, http.StatusBadRequest, "BadRequest", "please choose file")
// 		return
// 	}

// 	httpHelper.JSON(w, http.StatusOK, "FileUploaded", "File Successfuly Uploaded")
// }
