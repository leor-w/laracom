package model

import (
	"time"

	pb "github.com/leor-w/laracom/product-service/proto/product"
)

type Product struct {
	Id           uint      `gorm:"primarykey;autoIncrement;<-:create"`
	BrandId      uint32    `gorm:"type:varchar(255)"`
	Sku          string    `gorm:"type:varchar(255)"`
	Name         string    `gorm:"type:varchar(255)"`
	Slug         string    `gorm:"type:varchar(255)"`
	Description  string    `gorm:"type:text"`
	Cover        string    `gorm:"type:varchar(255)"`
	Quantity     uint32    `gorm:"unsigned, default:0"`
	Price        float32   `grom:"type:decimal(8,2)"`
	SalePrice    float32   `gorm:"type:decimal(8,2)"`
	Status       uint8     `gorm:"unsigned,default:0"`
	Length       float32   `gorm:"type:decimal(8,2)"`
	Width        float32   `gorm:"type:decimal(8,2)"`
	Height       float32   `gorm:"type:decimal(8,2)"`
	Weight       float32   `gorm:"type:decimal(8,2)"`
	DistanceUnit string    `gorm:"type:varchar(255)"`
	MassUnit     string    `gorm:"type:varchar(255)"`
	CreatedAt    time.Time `gorm:"<-:create"`
	UpdatedAt    time.Time
	Brand        Brand
	Attributes   []*ProductAttribute
	Images       []*ProductImage
	Categories   []*Category `gorm:"many2many:category_product;"`
}

func (model *Product) ToORM(req *pb.Product) {
	if req.Id != 0 {
		model.Id = uint(req.Id)
	}
	if req.BrandId != 0 {
		model.BrandId = req.BrandId
	}
	if req.Sku != "" {
		model.Sku = req.Sku
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
	if req.Quantity != 0 {
		model.Quantity = req.Quantity
	}
	if req.Price != 0 {
		model.Price = req.Price
	}
	if req.SalePrice != 0 {
		model.SalePrice = req.SalePrice
	}
	if req.Status != 0 {
		model.Status = uint8(req.Status)
	}
	if req.Length != 0 {
		model.Length = req.Length
	}
	if req.Width != 0 {
		model.Width = req.Width
	}
	if req.Height != 0 {
		model.Height = req.Height
	}
	if req.Weight != 0 {
		model.Weight = req.Weight
	}
	if req.DistanceUint != "" {
		model.DistanceUnit = req.DistanceUint
	}
	if req.MassUint != "" {
		model.MassUnit = req.MassUint
	}
}

func (model *Product) ToProtobuf() *pb.Product {
	product := &pb.Product{
		Id:           uint32(model.Id),
		BrandId:      model.BrandId,
		Sku:          model.Sku,
		Name:         model.Name,
		Slug:         model.Slug,
		Description:  model.Description,
		Cover:        model.Cover,
		Quantity:     model.Quantity,
		Price:        model.Price,
		SalePrice:    model.SalePrice,
		Status:       uint32(model.Status),
		Length:       model.Length,
		Width:        model.Width,
		Height:       model.Height,
		Weight:       model.Weight,
		DistanceUint: model.DistanceUnit,
		MassUint:     model.MassUnit,
		CreatedAt:    model.CreatedAt.Format(Format),
		UpdatedAt:    model.UpdatedAt.Format(Format),
	}
	if model.Images != nil {
		product.Images = ProductImageToProtobufArray(model.Images)
	}
	if model.Brand.ID != 0 {
		product.Brand = model.Brand.ToProtobuf()
	}
	if model.Categories != nil {
		product.Categories = CategoryToProtobufArray(model.Categories)
	}
	if model.Attributes != nil {
		product.Attributes = ProductAttributeToProtobufArray(model.Attributes)
	}
	return product
}

func ProductProtobufToArray(models []*Product) []*pb.Product {
	products := []*pb.Product{}
	for _, model := range models {
		products = append(products, model.ToProtobuf())
	}
	return products
}
