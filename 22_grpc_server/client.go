package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	grpcserver "go_andr_less/22_grpc_server/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"os"
	"strings"
	"time"
)

func writeRoutine(end chan interface{}, ctx context.Context, connCreate grpcserver.CreateEventClient, connList grpcserver.GetListClient) {
	scanner := bufio.NewScanner(os.Stdin)
OUTER:
	for {
		select {
		case <-ctx.Done():
			break OUTER
		default:
			if !scanner.Scan() {
				break OUTER
			}
			input := scanner.Text()
			if input == "exit" {
				break OUTER
			}
			request := Request{}
			err := json.Unmarshal([]byte(input), &request)
			if err != nil {
				fmt.Printf("error: %s\n", status.Convert(err).Message())
				break
			}
			if request.Source == "create" {
				event := grpcserver.Event{}
				err = jsonpb.Unmarshal(strings.NewReader(request.Data), &event)
				if err != nil {
					fmt.Printf("error: %s\n", status.Convert(err).Message())
					break
				}
				msg, err := connCreate.Create(context.Background(), &event)
				if err != nil {
					fmt.Printf("error: %s\n", status.Convert(err).Message())
				}

				if msg != nil {
					created := msg.StartDate.AsTime()
					fmt.Printf("[%s]id:%d msg:%s\n", created.Local(), msg.Id, msg.Content)
				}
			} else {
				date := timestamppb.Timestamp{}
				err = jsonpb.Unmarshal(strings.NewReader(request.Data), &date)
				if err != nil {
					fmt.Printf("error: %s\n", status.Convert(err).Message())
					break
				}
				msg, err := connList.List(context.Background(), &date)
				if err != nil {
					fmt.Printf("error: %s\n", status.Convert(err).Message())
				}

				if msg != nil {
					count := msg.Events
					fmt.Printf("msg:%v\n", count)
				}
			}
		}

	}
	log.Printf("Finished writeRoutine")
	close(end)
}

type Request struct {
	Source string
	Data   string
}

// Request example: {"content":"some value","start_date":"2023-01-09T00:00:00Z"}
func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)

	cc, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer func(cc *grpc.ClientConn) {
		err := cc.Close()
		if err != nil {
			log.Fatalf("could not close connection: %v", err)
		}
	}(cc)
	c := grpcserver.NewCreateEventClient(cc)
	c2 := grpcserver.NewGetListClient(cc)
	end := make(chan interface{})
	go writeRoutine(end, ctx, c, c2)

	<-end
}
