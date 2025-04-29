package services

import (
	"github.com/vidinine-ecommerce/aut-service/config"
	"github.com/vidinine-ecommerce/aut-service/models"
	"github.com/vidinine-ecommerce/aut-service/utils"
)

func RegisterUser(user *models.User) error {
	if err := utils.ValidateEmail(user.Email); err != nil {
		return err
	}
	if err := utils.ValidatePassword(user.Password); err != nil {
		return err
	}

	if err := config.DB.Create(user).Error; err != nil {
		serviceLogger.Errorf("Erro ao registrar: %v", err)
		return err
	}
	return nil
}
