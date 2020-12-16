package repo

import (
	"github.com/leor-w/laracom/product-service/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ImageRepositoryInterface interface {
	Create(image *model.ProductImage) error
	Update(image *model.ProductImage) error
	Delete(image *model.ProductImage) error
	GetById(id uint) (*model.ProductImage, error)
	GetByProductId(productId uint) ([]*model.ProductImage, error)
}

type ImageRepository struct {
	Db *gorm.DB
}

func (repo *ImageRepository) Create(image *model.ProductImage) error {
	if err := repo.Db.Create(image).Error; err != nil {
		logrus.Errorf("Create product image failed : image = [%v] err = %v", *image, err)
		return err
	}
	return nil
}

func (repo *ImageRepository) Update(image *model.ProductImage) error {
	if err := repo.Db.Save(image).Error; err != nil {
		logrus.Errorf("Update product image failed : image = [%v] err = %v", *image, err)
		return err
	}
	return nil
}

func (repo *ImageRepository) Delete(image *model.ProductImage) error {
	if err := repo.Db.Delete(image).Error; err != nil {
		logrus.Errorf("Delete product image failed : image = [%v] err = %v", *image, err)
		return err
	}
	return nil
}

func (repo *ImageRepository) GetById(id uint) (*model.ProductImage, error) {
	image := &model.ProductImage{}
	if err := repo.Db.First(image, id).Error; err != nil {
		logrus.Errorf("Not find image record : id = %d", id)
		return nil, err
	}
	return image, nil
}

func (repo *ImageRepository) GetByProductId(productId uint) ([]*model.ProductImage, error) {
	images := []*model.ProductImage{}
	if err := repo.Db.Where("product_id = ?", productId).Find(&images).Error; err != nil {
		logrus.Errorf("Not find image with product : productId = [%d], err = %v", productId, err)
		return nil, err
	}
	return images, nil
}
