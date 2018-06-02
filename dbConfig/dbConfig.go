package dbConfig

import (
	"../models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
}

func GetDb() (db *gorm.DB) {
	db, err := gorm.Open("mysql", "root:barda87@tcp(mysql)/sample")
	db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
	return
}
