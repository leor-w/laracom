package model

import (
	"time"

	pb "github.com/leor-w/laracom/product-service/proto/product"
)

type AttributeValue struct {
	ID                uint   `gorm:"primary_key;autoIncrement;<-:create"`
	Value             string `gorm:"type:varchar(255)"`
	AttributeId       uint   `gorm:"undefined,default:0;index"`
	Attribute         Attribute
	ProductAttributes []*ProductAttribute `gorm:"many2many:attribute_value_product_attribute"`
	CreatedAt         time.Time           `gorm:"<-:create"`
	UpdatedAt         time.Time
}

func (model *AttributeValue) ToORM(req *pb.AttributeValue) {
	if req.Id != 0 {
		model.ID = uint(req.Id)
	}
	if req.Value != "" {
		model.Value = req.Value
	}
	if req.AttributeId != 0 {
		model.AttributeId = uint(req.AttributeId)
	}
}

func (model *AttributeValue) ToProtobuf() *pb.AttributeValue {
	attributeValue := &pb.AttributeValue{
		Id:          uint32(model.ID),
		Value:       model.Value,
		AttributeId: uint32(model.AttributeId),
		CreatedAt:   model.CreatedAt.Format(Format),
		UpdatedAt:   model.UpdatedAt.Format(Format),
	}
	if model.ProductAttributes != nil {
		attributeValue.ProductAttributes = ProductAttributeToProtobufArray(model.ProductAttributes)
	}
	return attributeValue
}

func AttributeValueToProtobufArray(values []*AttributeValue) []*pb.AttributeValue {
	attributeValues := []*pb.AttributeValue{}
	for _, value := range values {
		attributeValues = append(attributeValues, value.ToProtobuf())
	}
	return attributeValues
}
