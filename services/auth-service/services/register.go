package services

import (
	"github.com/vidinine-ecommerce/auth-service/config"
	"github.com/vidinine-ecommerce/auth-service/models"
	"github.com/vidinine-ecommerce/auth-service/utils"
)

func RegisterUser(user *models.User) error {
	if err := utils.ValidateEmail(user.Email); err != nil {
		serviceLogger.Error("Incorrect email format")
		return err
	}
	if err := utils.ValidatePassword(user.Password); err != nil {
		serviceLogger.Error("Password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one number, and one special character")
		return err
	}

	// Criptografar a senha
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		serviceLogger.Error("Error hashing password:", err)
		return nil
	}
	user.Password = hashedPassword

	if err := config.DB.Create(user).Error; err != nil {
		serviceLogger.Errorf("Error registering user: %v", err)
		return err
	}

	serviceLogger.Success("User registered successfully")
	return nil
}
