package services

import (
	"fmt"

	"github.com/vidinine-ecommerce/auth-service/config"
	"github.com/vidinine-ecommerce/auth-service/models"
)

func DeleteUsers(userID uint) error {
	var user models.User
	result := config.DB.First(&user, userID)
	if result.Error != nil {
		serviceLogger.Error("Erro ao buscar usuário para exclusão")
		return result.Error
	}

	updates := map[string]interface{}{
		"Name":     "User Deleted",
		"Email":    fmt.Sprintf("deleted_%d@deleted.com", userID),
		"Password": "",
	}

	if err := config.DB.Model(&user).Updates(updates).Error; err != nil {
		serviceLogger.Errorf("Erro ao anonimizar usuário (ID: %d): %v", userID, err)
		return fmt.Errorf("erro ao anonimizar dados do usuário")
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		serviceLogger.Error("Erro ao excluir usuário do banco de dados")
		return err
	}

	serviceLogger.Success("Usuário excluído com sucesso")
	return nil
}
