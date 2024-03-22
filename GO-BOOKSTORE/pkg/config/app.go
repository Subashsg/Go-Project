package config

import (
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "mysql", "root:salahari@1121996@tcp(127.0.0.1:3306)/gobookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
