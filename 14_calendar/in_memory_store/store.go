package in_memory_store

import (
	"fmt"
	"go_andr_less/14_calendar/event"
	"sync"
)

type Events []event.Event

var mu = sync.Mutex{}

func (events *Events) Add(e event.Event) {
	*events = append(*events, e)
}

func (events Events) Get(id int) (event.Event, error) {
	for _, el := range events {
		if el.Id == id {
			return el, nil
		}
	}
	return event.Event{}, fmt.Errorf("not Found")
}

func (events Events) Update(e event.Event) (event.Event, error) {
	for _, el := range events {
		if el.Id == e.Id {
			mu.Lock()
			el = e
			mu.Unlock()
			return el, nil
		}
	}
	return event.Event{}, nil
}

func (events *Events) Delete(e *event.Event) error {
	for idx, el := range *events {
		if el.Id == e.Id {
			mu.Lock()
			*events = append((*events)[:idx], (*events)[idx+1:]...)
			mu.Unlock()
			return nil
		}
	}
	return fmt.Errorf("not Found")
}

func (events *Events) Count() int {
	mu.Lock()
	defer mu.Unlock()
	return len(*events)
}
