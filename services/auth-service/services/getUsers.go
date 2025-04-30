package services

import (
	"github.com/vidinine-ecommerce/auth-service/config"
	"github.com/vidinine-ecommerce/auth-service/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	result := config.DB.Select("id, name, email, role, created_at").Find(&users)
	return users, result.Error
}
