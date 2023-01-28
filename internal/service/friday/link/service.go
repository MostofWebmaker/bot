package link

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Link {
	return allEntities
}

func (s *Service) Get(idx int) (*Link, error) {
	return &allEntities[idx], nil
}
