package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vidinine-ecommerce/auth-service/services"
)

func DeleteHandler(c *gin.Context) {
	userID, exist := c.Get("userID")
	if !exist {
		controllerLogger.Error("Falha ao deletar conta: usuário não autenticado")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	err := services.DeleteUsers(userID.(uint))
	if err != nil {
		controllerLogger.Error("Erro ao deletar usuário: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	controllerLogger.Success("Conta deletada com sucesso")
	c.JSON(http.StatusOK, gin.H{
		"message": "Conta deletada com sucesso",
	})
}
