package handler

import (
	"context"
	"fmt"

	"github.com/leor-w/laracom/product-service/model"
	pb "github.com/leor-w/laracom/product-service/proto/product"
	"github.com/leor-w/laracom/product-service/repo"
	"github.com/sirupsen/logrus"
)

type AttributeService struct {
	AttributeRepo repo.AttributeRepositoryInterface
}

func (srv *AttributeService) CreateAttribute(ctx context.Context, req *pb.Attribute, resp *pb.AttributeResponse) error {
	model := &model.Attribute{}
	model.ToORM(req)
	err := srv.AttributeRepo.CreateAttribute(model)
	if err != nil {
		return err
	}
	resp.Attribute = model.ToProtobuf()
	return nil
}

func (srv *AttributeService) DeleteAttribute(ctx context.Context, req *pb.Attribute, resp *pb.AttributeResponse) error {
	model := &model.Attribute{}
	model.ToORM(req)
	err := srv.AttributeRepo.DeleteAttribute(model)
	if err != nil {
		return err
	}
	resp.Attribute = nil
	return nil
}

func (srv *AttributeService) UpdateAttribute(ctx context.Context, req *pb.Attribute, resp *pb.AttributeResponse) error {
	model := &model.Attribute{}
	model.ToORM(req)
	err := srv.AttributeRepo.UpdateAttribute(model)
	if err != nil {
		return err
	}
	resp.Attribute = model.ToProtobuf()
	return nil
}

func (srv *AttributeService) CreateValue(ctx context.Context, req *pb.AttributeValue, resp *pb.AttributeValueResponse) error {
	model := &model.AttributeValue{}
	model.ToORM(req)
	err := srv.AttributeRepo.CreateValue(model)
	if err != nil {
		return err
	}
	resp.Value = model.ToProtobuf()
	return nil
}

func (srv *AttributeService) DeleteValue(ctx context.Context, req *pb.AttributeValue, resp *pb.AttributeValueResponse) error {
	model := &model.AttributeValue{}
	model.ToORM(req)
	err := srv.AttributeRepo.DeleteValue(model)
	if err != nil {
		return err
	}
	resp.Value = nil
	return nil
}

func (srv *AttributeService) UpdateValue(ctx context.Context, req *pb.AttributeValue, resp *pb.AttributeValueResponse) error {
	model := &model.AttributeValue{}
	model.ToORM(req)
	err := srv.AttributeRepo.UpdateValue(model)
	if err != nil {
		return err
	}
	resp.Value = model.ToProtobuf()
	return nil
}

func (srv *AttributeService) CreateProductAttribute(ctx context.Context, req *pb.ProductAttribute, resp *pb.ProductAttributeResponse) error {
	model := &model.ProductAttribute{}
	model.ToORM(req)
	err := srv.AttributeRepo.CreateProductAttribute(model)
	if err != nil {
		return err
	}
	resp.ProductAttribute = model.ToProtobuf()
	return nil
}

func (srv *AttributeService) DeleteProductAttribute(ctx context.Context, req *pb.ProductAttribute, resp *pb.ProductAttributeResponse) error {
	model := &model.ProductAttribute{}
	model.ToORM(req)
	err := srv.AttributeRepo.DeleteProductAttribute(model)
	if err != nil {
		return err
	}
	resp.ProductAttribute = nil
	return nil
}

func (srv *AttributeService) UpdateProductAttribute(ctx context.Context, req *pb.ProductAttribute, resp *pb.ProductAttributeResponse) error {
	model := &model.ProductAttribute{}
	model.ToORM(req)
	err := srv.AttributeRepo.UpdateProductAttribute(model)
	if err != nil {
		return err
	}
	resp.ProductAttribute = model.ToProtobuf()
	return nil
}

func (srv *AttributeService) GetAttribute(ctx context.Context, req *pb.Attribute, resp *pb.AttributeResponse) error {
	if req.Id == 0 {
		logrus.Errorf("attribute id not be empty")
		return fmt.Errorf("属性 ID 不能为空")
	}
	attribute, err := srv.AttributeRepo.GetAttribute(uint(req.Id))
	if err != nil {
		return err
	}
	resp.Attribute = attribute.ToProtobuf()
	return nil
}

func (srv *AttributeService) GetAttributes(ctx context.Context, req *pb.Request, resp *pb.AttributeResponse) error {
	attributes, err := srv.AttributeRepo.GetAttributes()
	if err != nil {
		return err
	}
	resp.Attributes = model.AttributeToProtobufArray(attributes)
	return nil
}

func (srv *AttributeService) GetValue(ctx context.Context, req *pb.AttributeValue, resp *pb.AttributeValueResponse) error {
	if req.Id == 0 {
		logrus.Errorf("attribute value id not be empty")
		return fmt.Errorf("属性值 ID 不能为空")
	}
	value, err := srv.AttributeRepo.GetValue(uint(req.Id))
	if err != nil {
		return err
	}
	resp.Value = value.ToProtobuf()
	return nil
}

func (srv *AttributeService) GetValues(ctx context.Context, req *pb.Request, resp *pb.AttributeValueResponse) error {
	values, err := srv.AttributeRepo.GetValues()
	if err != nil {
		return err
	}
	resp.Values = model.AttributeValueToProtobufArray(values)
	return nil
}

func (srv *AttributeService) GetProductAttribute(ctx context.Context, req *pb.ProductAttribute, resp *pb.ProductAttributeResponse) error {
	if req.Id == 0 {
		logrus.Errorf("product attribute id not be empty")
		return fmt.Errorf("商品属性 ID 不能为空")
	}
	productAttribute, err := srv.AttributeRepo.GetProductAttribute(uint(req.Id))
	if err != nil {
		return err
	}
	resp.ProductAttribute = productAttribute.ToProtobuf()
	return nil
}

func (srv *AttributeService) GetProductAttributes(ctx context.Context, req *pb.Product, resp *pb.ProductAttributeResponse) error {
	productAttributes, err := srv.AttributeRepo.GetProductAttributes()
	if err != nil {
		return err
	}
	resp.ProductAttributes = model.ProductAttributeToProtobufArray(productAttributes)
	return nil
}
