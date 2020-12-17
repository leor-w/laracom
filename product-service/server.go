package main

import (
	"net/http"

	"github.com/leor-w/laracom/product-service/db"
	"github.com/leor-w/laracom/product-service/handler"
	"github.com/leor-w/laracom/product-service/model"
	pb "github.com/leor-w/laracom/product-service/proto/product"
	"github.com/leor-w/laracom/product-service/repo"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

func prometheusBoot() {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(":9092", nil)
		if err != nil {
			logrus.Fatalf("ListenAndServer: %v", err)
		}
	}()
}

func main() {
	database, err := db.CreateConnection()
	if err != nil {
		logrus.Fatalf("Could not connect to DB: %v", err)
	}

	err = database.AutoMigrate(&model.Product{})
	err = database.AutoMigrate(&model.ProductImage{})
	err = database.AutoMigrate(&model.Brand{})
	err = database.AutoMigrate(&model.Category{})
	err = database.AutoMigrate(&model.Attribute{})
	err = database.AutoMigrate(&model.AttributeValue{})
	err = database.AutoMigrate(&model.ProductAttribute{})
	if err != nil {
		logrus.Fatalf("Migrate database failed: %v", err)
	}

	srv := micro.NewService(
		micro.Name("laracom.service.product"),
		micro.Version("latest"),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)

	srv.Init()

	pb.RegisterProductServiceHandler(srv.Server(), &handler.ProductService{ProductRepo: &repo.ProductRepository{Db: database}})
	pb.RegisterImageServiceHandler(srv.Server(), &handler.ImageService{ImageRepo: &repo.ImageRepository{Db: database}})
	pb.RegisterBrandServiceHandler(srv.Server(), &handler.BrandService{BrandRepo: &repo.BrandRepository{Db: database}})
	pb.RegisterCategoryServiceHandler(srv.Server(), &handler.CategoryService{CategoryRepo: &repo.CategoryRepository{Db: database}})
	pb.RegisterAttributeServiceHandler(srv.Server(), &handler.AttributeService{AttributeRepo: &repo.AttributeRepository{Db: database}})

	prometheusBoot()

	if err := srv.Run(); err != nil {
		logrus.Fatalf("Running Service failed: %v", err)
	}
}
