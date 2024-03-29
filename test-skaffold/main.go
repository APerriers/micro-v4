package main

import (
	"test-skaffold/handler"
	pb "test-skaffold/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "test-skaffold"
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
	pb.RegisterTestSkaffoldHandler(srv.Server(), new(handler.TestSkaffold))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
