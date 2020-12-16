package model

import (
	"time"

	pb "github.com/leor-w/laracom/product-service/proto/product"
)

type Brand struct {
	ID        uint   `gorm:"primary_key;autoIncrement;<-:create"`
	Name      string `gorm:"type:varchar(255)"`
	Products  []*Product
	CreatedAt *time.Time `gorm:"<-:create"`
	UpdatedAt *time.Time
}

func (model *Brand) ToORM(req *pb.Brand) {
	if req.Id != 0 {
		model.ID = uint(req.Id)
	}
	if req.Name != "" {
		model.Name = req.Name
	}
}

func (model *Brand) ToProtobuf() *pb.Brand {
	brandProto := &pb.Brand{
		Id:        uint32(model.ID),
		Name:      model.Name,
		CreatedAt: model.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: model.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	if model.Products != nil {
		products := []*pb.Product{}
		for _, product := range model.Products {
			products = append(products, product.ToProtobuf())
		}
		brandProto.Products = products
	}
	return brandProto
}

func BrandToProtobufArray(models []*Brand) []*pb.Brand {
	brands := []*pb.Brand{}
	for _, model := range models {
		brands = append(brands, model.ToProtobuf())
	}
	return brands
}
