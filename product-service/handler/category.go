package handler

import (
	"context"
	"fmt"

	"github.com/leor-w/laracom/product-service/model"
	pb "github.com/leor-w/laracom/product-service/proto/product"
	"github.com/leor-w/laracom/product-service/repo"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type CategoryService struct {
	CategoryRepo repo.CategoryRepositoryInterface
}

func (srv *CategoryService) Create(ctx context.Context, req *pb.Category, resp *pb.CategoryResponse) error {
	model := &model.Category{}
	model.ToORM(req)
	err := srv.CategoryRepo.Create(model)
	if err != nil {
		return err
	}
	resp.Category = model.ToProtobuf()
	return nil
}

func (srv *CategoryService) Update(ctx context.Context, req *pb.Category, resp *pb.CategoryResponse) error {
	model := &model.Category{}
	model.ToORM(req)
	err := srv.CategoryRepo.Update(model)
	if err != nil {
		return err
	}
	resp.Category = model.ToProtobuf()
	return nil
}

func (srv *CategoryService) Delete(ctx context.Context, req *pb.Category, resp *pb.CategoryResponse) error {
	model := &model.Category{}
	model.ToORM(req)
	err := srv.CategoryRepo.Delete(model)
	if err != nil {
		return err
	}
	resp.Category = nil
	return nil
}

func (srv *CategoryService) Get(ctx context.Context, req *pb.Category, resp *pb.CategoryResponse) error {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}
	var sp opentracing.Span
	wireContext, _ := opentracing.GlobalTracer().Extract(opentracing.TextMap, opentracing.TextMapCarrier(md))
	sp = opentracing.StartSpan("GetCategory", opentracing.ChildOf(wireContext))
	sp.SetTag("req", req)
	defer func() {
		sp.SetTag("resp", resp)
		sp.Finish()
	}()
	if req.Id == 0 {
		logrus.Errorf("category id not be empty")
		return fmt.Errorf("分类 ID 不能为空")
	}
	category, err := srv.CategoryRepo.GetById(uint(req.Id))
	if err != nil {
		return err
	}
	resp.Category = category.ToProtobuf()
	return nil
}

func (srv *CategoryService) GetAll(ctx context.Context, req *pb.Request, resp *pb.CategoryResponse) error {
	categories, err := srv.CategoryRepo.GetAll()
	if err != nil {
		return err
	}
	resp.Categories = model.CategoryToProtobufArray(categories)
	return nil
}

func (srv *CategoryService) GetWithProducts(ctx context.Context, req *pb.Category, resp *pb.CategoryResponse) error {
	if req.Id == 0 {
		logrus.Errorf("category id not be empty")
		return fmt.Errorf("分类 ID 不能为空")
	}
	category, err := srv.CategoryRepo.GetWithProducts(uint(req.Id))
	if err != nil {
		return err
	}
	resp.Category = category.ToProtobuf()
	return nil
}
