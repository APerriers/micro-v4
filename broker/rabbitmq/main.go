package main

import (
	"fmt"
	"github.com/asim/go-micro/plugins/broker/rabbitmq/v4"
	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/server"
	"go-micro.dev/v4/util/log"
	"rabbitmq/handler"
	"time"
)

var (
	service = "go.micro.srv.rabbitmq"
	version = "latest"
)

func main() {
	rabbitmq.DefaultRabbitURL = "amqp://xtec:xtec0755@10.10.10.44:5672/test"

	bro := rabbitmq.NewBroker()
	bro.Init()
	if err := bro.Connect(); err != nil {
		log.Logf("cant conect to broker, skip: %v", err)
	}

	s := server.NewServer(server.Broker(bro))

	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Server(s),
		micro.Broker(bro),
	)

	//Register a subscriber
	brkrSub := broker.NewSubscribeOptions(
		broker.Queue("queue.default"),
		broker.DisableAutoAck(),
		rabbitmq.DurableQueue(),
	)

	h := new(handler.Rabbitmq)
	micro.RegisterSubscriber(
		"topic",
		srv.Server(),
		h.Handler,
		server.SubscriberContext(brkrSub.Context),
		server.SubscriberQueue("queue.default"),
	)

	//go pub("topic", bro)
	//go pub("topic1", bro)
	//go sub1("topic", bro)
	//go sub2("topic", bro)

	//srv.Init()

	// Register handler
	//pb.RegisterRabbitmqHandler(srv.Server(), new(handler.Rabbitmq))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}

// send events using the publisher
func pub(topic string, b broker.Broker) {
	t := time.NewTicker(time.Second)
	i := 0
	for _ = range t.C {
		// create new event
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(fmt.Sprintf("%d: %s", i, time.Now().String())),
		}

		//log.Logf("publishing %+v\n", msg)

		// publish an event
		if err := b.Publish(topic, msg); err != nil {
			log.Logf("[pub] failed: %v", err)
		} else {
			fmt.Println("[pub] pubbed message:", string(msg.Body))
		}

		i++
	}
}

func sub1(topic string, b broker.Broker) {
	_, err := b.Subscribe(topic, func(p broker.Event) error {
		fmt.Println("[sub1] received message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func sub2(topic string, b broker.Broker) {
	_, err := b.Subscribe(topic, func(p broker.Event) error {
		fmt.Println("[sub2] received message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
