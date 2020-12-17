package model

import pb "github.com/leor-w/laracom/product-service/proto/product"

type ProductImage struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;<-:create"`
	ProductId uint   `gorm:"unsigned, default:0;index"`
	Src       string `gorm:"type:varchar(255)"`
}

func (model *ProductImage) ToORM(req *pb.ProductImage) {
	if req.Id != 0 {
		model.ID = uint(req.Id)
	}
	if req.ProductId != 0 {
		model.ProductId = uint(req.ProductId)
	}
	if req.Src != "" {
		model.Src = req.Src
	}
}

func (model *ProductImage) ToProtobuf() *pb.ProductImage {
	return &pb.ProductImage{
		Id:        uint32(model.ID),
		ProductId: uint32(model.ProductId),
		Src:       model.Src,
	}
}

func ProductImageToProtobufArray(models []*ProductImage) []*pb.ProductImage {
	productImages := []*pb.ProductImage{}
	for _, model := range models {
		productImages = append(productImages, model.ToProtobuf())
	}
	return productImages
}
