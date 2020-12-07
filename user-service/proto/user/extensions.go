package laracom_service_user

import (
	gouuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func (model *User) BeforeCreate(db *gorm.DB) error {
	uuid := gouuid.NewV4()
	model.Id = uuid.String()
	return nil
}
