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
	AppUserId int    `sql:"not null;unique"`
	Password  []byte `sql:"not null"`
}

func GetAllAppUsers() []AppUser {
	var res []AppUser
	if err := _db.Find(&res).Error; err != nil {
		panic(err)
	}
	return res
}

const (
	adminUserId = "ah@jabara.info"
)

func getAppUserByUserId(db *gorm.DB, userId string) *AppUser {
	var res []AppUser
	if err := db.Where("user_id = ?", userId).First(&res).Error; err != nil {
		panic(err)
	}
	if len(res) > 1 {
		panic("結果が２件以上あるのは異常事態. AppUser.UserId -> " + userId)
	}
	if len(res) == 1 {
		return &res[0]
	}
	return nil
}

func createAdminUserIfNecessary() {
	tx := _db.Begin().LogMode(true)
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()

	if adminUser := getAppUserByUserId(tx, adminUserId); adminUser != nil {
		return
	}
	adminUser := AppUser{UserId: adminUserId}
	credential := AppUserCredential{AppUser: adminUser, Password: []byte("hogehoge")}
	if err := tx.Create(&credential).Error; err != nil {
		panic(err)
	}
}
