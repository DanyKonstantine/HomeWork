package person

import (
	"fmt"
	"github.com/upper/db/v4"
)

type PerRepository interface {
	FindAllPer() ([]Person, error)
	FindOnePer(id int64) (*Person, error)
	AddNewPerson(person *Person) (*db.ID, error)
	UpdatePerson(person *Person) (*Person, error)
	DeletePerson(id uint64) (string, error)
}

type perrepository struct {
	sess db.Session
}

func NewRepository(sess db.Session) PerRepository {
	return &perrepository{sess: sess}
}

func (pr *perrepository) FindAllPer() ([]Person, error) {
	var persons []Person
	err := pr.sess.Collection("persons").Find().OrderBy("name").All(&persons)
	if err != nil {
		fmt.Println("Records not found : ", err)
		return nil, err
	}
	return persons, nil
}
func (pr *perrepository) FindOnePer(id int64) (*Person, error) {
	var person Person
	err := pr.sess.Collection("persons").Find(id).One(&person)
	if err != nil {
		fmt.Println("One person not found : ", err)
		return nil, err
	}
	return &person, nil
}
func (pr *perrepository) AddNewPerson(person *Person) (*db.ID, error) {
	personReturn, err := pr.sess.Collection("persons").Insert(person)
	personId := personReturn.ID()
	if err != nil {
		fmt.Println("Can`t add person", err)
		return nil, err
	}
	return &personId, nil
}
func (pr *perrepository) UpdatePerson(person *Person) (*Person, error) {
	err := pr.sess.Collection("persons").Find(person.ID).Update(person)
	if err != nil {
		fmt.Println("One person not found : ", err)
		return nil, err
	}
	return person, nil
}
func (pr *perrepository) DeletePerson(id uint64) (string, error) {
	err := pr.sess.Collection("persons").Find(id).Delete()
	if err != nil {
		fmt.Println("One person not found : ", err)
		return " ", err
	}
	return "Record successfully deleted", nil
}
