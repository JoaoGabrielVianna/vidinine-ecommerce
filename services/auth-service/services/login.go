package services

import (
	"errors"

	"github.com/vidinine-ecommerce/auth-service/config"
	"github.com/vidinine-ecommerce/auth-service/models"
	"github.com/vidinine-ecommerce/auth-service/utils"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrWrongPassword = errors.New("incorrect password")
)

func Login(email, password string) (string, error) {
	var user models.User

	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			serviceLogger.Error("User not found in the database")
			return "", ErrUserNotFound
		}
		serviceLogger.Error("Database error during user lookup:", err)
		return "", err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		serviceLogger.Error("Incorrect password provided")
		return "", ErrWrongPassword
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		serviceLogger.Error("Failed to generate authentication token:", err)
		return "", err
	}

	serviceLogger.Success("User authenticated successfully")
	return token, nil
}
