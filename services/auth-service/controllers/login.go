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
		c.JSON(400, gin.H{"error": "Dados inválidos"})
		return
	}

	token, err := services.Login(creds.Email, creds.Password)
	switch err {
	case services.ErrUserNotFound:
		c.JSON(http.StatusNotFound, gin.H{ // 404 ou 401
			"error":   "usuario_nao_encontrado",
			"message": "Nenhuma conta com este email foi encontrada",
		})
		return
	case services.ErrWrongPassword:
		c.JSON(http.StatusUnauthorized, gin.H{ // 401
			"error":   "senha_incorreta",
			"message": "A senha está incorreta",
		})
		return
	case nil:
		// Token gerado com sucesso
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "erro_interno",
			"message": "Ocorreu um erro inesperado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": "Bearer " + token,
	})
}
