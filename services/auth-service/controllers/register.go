package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vidinine-ecommerce/aut-service/models"
	"github.com/vidinine-ecommerce/aut-service/services"
)

func RegisterHandler(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		controllerLogger.Errorf("Erro ao fazer bind: %v", err)
		return
	}

	if err := services.RegisterUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		controllerLogger.Errorf("Erro ao registrar usuário: %v", err)
		return
	}

	response := gin.H{
		"msg":    "Usuário registrado com sucesso",
		"status": http.StatusCreated,
		"data": gin.H{
			"id":         user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"password":   user.Password,
			"created_at": user.CreatedAt,
		},
	}

	c.JSON(http.StatusCreated, response)
	controllerLogger.Successf("Usuário registrado com sucesso: %s", user.Email)
}
