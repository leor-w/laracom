package main

import (
	"github.com/leor-w/laracom/product-service/db"
	"github.com/leor-w/laracom/product-service/handler"
	"github.com/leor-w/laracom/product-service/model"
	pb "github.com/leor-w/laracom/product-service/proto/product"
	"github.com/leor-w/laracom/product-service/repo"
	"github.com/micro/go-micro/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	database, err := db.CreateConnection()
	if err != nil {
		logrus.Fatalf("Could not connect to DB: %v", err)
	}

	err = database.AutoMigrate(&model.Product{})
	if err != nil {
		logrus.Fatalf("Migrate database failed: %v", err)
	}

	srv := micro.NewService(
		micro.Name("laracom.service.product"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterProductServiceHandler(
		srv.Server(),
		&handler.ProductService{
			ProductRepo: &repo.ProductRepository{
				Db: database,
			},
		})

	if err := srv.Run(); err != nil {
		logrus.Fatalf("Running Service failed: %v", err)
	}
}
