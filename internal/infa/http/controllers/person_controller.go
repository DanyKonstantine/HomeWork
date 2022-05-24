package controllers

import (
	"HomeWork/internal/domain/person"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

type PerconControler struct {
	service *person.Service
}

func NewPerconControler(ps *person.Service) *PerconControler {
	return &PerconControler{
		service: ps,
	}
}

func (c *PerconControler) FindAllPer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events, err := (*c.service).FindAllPer()
		if err != nil {
			fmt.Printf("PerconControler.FindAllPer(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("PerconControler.FindAllPer(): %s", err)
			}
			return
		}

		err = success(w, events)
		if err != nil {
			fmt.Printf("PerconControler.FindAllPer(): %s", err)
		}
	}
}

func (c *PerconControler) FindOnePer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("PerconControler.FindOnePer(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("PerconControler.FindOnePer(): %s", err)
			}
			return
		}
		event, err := (*c.service).FindOnePer(id)
		if err != nil {
			fmt.Printf("PerconControler.FindOnePer(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("PerconControler.FindOnePer(): %s", err)
			}
			return
		}

		err = success(w, event)
		if err != nil {
			fmt.Printf("PerconControler.FindOnePer(): %s", err)
		}
	}
}
func (c *PerconControler) AddNewPeson() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		formdata := map[string]string{}
		err_post := json.NewDecoder(r.Body).Decode(&formdata)
		if err_post != nil {
			fmt.Println("Error: ", err_post)
		}
		defer r.Body.Close()
		age := formdata["Age"]
		ageint, err_conv := strconv.Atoi(age)
		if err_conv != nil {
			fmt.Printf("Сonvert: ", err_conv)
		}
		event_id := formdata["Event_ID"]
		event_id_int, err_conv := strconv.Atoi(event_id)
		if err_conv != nil {
			fmt.Printf("Сonvert: ", err_conv)
		}
		person := person.Person{
			Name:     formdata["Name"],
			Sername:  formdata["Sername"],
			Age:      ageint,
			Event_id: event_id_int,
		}
		log.Print(person)
		err := (*c.service).AddNewPerson(&person)
		if err != nil {
			fmt.Printf("PerconControler.AddNewPeson(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("PerconControler.AddNewPeson(): %s", err)
			}
			return
		}

		err = noContent(w)
		if err != nil {
			fmt.Printf("PerconControler.AddNewPeson(): %s", err)
		}
	}
}

func (c *PerconControler) UppdatePeson() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("PerconControler.FindOnePer(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("PerconControler.FindOnePer(): %s", err)
			}
			return
		}
		formdata := map[string]string{}
		err_post := json.NewDecoder(r.Body).Decode(&formdata)
		if err_post != nil {
			fmt.Println("Error: ", err_post)
		}
		defer r.Body.Close()
		age := formdata["Age"]
		ageint, err_conv := strconv.Atoi(age)
		if err_conv != nil {
			fmt.Printf("Сonvert: ", err_conv)
		}
		event_id := formdata["Event_ID"]
		event_id_int, err_conv := strconv.Atoi(event_id)
		if err_conv != nil {
			fmt.Printf("Сonvert: ", err_conv)
		}
		person := person.Person{
			Name:     formdata["Name"],
			Sername:  formdata["Sername"],
			Age:      ageint,
			Event_id: event_id_int,
		}
		log.Print(person)
		err = (*c.service).UpdatePerson(&person, id)
		if err != nil {
			fmt.Printf("PerconControler.AddNewPeson(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("PerconControler.AddNewPeson(): %s", err)
			}
			return
		}

		err = noContent(w)
		if err != nil {
			fmt.Printf("PerconControler.AddNewPeson(): %s", err)
		}
	}
}
