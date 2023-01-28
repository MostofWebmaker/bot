package link

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Link {
	return allEntities
}

func (s *Service) Get(linkType string) (*Link, error) {
	links := s.List()
	for _, link := range links {
		if link.LinkType == linkType {
			return &link, nil
		}
	}

	return &Link{}, nil
}

func (s *Service) New(title string, http string, linkType string) (*Link, error) {
	return &Link{
		Title:    title,
		Http:     http,
		LinkType: linkType,
	}, nil
}

func (s *Service) Delete(linkType string) (*Link, error) {
	links := s.List()
	for i, link := range links {
		if link.LinkType == linkType {
			copy(links[i:], links[i+1:])
			links[len(links)-1] = Link{}
			links = links[:len(links)-1]
		}
	}
	return &Link{}, nil
	//// 1. Выполнить сдвиг a[i+1:] влево на один индекс.
	//copy(a[i:], a[i+1:])
	//
	//// 2. Удалить последний элемент (записать нулевое значение).
	//a[len(a)-1] = ""
	//
	//// 3. Усечь срез.
	//a = a[:len(a)-1]

	//return &Link{
	//	Title:    title,
	//	Http:     http,
	//	LinkType: linkType,
	//}, nil
}
