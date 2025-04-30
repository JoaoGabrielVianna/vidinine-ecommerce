package middlewares

import (
	"net/http"
	"strings" // Ajuste o caminho conforme sua estrutura

	"github.com/gin-gonic/gin"
	"github.com/vidinine-ecommerce/aut-service/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ParseToken(tokenString)
		if err != nil {

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		// 4. Adicionar o userID ao contexto do Gin
		c.Set("userID", claims.UserID)

		// 5. Continuar para o próximo handler
		c.Next()
	}
}
