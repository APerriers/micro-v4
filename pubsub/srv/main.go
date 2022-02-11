package main

import (
	"context"
	"fmt"
	proto "go-micro.dev/v4/api/proto"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/metadata"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
	"go-micro.dev/v4/server"
	"srv/handler"
	pb "srv/proto"
	"time"

	"go-micro.dev/v4"
	"go-micro.dev/v4/util/log"
)

var (
	service = "go.micro.srv.pubsub"
	version = "latest"
)

func logWrapper(handlerFunc server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Logf("[Log Wrapper] Before serving request method: %v", req.Endpoint())
		md, _ := metadata.FromContext(ctx)
		fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())

		err := handlerFunc(ctx, req, rsp)
		log.Logf("[Log Wrapper] After serving request")
		return err
	}
}

func logSubWrapper(subscriberFunc server.SubscriberFunc) server.SubscriberFunc {
	return func(ctx context.Context, msg server.Message) error {
		log.Logf("[Log Sub Wrapper] Before serving publication topic: %v", msg.Topic())
		err := subscriberFunc(ctx, msg)
		log.Log("[Log Sub Wrapper] After serving publication")
		return err
	}
}

func metricsWrap(callFunc client.CallFunc) client.CallFunc {
	return func(ctx context.Context, node *registry.Node, req client.Request, rsp interface{}, opts client.CallOptions) error {
		t := time.Now()
		err := callFunc(ctx, node, req, rsp, opts)
		log.Logf("[Metrics Wrapper] called: %v %s.%s duration: %v\n", node, req.Service(), req.Endpoint(), time.Since(t))
		return err
	}
}

// Return a new first node selector
func FirstNodeSelector(opts ...selector.Option) selector.Selector {
	var sopts selector.Options
	for _, opt := range opts {
		opt(&sopts)
	}
	if sopts.Registry == nil {
		sopts.Registry = registry.DefaultRegistry
	}
	return &handler.FirstNodeSelector{sopts}
}

func main() {
	md := server.DefaultOptions().Metadata
	md["datacenter"] = "local1111111"

	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.WrapHandler(logWrapper),
		micro.WrapSubscriber(logSubWrapper),
		micro.WrapCall(metricsWrap),
		micro.Metadata(md),
		micro.Selector(FirstNodeSelector()),
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
