package model

import (
	"fmt"
	"github.com/jabaraster/go-web-scaffold/src/go/env"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
    "github.com/jabaraster/go-web-scaffold/src/go/configuration"
)

var (
	_db *gorm.DB
)

func init() {
    config := configuration.Get().Database
    switch (config.Kind) {
    case configuration.DbKind_SQLite:
        db, err := gorm.Open("sqlite3", config.SQLite.DatabaseFilePath)
        if err != nil {
            panic(err)
        }
        _db = &db
        initializeDatabase()
    case configuration.DbKind_PostgreSQL:
        cs := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=false",
            env.ResolveEnv(config.PostgreSQL.Host),
            env.ResolveEnv(config.PostgreSQL.Port),
            env.ResolveEnv(config.PostgreSQL.User),
            env.ResolveEnv(config.PostgreSQL.Database),
            env.ResolveEnv(config.PostgreSQL.Password),
        )
        fmt.Println("★", cs)
        db ,err := gorm.Open("postgres", cs)
        if err != nil {
            panic(err)
        }
        _db = &db
        initializeDatabase()
    }
//	if env.DbKind() == env.DbKindSQLite3 {
//		db, err = gorm.Open("sqlite3", "./app.db")
//	} else {
//		cs := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=%s",
//			env.PostgresHost(),
//			env.PostgresUser(),
//			env.PostgresDbName(),
//			env.PostgresPassword(),
//			env.PostgresSslMode())
//		fmt.Println(cs)
//		db, err = gorm.Open("postgres", cs)
//		if err != nil {
//			panic(err)
//		}
//	}
//	if err != nil {
//		panic(err)
//	}
//	_db = &db
}

func initializeDatabase() {
    _db.LogMode(true)
    _db.CreateTable(&Product{})
    _db.CreateTable(&Order{})
    _db.CreateTable(&AppUser{})
    _db.CreateTable(&AppUserCredential{})
    _db.AutoMigrate(&Product{}, &Order{}, &AppUser{}, &AppUserCredential{})

    createAdminUserIfNecessary()
}

type NotFound interface {
	// nodef
}
type notFoundImpl struct {
	// nodef
}

func NewNotFound() NotFound {
	return notFoundImpl{}
}

type InvalidValue interface {
	GetDescription() string
}

type invalidValue struct {
	description string
}

func (e *invalidValue) GetDescription() string {
	return e.description
}

func NewInvalidValue(description string) InvalidValue {
	return &invalidValue{description: description}
}

func mustInsert(db *gorm.DB, entity interface{}) {
	if err := db.Create(entity).Error; err != nil {
		panic(err)
	}
}
