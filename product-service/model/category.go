package model

import (
	"time"

	pb "github.com/leor-w/laracom/product-service/proto/product"
)

type Category struct {
	ID          uint      `gorm:"primary_key;autoIncrement;<-:create"`
	Name        string    `gorm:"type:varchar(255);unique_index"`
	Slug        string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:text"`
	Cover       string    `gorm:"type:varchar(255)"`
	Status      uint8     `gorm:"unsigned,default:0"`
	ParentId    uint      `gorm:"unsigned,default:0"`
	Lft         uint32    `gorm:"undefined,default:0;index"`
	Rgt         uint32    `gorm:"undefined,default:0;index"`
	CreatedAt   time.Time `gorm:"<-:create"`
	UpdatedAt   time.Time
	Products    []*Product `gorm:"many2many:category_product;"`
}

func (model *Category) ToORM(req *pb.Category) {
	if req.Id != 0 {
		model.ID = uint(req.Id)
	}
	if req.Name != "" {
		model.Name = req.Name
	}
	if req.Slug != "" {
		model.Slug = req.Slug
	}
	if req.Description != "" {
		model.Description = req.Description
	}
	if req.Cover != "" {
		model.Cover = req.Cover
	}
	if req.Status != 0 {
		model.Status = uint8(req.Status)
	}
	model.ParentId = uint(req.ParentId)
	model.Lft = req.Lft
	model.Rgt = req.Rgt
}

func (model *Category) ToProtobuf() *pb.Category {
	category := &pb.Category{
		Id:          uint32(model.ID),
		Name:        model.Name,
		Slug:        model.Slug,
		Description: model.Description,
		Cover:       model.Cover,
		Status:      uint32(model.Status),
		ParentId:    uint32(model.ParentId),
		Lft:         model.Lft,
		Rgt:         model.Rgt,
		CreatedAt:   model.CreatedAt.Format(Format),
		UpdatedAt:   model.UpdatedAt.Format(Format),
	}
	if model.Products != nil {
		category.Products = ProductProtobufToArray(model.Products)
	}
	return category
}

func CategoryToProtobufArray(models []*Category) []*pb.Category {
	categorys := []*pb.Category{}
	for _, model := range models {
		categorys = append(categorys, model.ToProtobuf())
	}
	return categorys
}
