package info

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Info {
	return allEntities
}

func (s *Service) Get(idx int) (*Info, error) {
	return &allEntities[idx], nil
}
