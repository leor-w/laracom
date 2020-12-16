package repo

import (
	"github.com/leor-w/laracom/product-service/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AttributeRepositoryInterface interface {
	CreateAttribute(attribute *model.Attribute) error
	UpdateAttribute(attribute *model.Attribute) error
	DeleteAttribute(attribute *model.Attribute) error

	CreateValue(value *model.AttributeValue) error
	UpdateValue(value *model.AttributeValue) error
	DeleteValue(value *model.AttributeValue) error

	CreateProductAttribute(attribute *model.ProductAttribute) error
	UpdateProductAttribute(attribute *model.ProductAttribute) error
	DeleteProductAttribute(attribute *model.ProductAttribute) error

	GetAttribute(id uint) (*model.Attribute, error)
	GetAttributes() ([]*model.Attribute, error)
	GetValue(id uint) (*model.AttributeValue, error)
	GetValues() ([]*model.AttributeValue, error)
	GetProductAttribute(id uint) (*model.ProductAttribute, error)
	GetProductAttributes() ([]*model.ProductAttribute, error)
}

type AttributeRepository struct {
	Db *gorm.DB
}

func (repo *AttributeRepository) CreateAttribute(attribute *model.Attribute) error {
	if err := repo.Db.Create(attribute).Error; err != nil {
		logrus.Errorf("create attribute failed : err = %v", err)
		return err
	}
	return nil
}

func (repo *AttributeRepository) UpdateAttribute(attribute *model.Attribute) error {
	if err := repo.Db.Save(attribute).Error; err != nil {
		logrus.Errorf("update attribute failed : err = %v", err)
		return err
	}
	return nil
}

func (repo *AttributeRepository) DeleteAttribute(attribute *model.Attribute) error {
	if err := repo.Db.Delete(attribute).Error; err != nil {
		logrus.Errorf("delete attribute failed : err = %v", err)
		return err
	}
	return nil
}

func (repo *AttributeRepository) CreateValue(value *model.AttributeValue) error {
	if err := repo.Db.Create(value).Error; err != nil {
		logrus.Errorf("create attribute value failed : err = %v", err)
		return err
	}
	return nil
}

func (repo *AttributeRepository) UpdateValue(value *model.AttributeValue) error {
	if err := repo.Db.Save(value).Error; err != nil {
		logrus.Errorf("update attribute value failed : err = %v", err)
		return err
	}
	return nil
}

func (repo *AttributeRepository) DeleteValue(value *model.AttributeValue) error {
	if err := repo.Db.Delete(value).Error; err != nil {
		logrus.Errorf("delete attribute value failed : err = %v", err)
		return err
	}
	return nil
}

func (repo *AttributeRepository) CreateProductAttribute(attribute *model.ProductAttribute) error {
	if err := repo.Db.Create(attribute).Error; err != nil {
		logrus.Errorf("create product attribute failed : err = %v", err)
		return err
	}
	return nil
}

func (repo *AttributeRepository) UpdateProductAttribute(attribute *model.ProductAttribute) error {
	if err := repo.Db.Save(attribute).Error; err != nil {
		logrus.Errorf("update product attribute failed : err = %v", err)
		return err
	}
	return nil
}

func (repo *AttributeRepository) DeleteProductAttribute(attribute *model.ProductAttribute) error {
	if err := repo.Db.Delete(attribute).Error; err != nil {
		logrus.Errorf("delete product attribute failed : err = %v", err)
		return err
	}
	return nil
}

func (repo *AttributeRepository) GetAttribute(id uint) (*model.Attribute, error) {
	model := &model.Attribute{}
	if err := repo.Db.First(model, id).Error; err != nil {
		logrus.Errorf("not find attribute record : err = %v", err)
		return nil, err
	}
	return model, nil
}

func (repo *AttributeRepository) GetAttributes() ([]*model.Attribute, error) {
	attrbutes := []*model.Attribute{}
	if err := repo.Db.Find(&attrbutes).Error; err != nil {
		return nil, err
	}
	return attrbutes, nil
}

func (repo *AttributeRepository) GetValue(id uint) (*model.AttributeValue, error) {
	model := &model.AttributeValue{}
	if err := repo.Db.First(model, id).Error; err != nil {
		logrus.Errorf("not find attribute value record : err = %v", err)
		return nil, err
	}
	return model, nil
}

func (repo *AttributeRepository) GetValues() ([]*model.AttributeValue, error) {
	values := []*model.AttributeValue{}
	if err := repo.Db.Find(&values).Error; err != nil {
		logrus.Errorf("not find attribute value record : err = %v", err)
		return nil, err
	}
	return values, nil
}

func (repo *AttributeRepository) GetProductAttribute(id uint) (*model.ProductAttribute, error) {
	model := &model.ProductAttribute{}
	if err := repo.Db.First(model, id).Error; err != nil {
		logrus.Errorf("not find product attribute record : err = %v", err)
		return nil, err
	}
	return model, nil
}

func (repo *AttributeRepository) GetProductAttributes() ([]*model.ProductAttribute, error) {
	models := []*model.ProductAttribute{}
	if err := repo.Db.Find(&models).Error; err != nil {
		logrus.Errorf("not find product attribute record : err %v", err)
		return nil, err
	}
	return models, nil
}
