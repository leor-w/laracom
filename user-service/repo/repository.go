package repo

import (
	"github.com/leor-w/laracom/user-service/model"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Create(user *model.User) error
	Get(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetAll() ([]*model.User, error)
	Update(user *model.User) error
}

type UserRepository struct {
	Db *gorm.DB
}

func (repo *UserRepository) Create(user *model.User) error {
	if err := repo.Db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) Get(id uint) (*model.User, error) {
	user := &model.User{}
	user.ID = id
	err := repo.Db.First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetByEmail(email string) (*model.User, error) {
	user := &model.User{}
	err := repo.Db.Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetAll() ([]*model.User, error) {
	var users []*model.User
	err := repo.Db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Update(user *model.User) error {
	if err := repo.Db.Save(user).Error; err != nil {
		return err
	}
	return nil
}
