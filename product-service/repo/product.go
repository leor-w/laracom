package repo

import (
	"github.com/leor-w/laracom/product-service/model"
	"gorm.io/gorm"
)

type ProductRepositoryInterface interface {
	Create(product *model.Product) error
	Update(product *model.Product) error
	Delete(product *model.Product) error
	GetById(id uint) (*model.Product, error)
	GetBySlug(slug string) (*model.Product, error)
	GetAll() ([]*model.Product, error)
}

type ProductRepository struct {
	Db *gorm.DB
}

func (repo *ProductRepository) Create(product *model.Product) error {
	if err := repo.Db.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (repo *ProductRepository) Update(product *model.Product) error {
	if err := repo.Db.Save(product).Error; err != nil {
		return err
	}
	return nil
}

func (repo *ProductRepository) Delete(product *model.Product) error {
	if err := repo.Db.Delete(product).Error; err != nil {
		return err
	}
	return nil
}

func (repo *ProductRepository) GetById(id uint) (*model.Product, error) {
	product := &model.Product{}
	if err := repo.Db.Where("id = ?", id).First(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (repo *ProductRepository) GetBySlug(slug string) (*model.Product, error) {
	product := &model.Product{}
	if err := repo.Db.Where("slug = ?", slug).First(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (repo *ProductRepository) GetAll() ([]*model.Product, error) {
	products := []*model.Product{}
	if err := repo.Db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
