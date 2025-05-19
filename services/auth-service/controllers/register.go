package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vidinine-ecommerce/auth-service/models"
	"github.com/vidinine-ecommerce/auth-service/services"
	"github.com/vidinine-ecommerce/auth-service/utils"
)

func RegisterHandler(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		controllerLogger.Errorf("Failed to bind request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid_request",
			"message": "Invalid request data",
		})
		return
	}

	if err := services.RegisterUser(&user); err != nil {
		controllerLogger.Errorf("Failed to register user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "registration_failed",
			"message": "Could not register user",
		})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		controllerLogger.Errorf("Failed to generate token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "token_generation_failed",
			"message": "Could not generate authentication token",
		})
		return
	}
	c.Header("Authorization", "Bearer "+token)

	response := gin.H{
		"message": "User registered successfully",
		"status":  http.StatusCreated,
		"data": gin.H{
			"id":         user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"created_at": user.CreatedAt,
			"role":       user.Role,
		},
		"token":      token,
		"token_type": "Bearer",
	}

	controllerLogger.Successf("User registered successfully: %s", user.Email)
	c.JSON(http.StatusCreated, response)
}
