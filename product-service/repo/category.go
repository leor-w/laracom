package repo

import (
	"github.com/leor-w/laracom/product-service/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CategoryRepositoryInterface interface {
	Create(category *model.Category) error
	Update(category *model.Category) error
	Delete(category *model.Category) error
	GetById(id uint) (*model.Category, error)
	GetAll() ([]*model.Category, error)
	GetWithProducts(categoryId uint) (*model.Category, error)
}

type CategoryRepository struct {
	Db *gorm.DB
}

func (repo *CategoryRepository) Create(category *model.Category) error {
	if err := repo.Db.Create(category).Error; err != nil {
		logrus.Errorf("create category failed : err = %v", err)
		return err
	}
	return nil
}

func (repo *CategoryRepository) Update(category *model.Category) error {
	if err := repo.Db.Save(category).Error; err != nil {
		logrus.Errorf("update category failed : err = %v", err)
		return err
	}
	return nil
}

func (repo *CategoryRepository) Delete(category *model.Category) error {
	if err := repo.Db.Delete(category).Error; err != nil {
		logrus.Errorf("delete category failed : err = %v", err)
		return err
	}
	return nil
}

func (repo *CategoryRepository) GetById(id uint) (*model.Category, error) {
	category := &model.Category{}
	if err := repo.Db.First(category, id).Error; err != nil {
		logrus.Errorf("not find category record : err = %v", err)
		return nil, err
	}
	return category, nil
}

func (repo *CategoryRepository) GetAll() ([]*model.Category, error) {
	categories := []*model.Category{}
	if err := repo.Db.Find(&categories).Error; err != nil {
		logrus.Errorf("get all category failed : err = %v", err)
		return nil, err
	}
	return categories, nil
}

func (repo *CategoryRepository) GetWithProducts(categoryId uint) (*model.Category, error) {
	category := &model.Category{}
	if err := repo.Db.Where("id = ?", categoryId).Preload("Products").First(category).Error; err != nil {
		logrus.Errorf("get category with product failed : err = %v", err)
		return nil, err
	}
	return category, nil
}
