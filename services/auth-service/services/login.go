package services

import (
	"errors"

	"github.com/vidinine-ecommerce/auth-service/config"
	"github.com/vidinine-ecommerce/auth-service/models"
	"github.com/vidinine-ecommerce/auth-service/utils"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound  = errors.New("usuário não encontrado")
	ErrWrongPassword = errors.New("senha incorreta")
)

func Login(email, password string) (string, error) {
	var user models.User

	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			serviceLogger.Error("Usuário não encontrado no banco de dados")
			return "", ErrUserNotFound
		}
		return "", err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		serviceLogger.Error("Senha incorreta fornecida pelo usuário")
		return "", ErrWrongPassword
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	serviceLogger.Success("Login realizado com sucesso")
	return token, nil
}
