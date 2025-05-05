package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vidinine-ecommerce/product-service/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			middlewaresLogger.Error("Cabeçalho de autorização está ausente")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			middlewaresLogger.Error("Falha ao analisar o token: " + err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		// Adicionar o userID ao contexto do Gin
		c.Set("userID", claims.UserID)

		// Adicionar o userRole ao contexto do Gin após validar o token
		c.Set("userRole", claims.Role)

		// Continuar para o próximo handler
		middlewaresLogger.Successf("Token validado com sucesso para userID: %d, role: %s", claims.UserID, claims.Role)
		c.Next()
	}
}
