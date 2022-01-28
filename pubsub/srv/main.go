package main

import (
	"context"
	proto "go-micro.dev/v4/api/proto"
	"go-micro.dev/v4/metadata"
	"go-micro.dev/v4/server"
	"srv/handler"
	pb "srv/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/util/log"
)

var (
	service = "go.micro.srv.pubsub"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	pb.RegisterSrvHandler(srv.Server(), new(handler.Srv))

	// register subscriber
	micro.RegisterSubscriber("example.topic.pubsub.1", srv.Server(), new(handler.Sub))

	// register subscriber with queue, each message is delivered to a unique subscriber
	micro.RegisterSubscriber("example.topic.pubsub.2", srv.Server(), subEv, server.SubscriberQueue("queue.pubsub"))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}

// Alternatively a function can be used
func subEv(ctx context.Context, event *proto.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Logf("[pubsub.2] Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}
