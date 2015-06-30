package model

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	Code string `sql:"type:varchar(100);unique;not null"`
	gorm.Model
}
