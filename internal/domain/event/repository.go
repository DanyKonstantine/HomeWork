package event

import (
	"fmt"
	"github.com/upper/db/v4"
)

type Repository interface {
	FindAll() (*[]Event, error)
	FindOne(id uint64) (*Event, error)
	AddNewEvent(event *Event) (*db.ID, error)
	UpdateEvent(event *Event) (*Event, error)
	DeleteEvent(id uint64) (string, error)
}

type repository struct {
	sess db.Session
}

func NewRepository(sess db.Session) Repository {
	return &repository{sess: sess}
}

func (r *repository) FindAll() (*[]Event, error) {
	var events []Event
	err := r.sess.Collection("events").Find().OrderBy("title").All(&events)
	if err != nil {
		fmt.Println("Records not found : ", err)
		return nil, err
	}
	return &events, nil
}

func (r *repository) FindOne(id uint64) (*Event, error) {
	var event Event
	err := r.sess.Collection("events").Find(id).One(&event)
	if err != nil {
		fmt.Println("One event not found : ", err)
		return nil, err
	}
	return &event, nil
}
func (r *repository) AddNewEvent(event *Event) (*db.ID, error) {
	eventReturn, err := r.sess.Collection("events").Insert(event)
	eventId := eventReturn.ID()
	if err != nil {
		fmt.Println("Can`t add event : ", err)
		return nil, err
	}
	return &eventId, nil
}
func (r *repository) UpdateEvent(event *Event) (*Event, error) {
	err := r.sess.Collection("events").Find(event.ID).Update(event)
	if err != nil {
		fmt.Println("One event not found : ", err)
		return nil, err
	}
	return event, nil
}

func (r *repository) DeleteEvent(id uint64) (string, error) {
	err := r.sess.Collection("events").Find(id).Delete()
	if err != nil {
		fmt.Println("One event not found : ", err)
		return " ", err
	}
	return "Record successfully deleted", nil
}
