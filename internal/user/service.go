package user

import "context"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Create(ctx context.Context, payload CreateRequest) (bool, error) {
	if payload.Password == "" || payload.Username == "" {
		return false, ErrorBadRequest
	}
	success, err := s.repo.create(ctx, payload)
	return success, err
}
