package main

import (
	"context"
	"fmt"
	"github.com/pborman/uuid"
	"go-micro.dev/v4"
	"go-micro.dev/v4/util/log"
	proto "pubsub-srv/proto"
	"time"
)

var (
	service = "go.micro.cli.pubsub"
	version = "latest"
)

func main() {
	// create a service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	// parse command line
	srv.Init()

	// create publisher
	pub1 := micro.NewEvent("example.topic.pubsub.1", srv.Client())
	pub2 := micro.NewEvent("example.topic.pubsub.2", srv.Client())

	// pub to topic 1
	go sendEv("example.topic.pubsub.1", pub1)
	// pub to topic 2
	go sendEv("example.topic.pubsub.2", pub2)

	// block forever
	select {}
}

// send events using the publisher
func sendEv(topic string, p micro.Publisher) {
	t := time.NewTicker(time.Second)

	for _ = range t.C {
		// create new event
		ev := &proto.Event{
			Id:        uuid.NewUUID().String(),
			Timestamp: time.Now().Unix(),
			Message:   fmt.Sprintf("Messaging you all day on %s", topic),
		}

		log.Logf("publishing %+v\n", ev)

		// publish an event
		if err := p.Publish(context.Background(), ev); err != nil {
			log.Logf("error publishing: %v", err)
		}
	}
}
