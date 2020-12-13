package model

import (
	"time"

	"gorm.io/gorm"
)

// User 创建前执行的钩子
//func (model *User) BeforeCreate(db *gorm.DB) error {
//	uuid := gouuid.NewV4()
//	model.Id = uuid.String()
//	return nil
//}

// User 保存前执行的钩子
func (model *User) BeforeSave(db *gorm.DB) error {
	model.UpdatedAt = time.Now()
	return nil
}

// PasswordReset 创建前执行的钩子
func (model *PasswordReset) BeforeCreate(db *gorm.DB) error {
	model.CreateAt = time.Now()
	return nil
}
