package handler

import (
	"context"
	"fmt"

	"github.com/leor-w/laracom/product-service/model"
	pb "github.com/leor-w/laracom/product-service/proto/product"
	"github.com/leor-w/laracom/product-service/repo"
	"github.com/sirupsen/logrus"
)

type BrandService struct {
	BrandRepo repo.BrandRepositoryInterface
}

func (srv *BrandService) Create(ctx context.Context, req *pb.Brand, resp *pb.BrandResponse) error {
	model := &model.Brand{}
	model.ToORM(req)
	err := srv.BrandRepo.Create(model)
	if err != nil {
		return err
	}
	resp.Brand = model.ToProtobuf()
	return nil
}

func (srv *BrandService) Update(ctx context.Context, req *pb.Brand, resp *pb.BrandResponse) error {
	model := &model.Brand{}
	model.ToORM(req)
	err := srv.BrandRepo.Update(model)
	if err != nil {
		return err
	}
	resp.Brand = model.ToProtobuf()
	return nil
}

func (srv *BrandService) Delete(ctx context.Context, req *pb.Brand, resp *pb.BrandResponse) error {
	model := &model.Brand{}
	model.ToORM(req)
	err := srv.BrandRepo.Delete(model)
	if err != nil {
		return err
	}
	resp.Brand = nil
	return nil
}

func (srv *BrandService) Get(ctx context.Context, req *pb.Brand, resp *pb.BrandResponse) error {
	if req.Id == 0 {
		return fmt.Errorf("not find brand with id : id = %d", req.Id)
	}
	brand, err := srv.BrandRepo.GetById(uint(req.Id))
	if err != nil {
		return err
	}
	resp.Brand = brand.ToProtobuf()
	return nil
}

func (srv *BrandService) GetAll(ctx context.Context, req *pb.Request, resp *pb.BrandResponse) error {
	brands, err := srv.BrandRepo.GetAll()
	if err != nil {
		return err
	}
	resp.Brands = model.BrandToProtobufArray(brands)
	return nil
}

func (srv *BrandService) GetWithProducts(ctx context.Context, req *pb.Brand, resp *pb.BrandResponse) error {
	if req.Id == 0 {
		logrus.Errorf("brand id not be empty")
		return fmt.Errorf("品牌 ID 不能为空")
	}
	model, err := srv.BrandRepo.GetWithProduct(uint(req.Id))
	if err != nil {
		logrus.Errorf("not find brand with product record : brandId = [%v], err = %v", req.Id, err)
		return err
	}
	resp.Brand = model.ToProtobuf()
	return nil
}
