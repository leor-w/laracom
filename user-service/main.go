package main

import (
	"fmt"
	"log"

	database "github.com/leor-w/laracom/user-service/db"
	"github.com/leor-w/laracom/user-service/handler"
	pb "github.com/leor-w/laracom/user-service/proto/user"
	repository "github.com/leor-w/laracom/user-service/repo"
	"github.com/micro/go-micro/v2"
)

func main() {
	db, err := database.CreateConnection()
	defer db.Close()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	db.AutoMigrate(&pb.User{})

	repo := &repository.UserRepository{db}

	srv := micro.NewService(
		micro.Name("laracom.user.service"),
		micro.Version("latest"),
	)
	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), &handler.UserService{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
