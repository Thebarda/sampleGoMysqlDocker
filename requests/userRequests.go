package requests

import (
	"../dbConfig"
	"../models"
)

type UserRequests struct{}

func GetAllUsers() []models.User {
	db := dbConfig.GetDb()
	var users []models.User
	db.Find(&users)
	return users
}

func Insert(user models.User) bool {
	db := dbConfig.GetDb()
	res := db.NewRecord(user)
	db.Create(&user)
	return res
}
