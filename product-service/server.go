package main

import (
	"net/http"
	"os"

	"github.com/leor-w/laracom/product-service/db"
	"github.com/leor-w/laracom/product-service/handler"
	"github.com/leor-w/laracom/product-service/model"
	pb "github.com/leor-w/laracom/product-service/proto/product"
	"github.com/leor-w/laracom/product-service/repo"
	"github.com/leor-w/laracom/product-service/trace"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	tracePlugin "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

const ProductServiceName = "laracom.service.product"

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

	tracer, close, err := trace.NewTracer(ProductServiceName, os.Getenv("MICRO_TRACE_SERVER"))
	if err != nil {
		logrus.Fatalf("create trace failed : %v", err)
	}
	defer close.Close()
	opentracing.SetGlobalTracer(tracer)

	srv := micro.NewService(
		micro.Name(ProductServiceName),
		micro.Version("latest"),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
		micro.WrapHandler(tracePlugin.NewHandlerWrapper(opentracing.GlobalTracer())),
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
