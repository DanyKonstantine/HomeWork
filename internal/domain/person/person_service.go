package person

type Service interface {
	FindAllPer() ([]Person, error)
	FindOnePer(id int64) (*Person, error)
	AddNewPerson(person *Person) error
	UpdatePerson(person *Person, id uint64) error
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
	return (*ps.repo).FindlOnePer(id)
}
func (ps *pservice) AddNewPerson(person *Person) error {
	return (*ps.repo).AddNewPerson(person)
}
func (ps *pservice) UpdatePerson(person *Person, id uint64) error {
	return (*ps.repo).UpdatePerson(person, id)
}
