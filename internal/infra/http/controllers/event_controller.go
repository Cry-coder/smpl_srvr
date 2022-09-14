package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Cry-coder/smpl_srvr/internal/domain/event"
	"github.com/go-chi/chi/v5"
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

func (c *EventController) FindOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
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

func (c *EventController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.Delete(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Delete(): %s", err)
			}
			return
		}
		err = (*c.service).Delete(id)
		err = success(w, "Successful deleted.")
		if err != nil {
			fmt.Printf("EventController.Delete(): %s", err)
		}
	}

}

func (c *EventController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var g event.St
		err := json.NewDecoder(r.Body).Decode(&g)
		if err != nil {
			fmt.Print("Error while reading request body")
		}
		res, er := (*c.service).Create(&g)
		if er != nil {
			fmt.Printf("EventController.Create(): %s", er)
		}
		err = created(w, res) // changed without testing
		if er != nil {
			fmt.Printf("EventController.Create(): %s", er)
		}

	}
}

func (c *EventController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var g event.St
		err := json.NewDecoder(r.Body).Decode(&g)
		fmt.Println(g)
		if err != nil {
			fmt.Print("Error while reading request body")
		}
		err = (*c.service).Update(&g)
		if err != nil {
			fmt.Printf("EventController.Update(): %s", err)
			err = notFound(w, err)
			if err != nil {
				fmt.Printf("EventController.Delete(): %s", err)
			}
		} else {
			err = success(w, "Successful updated.")
		}
	}
}
