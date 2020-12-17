package model

import (
	"time"

	pb "github.com/leor-w/laracom/product-service/proto/product"
)

type ProductAttribute struct {
	ID              uint      `gorm:"primaryKey;autoIncrement;<-:create"`
	ProductId       uint      `gorm:"undefined,default:0;index"`
	Quantity        uint32    `gorm:"unsigned,default:0"`
	Price           float32   `gorm:"type:decimal(8,2)"`
	SalePrice       float32   `gorm:"type:decimal(8,2)"`
	Default         uint8     `gorm:"unsigned,default:0"`
	CreatedAt       time.Time `gorm:"<-:create"`
	UpdatedAt       time.Time
	Product         Product
	AttributeValues []*AttributeValue `gorm:"many2many:attribute_value_product_attribute"`
}

func (model *ProductAttribute) ToORM(req *pb.ProductAttribute) {
	if req.Id != 0 {
		model.ID = uint(req.Id)
	}
	if req.ProductId != 0 {
		model.ProductId = uint(req.ProductId)
	}
	if req.Quantiry != 0 {
		model.Quantity = req.Quantiry
	}
	if req.Price != 0 {
		model.Price = req.Price
	}
	if req.SalePrice != 0 {
		model.SalePrice = req.SalePrice
	}
	if req.Default != 0 {
		model.Default = uint8(req.Default)
	}
}

func (model *ProductAttribute) ToProtobuf() *pb.ProductAttribute {
	attribute := &pb.ProductAttribute{
		Id:        uint32(model.ID),
		ProductId: uint32(model.ProductId),
		Quantiry:  model.Quantity,
		Price:     model.Price,
		SalePrice: model.SalePrice,
		Default:   uint32(model.Default),
		CreatedAt: model.CreatedAt.Format(Format),
		UpdatedAt: model.UpdatedAt.Format(Format),
	}
	if model.AttributeValues != nil {
		attribute.AttributeValues = AttributeValueToProtobufArray(model.AttributeValues)
	}
	return attribute
}

func ProductAttributeToProtobufArray(productAttrs []*ProductAttribute) []*pb.ProductAttribute {
	attributes := []*pb.ProductAttribute{}
	for _, productAttr := range productAttrs {
		attributes = append(attributes, productAttr.ToProtobuf())
	}
	return attributes
}
