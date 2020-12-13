package model

import (
	"time"

	pb "github.com/leor-w/laracom/user-service/proto/user"
)

type PasswordReset struct {
	Email    string `gorm:"index"`
	Token    string `gorm:"not null"`
	CreateAt time.Time
}

func (model *PasswordReset) ToORM(req *pb.PasswordReset) error {
	if req.Email != "" {
		model.Email = req.Email
	}
	if req.Token != "" {
		model.Token = req.Token
	}
	return nil
}

func (model *PasswordReset) ToProtoBuf() (*pb.PasswordReset, error) {
	passwordReset := &pb.PasswordReset{}
	passwordReset.Email = model.Email
	passwordReset.Token = model.Token
	passwordReset.CreateAt = model.CreateAt.Format("2006-01-02 15:04:05")
	return passwordReset, nil
}
