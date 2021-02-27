package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	ID 			uuid.UUID		`gorm: "primary_key"`
	Email		string			`gorm: "email"`
	Username 	string 			`gorm: "username to login"`
	Password 	string 			`gorm: "password to login"`
	Name		string			`gorm: "name"`
	Date 		int64 			`gorm: "type:datetime as unix"`
	Role 		int16 			`gorm: "role"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (user *User) BeforeCreate() (err error) {
	user.ID = uuid.New()
	if user.Username == "" {
		panic(err)
	}
	return
}

func (user *User) Create(db *gorm.DB) {
	user.Password, _ = HashPassword(user.Password)
	if db.Create(&user).Error != nil{
		log.Panic("Unable to create user")
	}
}

func GetAll(db *gorm.DB) []User {
	var allUsers []User
	db.Find(&allUsers)
	return allUsers
}

func (user *User) Update(db *gorm.DB) {
	user.Password, _ = HashPassword(user.Password)
	if db.Model(&user).Where("id = ?", user.ID).Update(&user).Error != nil {
		log.Panic("No user with that id")
	}
	//if db.Where("id = ?", user.ID).Update(user).Error != nil {
	//	log.Panic("No user with this id")
	//}
}

func (user *User) GetByID(db *gorm.DB, id string) {
	if db.Where("id = ?", id).Find(&user).Error != nil {
		log.Panic("Some problems to find user with ID")
	}
}

func (user *User) Delete(db *gorm.DB)  {
	if db.Where("id = ?", user.ID).Delete(user).Error != nil {
		log.Panic("Some problems to delete user")
	}
}