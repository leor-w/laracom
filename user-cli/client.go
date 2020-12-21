package main

import (
	"context"
	"os"

	pb "github.com/leor-w/laracom/user-service/proto/user"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/sirupsen/logrus"
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

	client := pb.NewUserService("laracom.service.user", service.Client())
	service.Init(
		micro.Action(func(c *cli.Context) error {
			name := c.String("name")
			email := c.String("email")
			password := c.String("password")

			logrus.Infof("Request params: name [%s] email [%s] password [%s]", name, email, password)

			r, err := client.Create(context.TODO(), &pb.User{
				Name:     name,
				Email:    email,
				Password: password,
			})
			if err != nil {
				logrus.Fatalf("创建用户失败: %v", err)
			}
			logrus.Infof("创建用户成功: %s", r.User.Id)

			token, err := client.Auth(context.TODO(), &pb.User{
				Email:    email,
				Password: password,
			})

			if err != nil {
				logrus.Errorf("用户登陆失败：%v", err)
			}
			logrus.Infof("用户登陆成功：%s", token.Token)

			token, err = client.ValidateToken(context.TODO(), token)
			if err != nil {
				logrus.Errorf("用户验证失败: token [%v] error [%v]", token, err)
			}
			logrus.Infof("用户验证成功：%v", token.Valid)

			getAll, err := client.GetAll(context.Background(), &pb.Request{})
			if err != nil {
				logrus.Fatalf("获取所有用户失败: %v", err)
			}
			for _, v := range getAll.Users {
				logrus.Println(v)
			}
			os.Exit(0)
			return nil
		}),
	)

	if err := service.Run(); err != nil {
		logrus.Fatalf("用户端启动失败: %v", err)
	}
}
