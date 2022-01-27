package main

import (
	"job/handler"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "job"
	version = "latest"
)

func main() {
	// Create function
	fnc := micro.NewFunction(
		micro.Name(service),
		micro.Version(version),
	)
	fnc.Init()

	// Handle function
	fnc.Handle(new(handler.Job))

	// Run function
	if err := fnc.Run(); err != nil {
		log.Fatal(err)
	}
}
