package services

import (
	"errors"

	"github.com/vidinine-ecommerce/auth-service/config"
	"github.com/vidinine-ecommerce/auth-service/models"
)

func UpdateUser(userID uint, input models.UpdateUser) (*models.User, error) {
	var user models.User

	if err := config.DB.First(&user, userID).Error; err != nil {
		serviceLogger.Error("Erro ao buscar o usuário no banco de dados")
		return nil, err
	}

	if input.Email != "" && input.Email != user.Email {
		var existingUser models.User
		if err := config.DB.Where("email = ? AND id != ?", input.Email, userID).First(&existingUser).Error; err == nil {
			serviceLogger.Errorf("Tentativa de atualizar para e-mail em uso (ID: %d, Email: %s)", userID, input.Email)
			return nil, errors.New("email já está em uso por outro usuário")
		}
	}

	updates := map[string]interface{}{}
	if input.Name != "" {
		updates["name"] = input.Name
	}
	if input.Email != "" {
		updates["email"] = input.Email
	}

	if err := config.DB.Model(&user).Updates(updates).Error; err != nil {
		serviceLogger.Error("Erro ao atualizar os dados do usuário")
		return nil, err
	}

	serviceLogger.Success("Usuário atualizado com sucesso")
	return &user, nil
}
