package model

import (
	"time"

	"gorm.io/gorm"
)

func (model *Product) BeforeSave(db *gorm.DB) error {
	model.UpdatedAt = time.Now()
	return nil
}
