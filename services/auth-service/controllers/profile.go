// auth-service/controllers/profile.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vidinine-ecommerce/auth-service/config"
	"github.com/vidinine-ecommerce/auth-service/models"
)

func ProfileHandler(c *gin.Context) {
	// 1. Pegar userID (como uint) do contexto
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	// 2. Buscar usuário no banco (GORM)
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}
	// 3. Retornar dados seguros (sem senha)
	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"created_at": user.CreatedAt,
	})
}
