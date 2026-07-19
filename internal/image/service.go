package image

import (
	"context"
	"crypto/rand"
)

type Service struct {
	Repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		Repo: r,
	}
}

func (s *Service) Upload(ctx context.Context, req *UploadImage) (string, error) {

	fileName := req.OwnerID + "_" + rand.Text() + string(req.Ext)

	if err := s.Repo.Save(ctx, fileName, req.ImageReader); err != nil {
		return "", ErrFailedWrite
	}
	return fileName, nil
}
