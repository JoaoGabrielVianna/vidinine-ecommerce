package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vidinine-ecommerce/auth-service/services"
)

func LoginHandler(c *gin.Context) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindBodyWithJSON(&creds); err != nil {
		controllerLogger.Error("Failed to bind request body")
		c.JSON(400, gin.H{
			"error":   "invalid_request",
			"message": "Invalid request data",
		})
		return
	}

	token, err := services.Login(creds.Email, creds.Password)
	switch err {
	case services.ErrUserNotFound:
		controllerLogger.Error("User not found")
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "user_not_found",
			"message": "No account found with the provided email",
		})
		return
	case services.ErrWrongPassword:
		controllerLogger.Error("Incorrect password provided")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "incorrect_password",
			"message": "The password provided is incorrect",
		})
		return
	case nil:
		// Token gerado com sucesso
	default:
		controllerLogger.Error("Unexpected internal error")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal_error",
			"message": "An unexpected error occurred",
		})
		return
	}

	controllerLogger.Success("Login successful")
	c.JSON(http.StatusOK, gin.H{
		"token":      token,
		"token_type": "Bearer",
	})
}
