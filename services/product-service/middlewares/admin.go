package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vidinine-ecommerce/product-service/models"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		if !exists || userRole != models.AdminRole {
			middlewaresLogger.Error("Usuário não possui privilégios de administrador ou userRole está ausente")
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Acesso restrito a administradores",
			})
			return
		}
		c.Next()
	}
}
