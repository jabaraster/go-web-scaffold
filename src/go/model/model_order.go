package model

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	Product   Product
	ProductId int `sql:"not null"`
	gorm.Model
}
