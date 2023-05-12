package in_memory_store

import (
	"github.com/stretchr/testify/require"
	"go_andr_less/14_calendar/event"
	"testing"
)

func TestStore(t *testing.T) {
	events := Events{}
	event1 := event.Event{
		Id:      1,
		Content: "v1 content",
	}
	event2 := event.Event{
		Id: 2,
	}
	event3 := event.Event{
		Id: 3,
	}
	events.Add(event1)
	events.Add(event2)
	events.Add(event3)
	require.Equal(t, 3, events.Count(), "")

	ev2, _ := events.Get(2)
	require.Equal(t, 2, ev2.Id, "")

	events.Delete(&ev2)
	require.Equal(t, 2, events.Count(), "")

	_, err := events.Get(2)
	require.NotNil(t, err, "")

	err = events.Delete(&ev2)
	require.NotNil(t, err, "")

	event1.Content = "v2 content"
	ev1, _ := events.Update(event1)
	require.Equal(t, "v2 content", ev1.Content, "")

	_, err = events.Update(ev2)
	require.NotNil(t, err, "")

	err = events.Add(ev1)
	require.NotNil(t, err, "")
}
