package model

import (
    "bytes"
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
    b := []byte(password)
    c := bytes.Compare(b, credential.Password)
    return c == 0
}

func AddAppUser(userId, password, passwordConfirmation string) (*AppUser, InvalidValue) {
    if len(userId) == 0 || len(password) == 0 || len(passwordConfirmation) == 0 {
        return nil, NewInvalidValue("入力されていない項目があります.")
    }

    if password != passwordConfirmation {
        return nil, NewInvalidValue("パスワードが一致しません.")
    }

    newUser := AppUser {
        UserId: userId,
    }
    newCredential := AppUserCredential{
        AppUser: newUser,
        Password: []byte(password),
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
    if err := tx.Create(newCredential).Error; err != nil {
        panic(err)
    }
    return &newUser, nil
}

const (
	adminUserId = "ah@jabara.info"
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
	tx := _db.Begin().LogMode(true)
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()

	if _, nf := getAppUserByUserId(tx, adminUserId); nf == nil {
		return
	}
	adminUser := AppUser{UserId: adminUserId}
	credential := AppUserCredential{AppUser: adminUser, Password: []byte(adminUserPassword)}
	if err := tx.Create(&credential).Error; err != nil {
		panic(err)
	}
}
