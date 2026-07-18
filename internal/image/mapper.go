package image

import (
	"bytes"
	"io"
	"net/http"
)

func PartToImageRequest(r io.Reader, ownerID string) (*ImageRequest, error) {
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

	return &ImageRequest{
		ImageReader: imageReader,
		Ext:         ext,
		OwnerID:     ownerID,
	}, nil
}
