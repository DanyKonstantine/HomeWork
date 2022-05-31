package event

import "github.com/upper/db/v4"

type Service interface {
	FindAll() (*[]Event, error)
	FindOne(id uint64) (*Event, error)
	AddNewEvent(event *Event) (*db.ID, error)
	UpdateEvent(event *Event) (*Event, error)
	DeleteEvent(id uint64) (string, error)
}

type service struct {
	repo *Repository
}

func NewService(r *Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) FindAll() (*[]Event, error) {
	return (*s.repo).FindAll()
}

func (s *service) FindOne(id uint64) (*Event, error) {
	return (*s.repo).FindOne(id)
}
func (s *service) AddNewEvent(event *Event) (*db.ID, error) {
	return (*s.repo).AddNewEvent(event)
}
func (s *service) UpdateEvent(event *Event) (*Event, error) {
	return (*s.repo).UpdateEvent(event)
}
func (s *service) DeleteEvent(id uint64) (string, error) {
	return (*s.repo).DeleteEvent(id)
}
