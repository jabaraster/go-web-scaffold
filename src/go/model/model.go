package model

import (
	"../env"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var (
	_db gorm.DB
)

func init() {
	var db gorm.DB
	var err error
	if env.DbKind() == env.DbKindSQLite3 {
		db, err = gorm.Open("sqlite3", "./app.db")
	} else {
		cs := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=%s",
			env.PostgresHost(),
			env.PostgresUser(),
			env.PostgresDbName(),
			env.PostgresPassword(),
			env.PostgresSslMode())
		fmt.Println(cs)
		db, err = gorm.Open("postgres", cs)
		if err != nil {
			panic(err)
		}
	}
	if err != nil {
		panic(err)
	}
	_db = db
	db.LogMode(true)
	db.CreateTable(&Product{})
	db.CreateTable(&Order{})
	db.CreateTable(&AppUser{})
	db.CreateTable(&AppUserCredential{})
	db.AutoMigrate(&Product{}, &Order{})

    createAdminUserIfNecessary()
}

func UpdateDb() {
	tx := _db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()

	prod := Product{Code: "AAA"}

	var result *gorm.DB
	result = tx.Create(&prod)
	if result.Error != nil {
		panic(result.Error)
	}

	ord := Order{Product: prod}
	result = tx.Create(&ord)
	if result.Error != nil {
		panic(result.Error)
	}
}

func SelectDb() {
	var orders []Order
	_db.Preload("Product").Find(&orders)
	for idx, order := range orders {
		fmt.Println(idx+1, order.Product.Code)
	}
}
