package model

import (
	"github.com/jinzhu/gorm"
)

type AppUser struct {
	UserId string `sql:"type varchar(200);unique;"`
	gorm.Model
}

type AppUserCredential struct {
	AppUser   AppUser
	AppUserId int    `sql:"not null"`
	Password  []byte `sql:"not null"`
}
