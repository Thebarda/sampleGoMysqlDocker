package requests

import (
	"../dbConfig"
	"../models"
)

type UserRequests struct{}

func GetAllUsers() []models.User {
	db := dbConfig.GetDb()

	rows, err := db.Query("SELECT * FROM user") // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	var users []models.User
	for rows.Next() {
		var u models.User
		err = rows.Scan(&u.Id, &u.Firstname)
		if err != nil {
			panic(err)
		}
		users = append(users, u)
	}
	return users
}
