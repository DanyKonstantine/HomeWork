package controllers

import (
	"HomeWork/internal/domain/event"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

type EventController struct {
	service *event.Service
}

func NewEventController(s *event.Service) *EventController {
	return &EventController{
		service: s,
	}
}

func (c *EventController) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events, err := (*c.service).FindAll()
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindAll(): %s", err)
			}
			return
		}

		err = success(w, events)
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
		}
	}
}
func (c *EventController) AddNewEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var event event.Event
		err_post := json.NewDecoder(r.Body).Decode(&event)
		if err_post != nil {
			fmt.Println("Error: ", err_post)
		}
		defer r.Body.Close()
		log.Print(event)
		eventId, err := (*c.service).AddNewEvent(&event)
		if err != nil {
			fmt.Printf("EventController.AddNewEvent(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.AddNewEvent(): %s", err)
			}
			return
		}

		err = success(w, eventId)
		if err != nil {
			fmt.Printf("EventController.AddNewEvent(): %s", err)
		}
	}
}
func (c *EventController) FindOne() http.HandlerFunc {
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
		event, err := (*c.service).FindOne(id)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}

		err = success(w, event)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
		}
	}
}
func (c *EventController) UppdateEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var event event.Event
		err_post := json.NewDecoder(r.Body).Decode(&event)
		if err_post != nil {
			fmt.Println("Error: ", err_post)
		}
		defer r.Body.Close()
		log.Print(event)
		uEvent, err := (*c.service).UpdateEvent(&event)
		if err != nil {
			fmt.Printf("EventController.UppdateEvent(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.UppdateEvent(): %s", err)
			}
			return
		}

		err = success(w, uEvent)
		if err != nil {
			fmt.Printf("EventController.AddNewEvent(): %s", err)
		}
	}
}
func (c *EventController) DeletEvent() http.HandlerFunc {
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
		message, err := (*c.service).DeleteEvent(id)
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
