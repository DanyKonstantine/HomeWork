package event

import (
	"fmt"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
)

var setings = postgresql.ConnectionURL{
	Database: `Event`,
	Host:     `127.0.0.1:5432`,
	User:     `postgres`,
	Password: `postgres`,
}

type Repository interface {
	FindAll() ([]Event, error)
	FindOne(id uint64) (*Event, error)
	AddNewEvent(event *Event) error
	PersonOnEvent(id uint64) (*PersonOnEvent, error)
	UpdateEvent(event *Event, id uint64) error
}

type repository struct {
	// Some internal data
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindAll() ([]Event, error) {
	sess, err := postgresql.Open(setings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()
	eventColection := sess.Collection("Event")
	res := eventColection.Find()
	res = res.OrderBy("Title")
	var event []Event
	if err = res.All(&event); err != nil {
		log.Fatal("eventColection: ", err)
	}
	return event, nil
}

func (r *repository) FindOne(id uint64) (*Event, error) {
	sess, err := postgresql.Open(setings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()
	eventColection := sess.Collection("Event")
	res := eventColection.Find()
	count, _ := res.Count()
	if id <= count {
		q := sess.SQL().SelectFrom("Event").Where("\"ID\" =?", id)
		var event Event
		if err := q.One(&event); err != nil {
			log.Fatal("q.One:", err)
			return nil, err
		}
		return &event, nil
	} else {
		return nil, nil
	}
}
func (r *repository) AddNewEvent(event *Event) error {
	sess, err := postgresql.Open(setings)
	if err != nil {
		log.Fatal("Open: ", err)
		return err
	}
	defer sess.Close()
	res, errr := sess.SQL().
		InsertInto("Event").
		Columns("Title", "ShortDesc", "Desc", "Long", "Let", "Images", "Prewive").
		Values(event.Title, event.ShortDesc, event.Desc, event.Long, event.Let, event.Images, event.Prewive).
		Exec()
	if errr != nil {
		return err
	}
	if res != nil {
		fmt.Println("Corect Insert")
	}
	return nil
}
func (r *repository) PersonOnEvent(id uint64) (*PersonOnEvent, error) {
	var poe PersonOnEvent
	sess, err := postgresql.Open(setings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()
	eventColection := sess.Collection("Person")
	res := eventColection.Find()
	count, _ := res.Count()
	fmt.Printf("Count :", count)

	if id <= count {
		q, err := sess.SQL().Query("SELECT \"Event\".\"Title\", \"Event\".\"ShortDesc\", \"Person\".\"Name\",\"Person\".\"Sername\",\"Person\".\"Age\" FROM \"Person\" JOIN \"Event\" ON \"Person\".\"Event_ID\"=\"Event\".\"ID\" WHERE \"Person\".\"Event_ID\"=?", id)
		if err != nil {
			log.Fatal("Query: ", err)
		}
		if !q.Next() {
			log.Fatal("Expecting one row")
		}

		if err := q.Scan(&poe.Title, &poe.ShortDesc, &poe.Name, &poe.Sername, &poe.Age); err != nil {
			log.Fatal("Scan: ", err)
		}
		if err := q.Close(); err != nil {
			log.Fatal("Close: ", err)
		}
	}
	return &poe, nil
}
func (r *repository) UpdateEvent(event *Event, id uint64) error {
	sess, err := postgresql.Open(setings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()
	eventColection := sess.Collection("Event")
	res := eventColection.Find()
	count, _ := res.Count()
	if id <= count {

		res_up := sess.SQL().Update("Event")
		res_up.Set(event).Where("\"ID\"=?", id).Exec()
		if res_up != nil {
			fmt.Println("Corect Update")
		}
	}
	return err
}
