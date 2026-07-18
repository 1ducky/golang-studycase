package reader

import (
	"mime/multipart"
	"net/http"
)

func HttpRequestToMultipartReader(r *http.Request, maxReader int64) (*multipart.Reader, error) {
	r.Body = http.MaxBytesReader(nil, r.Body, maxReader)
	return r.MultipartReader()
}
