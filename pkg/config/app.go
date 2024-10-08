package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {

	//d, err := gorm.Open("mysql", "badr:ba0912dr/simplerest?charset=utf8&parseTime=True&loc=Local")
    d, err := gorm.Open("mysql", "badr:ba0912dr@tcp(localhost:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
