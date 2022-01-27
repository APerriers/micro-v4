package main

import (
	"apiuser/handler"
	pb "apiuser/proto"
	"fmt"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"time"
)

var (
	service = "go.micro.srv.apiuser"
	version = "latest"
)

func main() {
	// Create service
	var srv = micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		//程序参数
		micro.Flags(
			&cli.StringFlag{
				Name:  "string_flag",
				Usage: "This is a string flag",
			},
			&cli.IntFlag{
				Name:  "int_flag",
				Usage: "This is an int flag",
			},
			&cli.BoolFlag{
				Name:  "bool_flag",
				Usage: "This is a bool flag",
			},
		),
	)
	srv.Init(
		micro.Action(func(c *cli.Context) error {
			fmt.Printf("The string flag is: %s\n", c.String("string_flag"))
			fmt.Printf("The int flag is: %d\n", c.Int("int_flag"))
			fmt.Printf("The bool flag is: %t\n", c.Bool("bool_flag"))
			// let's just exit because
			//os.Exit(0)
			return nil
		}))

	// Register handler
	pb.RegisterApiuserHandler(srv.Server(), new(handler.Apiuser))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
