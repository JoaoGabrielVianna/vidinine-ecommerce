package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vidinine-ecommerce/product-service/controllers"
	"github.com/vidinine-ecommerce/product-service/middlewares"
)

func SetupRouter(r *gin.Engine) *gin.Engine {

	api := r.Group("/api/v1")
	{
		// 🌐 Rotas públicas
		publicRoutes := api.Group("/products")
		{
			// 📋 Listar produtos
			publicRoutes.GET("/list", controllers.ListPublicProductsHandler)
		}

		// 🔒 Rotas protegidas (usuário autenticado)
		protectedRoutes := api.Group("/")
		protectedRoutes.Use(middlewares.AuthMiddleware())
		{
			// 🚧 Adicione rotas protegidas aqui
		}

		// 🛠️ Rotas administrativas
		adminRoutes := api.Group("/management/products")
		adminRoutes.Use(middlewares.AuthMiddleware(), middlewares.AdminMiddleware())
		{
			// ➕ Criar produto
			adminRoutes.POST("/create", controllers.CreateProduct)
			// 📋 Listar produtos (admin)
			adminRoutes.GET("/list", controllers.ListProductsHandler)
		}
	}

	return r
}
