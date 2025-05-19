package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vidinine-ecommerce/auth-service/utils"
)

func ValidateTokenHandle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autorização não fornecido ou formato inválido"})
		return
	}

	tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))

	fmt.Println("Token Recebido: ", tokenString)
	claims, err := utils.ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido ou expirado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Token válido",
		"user-id": claims.UserID,
		"role":    claims.Role,
	})
}
