package controllers

import (
	"HomeWork/internal/domain/person"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
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
		var person person.Person
		err_post := json.NewDecoder(r.Body).Decode(&person)
		if err_post != nil {
			fmt.Println("Error: ", err_post)
		}
		defer r.Body.Close()
		personId, err := (*c.service).AddNewPerson(&person)
		if err != nil {
			fmt.Printf("PerconControler.AddNewPeson(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("PerconControler.AddNewPeson(): %s", err)
			}
			return
		}

		err = success(w, personId)
		if err != nil {
			fmt.Printf("PerconControler.AddNewPeson(): %s", err)
		}
	}
}

func (c *PerconControler) UppdatePeson() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var person person.Person
		err_post := json.NewDecoder(r.Body).Decode(&person)
		if err_post != nil {
			fmt.Println("Error: ", err_post)
		}
		defer r.Body.Close()
		uPerson, err := (*c.service).UpdatePerson(&person)
		if err != nil {
			fmt.Printf("PerconControler.AddNewPeson(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("PerconControler.AddNewPeson(): %s", err)
			}
			return
		}
		err = success(w, uPerson)
		if err != nil {
			fmt.Printf("PerconControler.AddNewPeson(): %s", err)
		}
	}
}
func (c *PerconControler) DeletePerson() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}
		message, err := (*c.service).DeletePerson(id)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}

		err = success(w, message)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
		}
	}
}
