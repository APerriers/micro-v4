package main

import (
	"fmt"
	"github.com/asim/go-micro/plugins/broker/rabbitmq/v4"
	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/server"
	"go-micro.dev/v4/util/log"
)

var (
	service = "consumer"
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

	//brkrSub := broker.NewSubscribeOptions(
	//	broker.Queue("queue.default"),
	//	broker.DisableAutoAck(),
	//	rabbitmq.DurableQueue(),
	//)

	//h := new(handler.Consumer)
	//micro.RegisterSubscriber(
	//	"topic",
	//	srv.Server(),
	//	h.Handler,
	//	server.SubscriberContext(brkrSub.Context),
	//	server.SubscriberQueue("queue.default"),
	//)

	go sub1("topic", bro)
	go sub2("topic1", bro)

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
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
