package event

type Service interface {
	FindAll() ([]Event, error)
	FindOne(id uint64) (*Event, error)
	AddNewEvent(event *Event) error
	PersonOnEvent(id uint64) (*PersonOnEvent, error)
	UpdateEvent(event *Event, id uint64) error
}

type service struct {
	repo *Repository
}

func NewService(r *Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) FindAll() ([]Event, error) {
	return (*s.repo).FindAll()
}

func (s *service) FindOne(id uint64) (*Event, error) {
	return (*s.repo).FindOne(id)
}
func (s *service) AddNewEvent(event *Event) error {
	return (*s.repo).AddNewEvent(event)
}
func (s *service) PersonOnEvent(id uint64) (*PersonOnEvent, error) {
	return (*s.repo).PersonOnEvent(id)
}
func (s *service) UpdateEvent(event *Event, id uint64) error {
	return (*s.repo).UpdateEvent(event, id)
}
