package main

import (
	"fmt"
	"log"

	database "github.com/leor-w/laracom/user-service/db"
	"github.com/leor-w/laracom/user-service/handler"
	pb "github.com/leor-w/laracom/user-service/proto/user"
	repository "github.com/leor-w/laracom/user-service/repo"
	"github.com/leor-w/laracom/user-service/service"
	"github.com/micro/go-micro/v2"
)

func main() {
	db, err := database.CreateConnection()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	db.AutoMigrate(&pb.User{})

	repo := &repository.UserRepository{db}
	token := &service.TokenService{repo}

	srv := micro.NewService(
		micro.Name("laracom.service.user"),
		micro.Version("latest"),
	)
	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), &handler.UserService{repo, token})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
