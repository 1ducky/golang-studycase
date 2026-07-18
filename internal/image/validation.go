package image

const MaxuploadImage = 10

func validationImagetype(contentType string) (ImageType, bool) {
	v, ok := AllowedImageType[contentType]
	if !ok {
		return "", false
	}
	return v, true
}
