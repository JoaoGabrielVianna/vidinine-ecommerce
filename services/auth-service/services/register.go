package services

import (
	"github.com/vidinine-ecommerce/auth-service/config"
	"github.com/vidinine-ecommerce/auth-service/models"
	"github.com/vidinine-ecommerce/auth-service/utils"
)

func RegisterUser(user *models.User) error {
	if err := utils.ValidateEmail(user.Email); err != nil {
		serviceLogger.Error("Erro ao validar o email")
		return err
	}
	if err := utils.ValidatePassword(user.Password); err != nil {
		serviceLogger.Error("Erro ao validar a senha")
		return err
	}

	// Criptografar a senha
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		serviceLogger.Error("Erro ao criptografar a senha")
		return nil
	}
	user.Password = hashedPassword

	if err := config.DB.Create(user).Error; err != nil {
		serviceLogger.Errorf("Erro ao registrar: %v", err)
		return err
	}

	serviceLogger.Success("Usuário registrado com sucesso")
	return nil
}
