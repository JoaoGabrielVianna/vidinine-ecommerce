package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vidinine-ecommerce/aut-service/utils"
)

func HomeHandler(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	var message string

	if tokenString != "" { // 👈 Corrigido: verifica se TEM token
		claims, err := utils.ParseToken(tokenString)
		if err == nil {
			message = fmt.Sprintf("Bem-vindo ao Auth-Service! (Logado com ID: %d)", claims.UserID) // 👈 UserID em vez de ID
		} else {
			message = "Bem-vindo ao Auth-Service! (Token inválido)"
		}
	} else {
		message = "Bem-vindo ao Auth-Service! (Não logado)"
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"status":  http.StatusOK,
		"routes": gin.H{
			"POST /register": "Register a new user",
			"POST /login":    "Login with email and password",
		},
	})
}
