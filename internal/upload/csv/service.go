package csv

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Validate() int {
	return 0
}
