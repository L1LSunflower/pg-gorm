package utils

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB //database

type dbConnection struct {
	HOST     string "env: database host"
	PORT     int    "env: database port"
	USER     string "env: database user"
	PASSWORD string "env: database password"
	DBNAME   string "env: database name"
}

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
	}

	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Println(err)
	}

	db = conn
	// db.Debug().AutoMigrate()

}

func GetDB() *gorm.DB {
	return db
}
