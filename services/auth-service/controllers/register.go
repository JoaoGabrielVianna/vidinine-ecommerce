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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		controllerLogger.Errorf("Erro ao fazer bind: %v", err)
		return
	}

	if err := services.RegisterUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		controllerLogger.Errorf("Erro ao registrar usuário: %v", err)
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		controllerLogger.Errorf("Erro ao gerar token: %v", err)
		return
	}
	c.Header("Authorization", "Bearer "+token)

	response := gin.H{
		"msg":    "Usuário registrado com sucesso",
		"status": http.StatusCreated,
		"data": gin.H{
			"id":         user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"password":   user.Password,
			"created_at": user.CreatedAt,
			"role":       user.Role,
		},
		"token": "Bearer " + token,
	}

	c.JSON(http.StatusCreated, response)
	controllerLogger.Successf("Usuário registrado com sucesso: %s", user.Email)
}
