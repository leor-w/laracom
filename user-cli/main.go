package main

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	pb "github.com/leor-w/laracom/user-service/proto/user"
)

func main() {
	service := micro.NewService(
		micro.Flags(
			&cli.StringFlag{
				Name:  "name",
				Usage: "Your name",
			},
			&cli.StringFlag{
				Name:  "email",
				Usage: "Your email",
			},
			&cli.StringFlag{
				Name:  "password",
				Usage: "Your password",
			},
		),
	)

	client := pb.
}
