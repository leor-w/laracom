package handler

import (
	"context"
	"fmt"

	"github.com/leor-w/laracom/product-service/model"
	pb "github.com/leor-w/laracom/product-service/proto/product"
	"github.com/leor-w/laracom/product-service/repo"
	"gorm.io/gorm"
)

type ProductService struct {
	ProductRepo repo.ProductRepositoryInterface
}

func (srv *ProductService) Create(ctx context.Context, req *pb.Product, resp *pb.Response) error {
	productModel := &model.Product{}
	productModel.ToORM(req)
	if err := srv.ProductRepo.Create(productModel); err != nil {
		return err
	}
	resp.Product = productModel.ToProtobuf()
	return nil
}

func (srv *ProductService) Delete(ctx context.Context, req *pb.Product, resp *pb.Response) error {
	productModel := &model.Product{}
	productModel.ToORM(req)
	if err := srv.ProductRepo.Delete(productModel); err != nil {
		return err
	}
	resp.Product = productModel.ToProtobuf()
	return nil
}

func (srv *ProductService) Update(ctx context.Context, req *pb.Product, resp *pb.Response) error {
	productModel := &model.Product{}
	productModel.ToORM(req)
	if err := srv.ProductRepo.Update(productModel); err != nil {
		return err
	}
	resp.Product = productModel.ToProtobuf()
	return nil
}

func (srv *ProductService) Get(ctx context.Context, req *pb.Product, resp *pb.Response) error {
	if req.Id == 0 && req.Slug == "" {
		return fmt.Errorf("商品 ID 与 Slug 不能都为空")
	}
	var (
		productModel *model.Product
		err          error
	)
	if req.Id != 0 {
		productModel, err = srv.ProductRepo.GetById(uint(req.Id))
	} else {
		productModel, err = srv.ProductRepo.GetBySlug(req.Slug)
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if productModel != nil {
		resp.Product = productModel.ToProtobuf()
	}
	return nil
}

func (srv *ProductService) GetDetail(ctx context.Context, req *pb.Product, resp *pb.Response) error {
	return nil
}

func (srv *ProductService) GetAll(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	productModels, err := srv.ProductRepo.GetAll()
	if err != nil {
		return err
	}
	var products []*pb.Product
	for _, productModel := range productModels {
		product := productModel.ToProtobuf()
		products = append(products, product)
	}
	resp.Products = products
	return nil
}
