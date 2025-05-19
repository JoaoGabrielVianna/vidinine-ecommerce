package services

import (
	"errors"

	"github.com/vidinine-ecommerce/auth-service/config"
	"github.com/vidinine-ecommerce/auth-service/models"
	"github.com/vidinine-ecommerce/auth-service/utils"
)

func UpdateUser(userID uint, input models.UpdateUser) (*models.User, error) {
	var user models.User

	if err := config.DB.First(&user, userID).Error; err != nil {
		serviceLogger.Errorf("Failed to find user in database (ID: %d): %v", userID, err)
		return nil, err
	}

	if input.Email != "" && input.Email != user.Email {
		var existingUser models.User
		if err := config.DB.Where("email = ? AND id != ?", input.Email, userID).First(&existingUser).Error; err == nil {
			serviceLogger.Errorf("Attempt to update to an email already in use (ID: %d, Email: %s)", userID, input.Email)
			return nil, errors.New("email is already in use by another user")
		}
	}

	updates := map[string]interface{}{}
	if input.Name != "" {
		updates["name"] = input.Name
	}
	if input.Email != "" {
		updates["email"] = input.Email
	}
	if input.Password != "" {
		// Criptografar a senha
		hashedPassword, err := utils.HashPassword(input.Password)
		if err != nil {
			serviceLogger.Errorf("Failed to hash password for user ID %d: %v", userID, err)
			return nil, errors.New("failed to update user password")
		}
		updates["password"] = hashedPassword
	}

	if err := config.DB.Model(&user).Updates(updates).Error; err != nil {
		serviceLogger.Errorf("Failed to update user data (ID: %d): %v", userID, err)
		return nil, err
	}

	serviceLogger.Successf("User updated successfully (ID: %d)", userID)
	return &user, nil
}
