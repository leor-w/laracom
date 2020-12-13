package repo

import (
	"github.com/leor-w/laracom/user-service/model"
	"gorm.io/gorm"
)

type PasswordRepositoryInterface interface {
	Create(reset *model.PasswordReset) error
	GetByToken(token string) (*model.PasswordReset, error)
	DeletePasswordReset(reset *model.PasswordReset) error
	GetByEmail(email string) (*model.PasswordReset, error)
}

type PasswordRepository struct {
	Db *gorm.DB
}

func (repo *PasswordRepository) Create(reset *model.PasswordReset) error {
	err := repo.Db.Create(reset).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *PasswordRepository) GetByToken(token string) (*model.PasswordReset, error) {
	var (
		reset = &model.PasswordReset{}
		err   error
	)
	err = repo.Db.Where("token = ?", token).First(&reset).Error
	if err != nil {
		return nil, err
	}
	return reset, nil
}

func (repo *PasswordRepository) DeletePasswordReset(reset *model.PasswordReset) error {
	return repo.Db.Delete(reset).Error
}

func (repo *PasswordRepository) GetByEmail(email string) (*model.PasswordReset, error) {
	reset := &model.PasswordReset{}
	err := repo.Db.Where("email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return reset, nil
}
