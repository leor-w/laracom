package model

import (
	"time"

	pb "github.com/leor-w/laracom/product-service/proto/product"
)

type Attribute struct {
	ID        uint              `gorm:"primaryKey;autoIncrement;<-:create"`
	Name      string            `gorm:"type:varchar(255);unique_index"`
	Values    []*AttributeValue `gorm:"foreignKey:ID"`
	CreatedAt time.Time         `gorm:"<-:create"`
	UpdatedAt time.Time
}

func (model *Attribute) ToORM(req *pb.Attribute) {
	if req.Id != 0 {
		model.ID = uint(req.Id)
	}
	if req.Name != "" {
		model.Name = req.Name
	}
}

func (model *Attribute) ToProtobuf() *pb.Attribute {
	attribute := &pb.Attribute{
		Id:        uint32(model.ID),
		Name:      model.Name,
		CreatedAt: model.CreatedAt.Format(Format),
		UpdatedAt: model.UpdatedAt.Format(Format),
	}
	if model.Values != nil {
		attribute.Values = AttributeValueToProtobufArray(model.Values)
	}
	return attribute
}

func AttributeToProtobufArray(attrs []*Attribute) []*pb.Attribute {
	attributes := []*pb.Attribute{}
	for _, attr := range attrs {
		attributes = append(attributes, attr.ToProtobuf())
	}
	return attributes
}
