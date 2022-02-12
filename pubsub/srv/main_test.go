package main

import (
	"context"
	"fmt"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/metadata"
	"go-micro.dev/v4/selector"
	"go-micro.dev/v4/util/log"
	pb "srv/proto"
	"testing"
)

func TestCall(t *testing.T) {
	// Create new request to service go.micro.srv.example, method Example.Call
	req := client.NewRequest("go.micro.srv.pubsub", "Srv.Call", &pb.CallRequest{
		Name: "John",
	})

	// create context with metadata
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "john",
		"X-From-Id": "script",
	})

	rsp := &pb.CallResponse{}

	// Call service
	if err := client.Call(ctx, req, rsp); err != nil {
		fmt.Println("call err: ", err, rsp)
		return
	}
	fmt.Println("Call:", "rsp:", rsp.Msg)

}

func TestGetSrv(t *testing.T) {
	srv := micro.NewService(
		micro.Name("TestGetSrv"),
		micro.Version(version),
	)
	srv.Init()

	mdnsRegistry := srv.Options().Registry
	getService, err := mdnsRegistry.GetService("go.micro.srv.pubsub") // product 是之前注册到consule的服务名
	if err != nil {
		log.Fatal(err)
	}

	getServiceList, err := mdnsRegistry.ListServices()
	if err != nil {
		log.Fatal(err)
	}

	log.Info(getServiceList)
	log.Info(mdnsRegistry.String())

	next := selector.Random(getService) // 如果有多个服务，则随机调用一个
	node, err := next()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(node.Id, node.Address, node.Metadata)

	fmt.Sprintln(mdnsRegistry)
}
