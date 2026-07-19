package image

import (
	"bytes"
	"io"
	"net/http"
)

func MapperPartToUploadImage(r io.Reader, ownerID string) (*UploadImage, error) {
	buf := make([]byte, 512)
	n, err := io.ReadFull(r, buf)
	if err != nil && err != io.ErrUnexpectedEOF {
		return nil, ErrInternalError
	}

	contentType := http.DetectContentType(buf[:n])
	ext, ok := validationImagetype(contentType)
	if !ok {
		return nil, ErrInvalidContentType
	}

	imageReader := io.MultiReader(bytes.NewReader(buf[:n]), r)

	return &UploadImage{
		ImageReader: imageReader,
		Ext:         ext,
		OwnerID:     ownerID,
	}, nil
}
