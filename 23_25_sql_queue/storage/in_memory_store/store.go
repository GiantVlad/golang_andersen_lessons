package in_memory_store

import (
	"fmt"
	"go_andr_less/23_25_sql_queue/http/domain"
	"sync"
	"time"
)

type Events []domain.Event

var mu = sync.Mutex{}

func (events *Events) Add(e *domain.Event) error {
	newId := 1
	for _, el := range *events {
		if el.Id == e.Id {
			return fmt.Errorf("the item #%d already exists", e.Id)
		}
		if newId <= el.Id {
			newId = el.Id + 1
		}
	}
	e.Id = newId
	*events = append(*events, *e)
	return nil
}

func (events Events) Get(id int) (domain.Event, error) {
	for _, el := range events {
		if el.Id == id {
			return el, nil
		}
	}
	return domain.Event{}, fmt.Errorf("not found")
}

func (events Events) Update(e *domain.Event) error {
	for i, el := range events {
		if el.Id == (*e).Id {
			mu.Lock()
			events[i] = *e
			mu.Unlock()
			return nil
		}
	}
	return fmt.Errorf("not found")
}

func (events *Events) Delete(e *domain.Event) error {
	for idx, el := range *events {
		if el.Id == e.Id {
			mu.Lock()
			*events = append((*events)[:idx], (*events)[idx+1:]...)
			mu.Unlock()
			return nil
		}
	}
	return fmt.Errorf("not found")
}

func (events *Events) List(dateFrom time.Time, dateTo time.Time) ([]domain.Event, error) {
	mu.Lock()
	defer mu.Unlock()
	var filtered []domain.Event
	var evs []domain.Event
	evs = *events
	for i := range evs {
		if evs[i].StartDate.Time.Equal(dateFrom) || (evs[i].StartDate.Time.After(dateFrom) && evs[i].StartDate.Time.Before(dateTo)) {
			filtered = append(filtered, evs[i])
		}
	}
	return filtered, nil
}

func (events *Events) Count() int {
	mu.Lock()
	defer mu.Unlock()
	return len(*events)
}
