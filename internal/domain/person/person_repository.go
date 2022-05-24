package person

import (
	"fmt"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
)

type PerRepository interface {
	FindAllPer() ([]Person, error)
	FindlOnePer(age int64) (*Person, error)
	AddNewPerson(person *Person) error
	UpdatePerson(person *Person, id uint64) error
}

const PersonCount = int64(10)

type perrepository struct {
}

var setings = postgresql.ConnectionURL{
	Database: `Event`,
	Host:     `127.0.0.1:5432`,
	User:     `postgres`,
	Password: `postgres`,
}

func NewRepository() PerRepository {
	return &perrepository{}
}

func (pr *perrepository) FindAllPer() ([]Person, error) {
	sess, err := postgresql.Open(setings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()
	personColection := sess.Collection("Person")
	res := personColection.Find()
	res = res.OrderBy("Name")
	var person []Person
	if err = res.All(&person); err != nil {
		log.Fatal("personColection: ", err)
	}
	return person, nil
}
func (pr *perrepository) FindlOnePer(id int64) (*Person, error) {
	sess, err := postgresql.Open(setings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()
	personColection := sess.Collection("Person")
	res := personColection.Find()
	count, _ := res.Count()
	if id <= int64(count) {
		q := sess.SQL().SelectFrom("Person").Where("\"ID\" =?", id)
		var person Person
		if err := q.One(&person); err != nil {
			log.Fatal("q.One:", err)
			return nil, err
		}
		return &person, nil
	} else {
		return nil, nil
	}
}
func (pr *perrepository) AddNewPerson(person *Person) error {
	sess, err := postgresql.Open(setings)
	if err != nil {
		log.Fatal("Open: ", err)
		return err
	}
	defer sess.Close()
	res, errr := sess.SQL().
		InsertInto("Person").
		Columns("Name", "Sername", "Age", "Event_ID").
		Values(person.Name, person.Sername, person.Age, person.Event_id).
		Exec()
	if errr != nil {
		return err
	}
	if res != nil {
		fmt.Println("Corect Insert")
	}
	return nil
}
func (pr *perrepository) UpdatePerson(person *Person, id uint64) error {
	sess, err := postgresql.Open(setings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()
	personColection := sess.Collection("Person")
	res := personColection.Find()
	count, _ := res.Count()
	if id <= count {
		res_up := sess.SQL().Update("Person")
		res_up.Set(person).Where("\"ID\"=?", id).Exec()
		if res_up != nil {
			fmt.Println("Corect Update")
		}
	}
	return err
}
