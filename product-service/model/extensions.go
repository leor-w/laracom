package model

import (
	"time"

	"gorm.io/gorm"
)

const Format = "2006-01-02 15:04:05"

func (model *Product) BeforeCreate(db *gorm.DB) error {
	model.CreatedAt = time.Now()
	return nil
}

func (model *Product) BeforeSave(db *gorm.DB) error {
	model.UpdatedAt = time.Now()
	return nil
}

func (model *ProductAttribute) BeforeCreate(db *gorm.DB) error {
	model.CreatedAt = time.Now()
	return nil
}

func (model *ProductAttribute) BeforeSave(db *gorm.DB) error {
	model.UpdatedAt = time.Now()
	return nil
}

func (model *AttributeValue) BeforeCreate(db *gorm.DB) error {
	model.CreatedAt = time.Now()
	return nil
}

func (model *AttributeValue) BeforeSave(db *gorm.DB) error {
	model.UpdatedAt = time.Now()
	return nil
}

func (model *Attribute) BeforeCreate(db *gorm.DB) error {
	model.CreatedAt = time.Now()
	return nil
}

func (model *Attribute) BeforeSave(db *gorm.DB) error {
	model.UpdatedAt = time.Now()
	return nil
}

func (model *Category) BeforeCreate(db *gorm.DB) error {
	model.CreatedAt = time.Now()
	return nil
}

func (model *Category) BeforeSave(db *gorm.DB) error {
	model.UpdatedAt = time.Now()
	return nil
}
