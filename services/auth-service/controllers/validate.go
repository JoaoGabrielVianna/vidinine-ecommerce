package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vidinine-ecommerce/auth-service/utils"
)

func ValidateTokenHandle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer") {
		controllerLogger.Error("Token autorization not provided or invalid format")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not provided or invalid format"})
		return
	}

	tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))

	claims, err := utils.ParseToken(tokenString)
	if err != nil {
		controllerLogger.Error("Failed to parse token: " + err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token or expired"})
		return
	}

	controllerLogger.Successf("Token is valid for user ID: %d", claims.UserID)
	c.JSON(http.StatusOK, gin.H{
		"message": "Token is valid",
		"user-id": claims.UserID,
		"role":    claims.Role,
	})
}
