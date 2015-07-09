package model

import (
	"bytes"
	"crypto/sha256"
	"github.com/jinzhu/gorm"
)

type AppUser struct {
	UserId string `sql:"type varchar(200);unique;"`
	gorm.Model
}

type AppUserCredential struct {
	AppUser   AppUser
	AppUserId uint   `sql:"not null;unique"`
	Password  []byte `sql:"not null"`
}

func GetAllAppUsers() []AppUser {
	var res []AppUser
	if err := _db.Find(&res).Error; err != nil {
		panic(err)
	}
	if res == nil {
		return make([]AppUser, 0)
	}
	return res
}

func Authenticate(userId string, password string) bool {
	credential, nf := getAppUserCredentialByUserId(_db, userId)
	if nf != nil {
		return false
	}
	c := bytes.Compare(toHash(password), credential.Password)
	return c == 0
}

func AddAppUser(userId, password string) (*AppUser, InvalidValue) {
	if len(userId) == 0 || len(password) == 0 {
		return nil, NewInvalidValue("入力されていない項目があります.")
	}

	tx := _db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()

	// TODO ユーザIDの重複検査

	newUser := AppUser{
		UserId: userId,
	}
	newCredential := AppUserCredential{
		AppUser:  newUser,
		Password: toHash(password),
	}
	mustInsert(tx, &newCredential)

	return &newUser, nil
}

const (
	adminUserId       = "ah@jabara.info"
	adminUserPassword = "hogehoge"
)

func getAppUserByUserId(db *gorm.DB, userId string) (AppUser, NotFound) {
	var res []AppUser
	if err := db.Where("user_id = ?", userId).First(&res).Error; err != nil {
		panic(err)
	}
	if len(res) > 1 {
		panic("結果が２件以上あるのは異常事態. AppUser.UserId -> " + userId)
	}
	if len(res) == 1 {
		return res[0], nil
	}
	return AppUser{}, NewNotFound()
}

func getAppUserCredentialByUserId(db *gorm.DB, userId string) (AppUserCredential, NotFound) {
	var res []AppUserCredential
	if err := db.Table("app_user_credentials").
		Select("app_user_credentials.password").
		Joins("inner join app_users on app_user_credentials.app_user_id = app_users.id").
		Where("app_users.user_id = ?", userId).
		First(&res).Error; err != nil {
		panic(err)
	}
	if len(res) == 0 {
		return AppUserCredential{}, NewNotFound()
	}
	return res[0], nil
}

func createAdminUserIfNecessary() {
	if _, nf := getAppUserByUserId(_db, adminUserId); nf == nil {
		return
	}
	AddAppUser(adminUserId, adminUserPassword)
}

func toHash(password string) []byte {
	p := sha256.Sum256([]byte(password))
	return p[:]
}
