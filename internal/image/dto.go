package image

import "mime/multipart"

const FileFieldName = "image"

type ImageType string

type Error struct {
	Code    string `json:"code"` //frontend code translate
	Status  int    `json:"status"`
	Message string `json:"message"`
}

const (
	JPEG ImageType = ".jpg"
	PNG  ImageType = ".png"
	WEBP ImageType = ".webp"
)

var AllowedImageType = map[string]ImageType{
	"image/jpeg": JPEG,
	"image/png":  PNG,
	"image/webp": WEBP,
}

type ImageRequest struct {
	MultipartReader *multipart.Reader
}
