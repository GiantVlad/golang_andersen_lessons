package event

//go:generate protoc --go_out=. ./event.proto
type Event struct {
	Id      int
	Content string
}

type EventStore interface {
	Add(e Event)
	Update(e Event) (Event, error)
	Delete(e *Event) error
	Get(id int) (Event, error)
	Count() int
}

var thisStore EventStore

func InitStore(store EventStore) {
	thisStore = store
}

func Add(event Event) {
	thisStore.Add(event)
}

func Update(event Event) (Event, error) {
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
