package person

import "github.com/upper/db/v4"

type Service interface {
	FindAllPer() ([]Person, error)
	FindOnePer(id int64) (*Person, error)
	AddNewPerson(person *Person) (*db.ID, error)
	UpdatePerson(person *Person) (*Person, error)
	DeletePerson(id uint64) (string, error)
}

type pservice struct {
	repo *PerRepository
}

func NewService(pr *PerRepository) Service {
	return &pservice{
		repo: pr,
	}
}
func (ps *pservice) FindAllPer() ([]Person, error) {
	return (*ps.repo).FindAllPer()
}
func (ps *pservice) FindOnePer(id int64) (*Person, error) {
	return (*ps.repo).FindOnePer(id)
}
func (ps *pservice) AddNewPerson(person *Person) (*db.ID, error) {
	return (*ps.repo).AddNewPerson(person)
}
func (ps *pservice) UpdatePerson(person *Person) (*Person, error) {
	return (*ps.repo).UpdatePerson(person)
}
func (ps *pservice) DeletePerson(id uint64) (string, error) {
	return (*ps.repo).DeletePerson(id)
}
