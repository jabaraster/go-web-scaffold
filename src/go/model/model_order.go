package model

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	Product   Product
	ProductId int `sql:"not null"`
	gorm.Model
}

func GetAllOrders() []Order {
	var orders []Order
	err := _db.Preload("Product").Find(&orders).Error
	if err != nil {
		panic(err)
	}
	return orders
}
