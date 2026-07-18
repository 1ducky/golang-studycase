package image

import (
	"bytes"
	"context"
	"crypto/rand"
	"io"
	"net/http"
)

type Service struct {
	Repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		Repo: r,
	}
}

func (s *Service) Upload(ctx context.Context, owner string, r io.Reader) (string, error) {
	buf := make([]byte, 512)
	n, err := io.ReadFull(r, buf)
	if err != nil && err != io.ErrUnexpectedEOF {
		return "", ErrInternalError
	}

	contentType := http.DetectContentType(buf[:n])
	ext, ok := validationImagetype(contentType)
	if !ok {
		return "", ErrInvalidContentType
	}

	fileName := owner + "_" + rand.Text() + string(ext)

	// gabungkan bytes yang sudah kebaca + sisa stream, jadi tidak perlu
	// buffer seluruh file ke memori
	full := io.MultiReader(bytes.NewReader(buf[:n]), r)

	if err := s.Repo.Save(ctx, fileName, full); err != nil {
		return "", ErrFailedWrite
	}

	return fileName, nil
}
