package user

import (
	"github.com/jinzhu/gorm"
	gouuid "github.com/satori/go.uuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := gouuid.NewV4()
	return scope.SetColumn("Id", uuid.String())
}
