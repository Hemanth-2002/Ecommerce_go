package database

import (
	model "ecommerce/db_model"

	"github.com/jinzhu/gorm"
)

type Dbclient struct {
	Db *gorm.DB
}

type DbProducts interface {
	GetProds() ([]model.Product, error)
}

func (s Dbclient) GetProds() ([]model.Product, error) {
	var products []model.Product
	s.Db.Find(&products)
	return products, nil
}
