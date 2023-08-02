package main

import (
	"log"
)

type NMessage struct {
	Level   int
	Message string
}

type Notifier interface {
	Notify(m NMessage) error
}

type CliNotifier struct {
}

func (cliNotifier CliNotifier) Notify(message NMessage) error {
	log.Printf("Sent notification: %s, level: %d \n", message.Message, message.Level)
	return nil
}
