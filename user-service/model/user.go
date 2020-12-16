package model

import (
	"strconv"
	"time"

	pb "github.com/leor-w/laracom/user-service/proto/user"
)

type User struct {
	Id            uint   `gorm:"primary_key;autoIncrement;<-:create"`
	Name          string `gorm:"type:varchar(100)"`
	Email         string `gorm:"type:varchar(100)"`
	Password      string
	Status        uint8 `gorm:"default:1"`
	StripeId      string
	CardBrand     string
	CardLastFour  string
	RememberToken string
	CreatedAt     time.Time `gorm:"<-:create"`
	UpdatedAt     time.Time
}

// 将 proto user 转换为 orm 的 user model
func (model *User) ToORM(req *pb.User) error {
	if req.Id != "" {
		id, err := strconv.ParseUint(req.Id, 10, 64)
		if err != nil {
			return err
		}
		model.Id = uint(id)
	}
	if req.Name != "" {
		model.Name = req.Name
	}
	if req.Email != "" {
		model.Email = req.Email
	}
	if req.Password != "" {
		model.Password = req.Password
	}
	if req.Status != "" {
		status, err := strconv.ParseUint(req.Status, 10, 64)
		if err != nil {
			return err
		}
		model.Status = uint8(status)
	}
	if req.StripeId != "" {
		model.StripeId = req.StripeId
	}
	if req.CardBrand != "" {
		model.CardBrand = req.CardBrand
	}
	if req.CardLastFour != "" {
		model.CardLastFour = req.CardLastFour
	}
	if req.RememberToken != "" {
		model.RememberToken = req.RememberToken
	}
	return nil
}

// 将 orm 转换为 proto user
func (model *User) ToProtoBuf() (*pb.User, error) {
	user := &pb.User{}
	user.Id = strconv.FormatUint(uint64(model.Id), 10)
	user.Name = model.Name
	user.Email = model.Email
	user.Status = strconv.FormatUint(uint64(model.Status), 10)
	user.StripeId = model.StripeId
	user.CardBrand = model.CardBrand
	user.CardLastFour = model.CardLastFour
	user.CreatedAt = model.CreatedAt.Format("2006-01-02 15:04:05")
	user.UpdatedAt = model.UpdatedAt.Format("2006-01-02 15:04:05")
	return user, nil
}
