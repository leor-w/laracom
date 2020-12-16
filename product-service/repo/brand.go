package repo

import (
	"github.com/leor-w/laracom/product-service/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BrandRepositoryInterface interface {
	Create(brand *model.Brand) error
	Update(brand *model.Brand) error
	Delete(brand *model.Brand) error
	GetById(id uint) (*model.Brand, error)
	GetAll() ([]*model.Brand, error)
	GetWithProduct(brandId uint) (*model.Brand, error)
}

type BrandRepository struct {
	Db *gorm.DB
}

func (repo *BrandRepository) Create(brand *model.Brand) error {
	if err := repo.Db.Create(brand).Error; err != nil {
		logrus.Errorf("create brand failed : brand = [%v] err = %v", *brand, err)
		return err
	}
	return nil
}

func (repo *BrandRepository) Update(brand *model.Brand) error {
	if err := repo.Db.Save(brand).Error; err != nil {
		logrus.Errorf("update brand failed : brand = [%v] err = %v", *brand, err)
		return err
	}
	return nil
}

func (repo *BrandRepository) Delete(brand *model.Brand) error {
	if err := repo.Db.Delete(brand).Error; err != nil {
		logrus.Errorf("delete brand failed : brand = [%v] err = %v", *brand, err)
		return err
	}
	return nil
}

func (repo *BrandRepository) GetById(id uint) (*model.Brand, error) {
	brand := &model.Brand{}
	if err := repo.Db.First(brand, id).Error; err != nil {
		logrus.Errorf("not find brand record : brand id = %d, err = %v", id, err)
		return nil, err
	}
	return brand, nil
}

func (repo *BrandRepository) GetAll() ([]*model.Brand, error) {
	brands := []*model.Brand{}
	if err := repo.Db.Find(&brands).Error; err != nil {
		logrus.Errorf("find all brand record failed : err = %v", err)
		return nil, err
	}
	return brands, nil
}

func (repo *BrandRepository) GetWithProduct(brandId uint) (*model.Brand, error) {
	brand := &model.Brand{}
	if err := repo.Db.Where("id = ?", brandId).Preload("Products").First(brand).Error; err != nil {
		logrus.Errorf("find brand with product failed : brandId = %d", brandId)
		return nil, err
	}
	return brand, nil
}
