package main

import (
	"fmt"
	"log"

	database "github.com/leor-w/laracom/user-service/db"
	"github.com/leor-w/laracom/user-service/handler"
	"github.com/leor-w/laracom/user-service/model"
	pb "github.com/leor-w/laracom/user-service/proto/user"
	repository "github.com/leor-w/laracom/user-service/repo"
	"github.com/leor-w/laracom/user-service/service"
	"github.com/micro/go-micro/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	db, err := database.CreateConnection()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	err = db.AutoMigrate(&model.User{})
	err = db.AutoMigrate(&model.PasswordReset{})
	if err != nil {
		logrus.Errorf("DatabaseAutoMigrate failed : %v", err)
		return
	}

	// 初始化 Repo 实例
	repo := &repository.UserRepository{db}
	passwordRepo := &repository.PasswordRepository{db}
	token := &service.TokenService{repo}

	// Micro 注册微服务
	srv := micro.NewService(
		micro.Name("laracom.service.user"),
		micro.Version("latest"),
	)
	srv.Init()

	pubSub := srv.Server().Options().Broker
	err = pubSub.Connect()
	defer pubSub.Disconnect()
	if err != nil {
		logrus.Fatalf("Connect nats failed : %v", err)
	}

	// 注册处理器
	pb.RegisterUserServiceHandler(
		srv.Server(),
		&handler.UserService{
			Repo:      repo,
			ResetRepo: passwordRepo,
			Token:     token,
			PubSub:    pubSub,
		},
	)

	// 启动服务
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
