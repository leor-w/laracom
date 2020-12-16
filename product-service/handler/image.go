package handler

import (
	"context"
	"fmt"

	"github.com/leor-w/laracom/product-service/model"
	pb "github.com/leor-w/laracom/product-service/proto/product"
	"github.com/leor-w/laracom/product-service/repo"
)

type ImageService struct {
	ImageRepo repo.ImageRepositoryInterface
}

func (srv *ImageService) Get(ctx context.Context, req *pb.ProductImage, resp *pb.ImageResponse) error {
	if req.Id == 0 {
		return fmt.Errorf("商品图片 ID 不能为空")
	}
	image, err := srv.ImageRepo.GetById(uint(req.Id))
	if err != nil {
		return err
	}
	resp.Image = image.ToProtobuf()
	return nil
}

func (srv *ImageService) GetByProduct(ctx context.Context, req *pb.Product, resp *pb.ImageResponse) error {
	if req.Id == 0 {
		return fmt.Errorf("商品 ID 不能为空")
	}
	images, err := srv.ImageRepo.GetByProductId(uint(req.Id))
	if err != nil {
		return err
	}
	resp.Images = model.ProductImageToProtobufArray(images)
	return nil
}

func (srv *ImageService) Create(ctx context.Context, req *pb.ProductImage, resp *pb.ImageResponse) error {
	model := &model.ProductImage{}
	model.ToORM(req)
	err := srv.ImageRepo.Create(model)
	if err != nil {
		return err
	}
	resp.Image = model.ToProtobuf()
	return nil
}

func (srv ImageService) Update(ctx context.Context, req *pb.ProductImage, resp *pb.ImageResponse) error {
	model := &model.ProductImage{}
	model.ToORM(req)
	err := srv.ImageRepo.Update(model)
	if err != nil {
		return err
	}
	resp.Image = model.ToProtobuf()
	return nil
}

func (srv *ImageService) Delete(ctx context.Context, req *pb.ProductImage, resp *pb.ImageResponse) error {
	model := &model.ProductImage{}
	model.ToORM(req)
	err := srv.ImageRepo.Delete(model)
	if err != nil {
		return err
	}
	resp.Image = nil
	return nil
}
