package main

import (
	"context"
	"fmt"
	"go-micro.dev/v4"
	api "go-micro.dev/v4/api/proto"
	"go-micro.dev/v4/server"
	"log"
	"redirect/handler"
	pb "redirect/proto"
	"sync"
	"testing"
	"time"
)

func TestRedirect(t *testing.T) {
	service := micro.NewService()
	service.Init()
	c := service.Client()

	request := c.NewRequest("redirect", "Redirect.Url", api.Request{
		Method: "1",
		Path:   "/11",
		Header: nil,
		Get:    nil,
		Post:   nil,
		Body:   "",
		Url:    "tst",
	})
	var response = new(api.Response)

	if err := c.Call(context.TODO(), request, &response); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)
}

func TestCancelSrv(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-time.After(time.Second * 15)
		log.Println("Shutdown example: shutting down service")
		cancel()
	}()

	// create service
	service := micro.NewService(
		// with our cancellation context
		micro.Context(ctx),
	)

	// init service
	service.Init()

	// run service
	service.Run()
}

func waitgroup(wg *sync.WaitGroup) server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			log.Println("waitgroup", rsp)
			wg.Add(1)
			defer wg.Done()
			return h(ctx, req, rsp)
		}
	}
}

func TestWait(t *testing.T) {
	var wg sync.WaitGroup

	service := micro.NewService(
		micro.Name("wait"),
		// wrap handlers with waitgroup wrapper
		micro.WrapHandler(waitgroup(&wg)),
		// waits for the waitgroup once stopped
		micro.AfterStop(func() error {
			// wait for handlers to finish
			wg.Wait()
			log.Println("service has stopped")
			return nil
		}),
	)

	pb.RegisterRedirectHandler(service.Server(), new(handler.Redirect))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func TestWaitGroup(t *testing.T) {
	service := micro.NewService()
	service.Init()
	c := service.Client()

	request := c.NewRequest("wait", "Redirect.Call", pb.CallRequest{Name: "john"})
	var response = new(pb.CallResponse)

	if err := c.Call(context.TODO(), request, &response); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)
}
