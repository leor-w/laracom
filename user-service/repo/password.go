package repo

import (
	pb "github.com/leor-w/laracom/user-service/proto/user"
	"gorm.io/gorm"
)

type PasswordRepositoryInterface interface {
	Create(reset *pb.PasswordReset) error
	GetByToken(token string) (*pb.PasswordReset, error)
}

type PasswordRepository struct {
	Db *gorm.DB
}

func (repo *PasswordRepository) Create(reset *pb.PasswordReset) error {
	err := repo.Db.Create(reset).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *PasswordRepository) GetByToken(token string) (*pb.PasswordReset, error) {
	var (
		reset = &pb.PasswordReset{}
		err   error
	)
	err = repo.Db.Where("token = ?", token).First(&reset).Error
	if err != nil {
		return nil, err
	}
	return reset, nil
}