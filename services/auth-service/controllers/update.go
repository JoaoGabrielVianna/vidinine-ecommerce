package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vidinine-ecommerce/auth-service/models"
	"github.com/vidinine-ecommerce/auth-service/services"
)

func UpdateHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		controllerLogger.Error("Usuário não autenticado")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	var input models.UpdateUser
	if err := c.ShouldBindJSON(&input); err != nil {
		controllerLogger.Error("Erro ao validar dados do usuário: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	updatedUser, err := services.UpdateUser(userID.(uint), input)
	if err != nil {
		controllerLogger.Error("Erro ao atualizar usuário: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	controllerLogger.Success("Usuário atualizado com sucesso")
	c.JSON(http.StatusOK, gin.H{
		"message": "Usuário atualizado com sucesso",
		"data": gin.H{
			"id":    updatedUser.ID,
			"name":  updatedUser.Name,
			"email": updatedUser.Email,
		},
	})
}
