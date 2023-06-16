package domain

import (
	"go_andr_less/21_http/utils"
	"time"
)

type Event struct {
	Id        int        `json:"id"`
	Content   string     `json:"content"`
	StartDate utils.Date `json:"start_date"`
}

type EventStore interface {
	Add(e *Event) error
	Update(e *Event) error
	Delete(e *Event) error
	Get(id int) (Event, error)
	List(dateFrom time.Time, dateTo time.Time) ([]Event, error)
	Count() int
}

var thisStore EventStore

func InitStore(store EventStore) *EventStore {
	thisStore = store
	return &thisStore
}

func Add(event *Event) error {
	return thisStore.Add(event)
}

func Update(event *Event) error {
	return thisStore.Update(event)
}

func Delete(e *Event) error {
	return thisStore.Delete(e)
}

func Get(id int) (Event, error) {
	return thisStore.Get(id)
}

func Count() int {
	return thisStore.Count()
}
