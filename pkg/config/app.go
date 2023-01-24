package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysqL"
)

var (
	db *gorm.DB
)

func Connect() {

	d, err := gorm.Open("mysql", "atanda0x:ethereumsolanasimplerest?charset=utf8&parseTime=True&loc=local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
