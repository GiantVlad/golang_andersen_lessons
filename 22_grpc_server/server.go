package main

import (
	"context"
	grpcserver "go_andr_less/22_grpc_server/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	grpcServer := grpc.NewServer()
	grpcserver.RegisterCreateEventServer(grpcServer, CreateEvent{})
	grpcserver.RegisterGetListServer(grpcServer, GetList{})
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

type GetList struct {
}

func (g GetList) List(ctx context.Context, timestamp *timestamppb.Timestamp) (*grpcserver.Events, error) {
	ev1 := grpcserver.Event{Id: 1, Content: "something"}
	ev2 := grpcserver.Event{Id: 1, Content: "something"}
	events := []*grpcserver.Event{
		&ev1,
		&ev2,
	}

	return &grpcserver.Events{Events: events}, nil
}

type CreateEvent struct {
}

func (c CreateEvent) Create(ctx context.Context, event *grpcserver.Event) (*grpcserver.Event, error) {
	return event, nil
}
