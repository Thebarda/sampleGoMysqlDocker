package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Id        int
	Firstname string
}
