package image

const FileFieldName = "image"

type ImageType string

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
