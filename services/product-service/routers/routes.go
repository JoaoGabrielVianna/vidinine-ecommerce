package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vidinine-ecommerce/product-service/controllers"
	"github.com/vidinine-ecommerce/product-service/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Welcome to the API",
			})
		})

		// Rotas protegidas
		protected := api.Group("/")
		protected.Use(middlewares.AuthMiddleware())

		admin := api.Group("/admin")
		admin.Use(middlewares.AuthMiddleware(), middlewares.AdminMiddleware())
		{
			admin.POST("/products", controllers.CreateProduct)
		}
	}

	return r
}
