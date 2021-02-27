package utils

import (
	"errors"
	"github.com/LILSunflower/DEKKER/models"
)

func MigrateDB() error{
	userModel := models.User{
		Email : "test",
		Username : "test",
		Password : "admin123",
		Name : "test",
		Date : 1150416000,
		Role : 16,
	}
	if db.AutoMigrate(userModel).Error != nil {
		return  errors.New("CAN NOT MIGRATE")
	}
	return nil
}
