package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vidinine-ecommerce/product-service/config"
	"github.com/vidinine-ecommerce/product-service/models"
)

var (
	AdminMiddlewareLogger = config.GetLogger("AdminMiddlewareLogger")
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		AdminMiddlewareLogger.Logf("userRole: %v, exists: %v\n", userRole, exists)
		if !exists || userRole != models.AdminRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Acesso restrito a administradores",
			})
			return
		}
		c.Next()
	}
}
