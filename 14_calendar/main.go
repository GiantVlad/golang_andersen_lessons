package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/file"
	"go_andr_less/14_calendar/event"
	"go_andr_less/14_calendar/in_memory_store"
	"go_andr_less/14_calendar/logger"
	"go_andr_less/14_calendar/server"
	"log"
	"time"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "config/cfg.yml", "config file")
}

type Config struct {
	HttpCfg  server.HttpCfg `yaml:"http_listen"`
	LogFile  string         `yaml:"log_file"`
	LogLevel string         `yaml:"log_level"`
}

func main() {
	flag.Parse()
	loader := confita.NewLoader(
		file.NewBackend(configFile),
	)
	cfg := Config{}

	loader.Load(context.Background(), &cfg)

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

	//mess := pb.EventMess{
	//	Id:      12,
	//	Content: "something",
	//}
	//out, err := proto.Marshal(&mess)
	//if err != nil {
	//	log.Fatalln("Failed to encode address book:", err)
	//}
	//
	//fmt.Printf("\nCount: %v", out)
	//

	sugar := logger.NewLogger(cfg.LogFile, cfg.LogLevel)
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "/home",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", "/home")
	sugar.Warnf("Failed to fetch URL: %s", "/home")
	sugar.Errorf("Failed to fetch URL: %s", "/home")
	server.Start(cfg.HttpCfg)
}
