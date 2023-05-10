package main

import (
	"fmt"
	"go_andr_less/14_calendar/event"
	"go_andr_less/14_calendar/event/pb"
	"go_andr_less/14_calendar/in_memory_store"
	"google.golang.org/protobuf/proto"
	"log"
)

func main() {
	event.InitStore(&in_memory_store.Events{})
	ev1 := event.Event{
		Id:      1,
		Content: "go to bed",
	}
	ev2 := event.Event{
		Id:      2,
		Content: "week up",
	}
	ev3 := event.Event{
		Id:      3,
		Content: "to have a dinner",
	}
	event.Add(ev1)
	event.Add(ev2)
	event.Add(ev3)
	err := event.Delete(&ev3)
	if err != nil {
		log.Fatalln(err)
	}
	ev2, err = event.Get(2)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Event: %s", ev2.Content)

	ev2.Content = "You can sleep, this is day off"
	ev2, err = event.Update(ev2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("\nEvent: %s", ev2.Content)

	fmt.Printf("\nCount: %d", event.Count())

	mess := pb.EventMess{
		Id:      12,
		Content: "something",
	}
	out, err := proto.Marshal(&mess)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}

	fmt.Printf("\nCount: %v", out)
}
