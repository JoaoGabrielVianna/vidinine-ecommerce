package services

import (
	"github.com/vidinine-ecommerce/auth-service/config"
	"github.com/vidinine-ecommerce/auth-service/models"
	"github.com/vidinine-ecommerce/auth-service/utils"
)

func RegisterUser(user *models.User) error {
	if err := utils.ValidateEmail(user.Email); err != nil {
		return err
	}
	if err := utils.ValidatePassword(user.Password); err != nil {
		return err
	}

	// Criptografar a senha
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil
	}
	user.Password = hashedPassword

	if err := config.DB.Create(user).Error; err != nil {
		serviceLogger.Errorf("Erro ao registrar: %v", err)
		return err
	}
	return nil
}
