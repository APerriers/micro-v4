package main

import (
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"redirect/handler"
	pb "redirect/proto"
)

var (
	service = "redirect"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	//srv.Server().Handle(srv.Server().NewHandler(new(handler.Redirect)))

	// Register handler
	pb.RegisterRedirectHandler(srv.Server(), new(handler.Redirect))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
