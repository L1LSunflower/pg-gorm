package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	ID 			int ""
	username 	string ""
	password 	string ""
	date 		string ""
	pruveleguu 	int16 ""
}

func AddUser(db *gorm.DB) string {
	return "Method: POST, Add users"
}

func GetAllUsers(db *gorm.DB) string {
	return "Method: GET, All users"
}

func UpdateUser(db *gorm.DB) string {
	return "Method PUT, Update users"
}

func GetUserByID(db *gorm.DB) string {
	return  "Method: GET, Get user by id"
}

func DeleteUser(db *gorm.DB) string {
	return "Method: DELETE, Delete user by id"
}