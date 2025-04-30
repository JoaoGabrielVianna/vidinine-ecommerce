package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vidinine-ecommerce/aut-service/config"
	"github.com/vidinine-ecommerce/aut-service/models"
)

func HomeHandler(c *gin.Context) {
	// Obter o userID do contexto (setado pelo AuthMiddleware)
	userID, exists := c.Get("userID")
	var message string

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	if exists {
		message = fmt.Sprintf("Bem-vindo ao Auth-Service! (Logado com : %d)", userID)
	} else {
		message = "Bem-vindo ao Auth-Service! (Não logado)"
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"status":  http.StatusOK,
		"user": gin.H{
			"id":         user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"created_at": user.CreatedAt,
		},
	})
}
