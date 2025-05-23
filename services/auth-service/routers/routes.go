package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vidinine-ecommerce/auth-service/controllers"
	"github.com/vidinine-ecommerce/auth-service/middlewares"
)

func SetupRouter(r *gin.Engine) *gin.Engine {
	r.Use(middlewares.CORSMiddleware())

	api := r.Group("/api/v1")
	{

		api.GET("/register", controllers.RegisterHandler)
		api.POST("/login", controllers.LoginHandler)
		api.GET("/validate", controllers.ValidateTokenHandle) // 🔥 aqui!

		// Rotas protegidas
		protected := api.Group("/")
		protected.Use(middlewares.AuthMiddleware())
		{
			protected.GET("/profile", controllers.ProfileHandler)
			protected.DELETE("/delete", controllers.DeleteHandler)
			protected.PUT("/update", controllers.UpdateHandler)
		}

		admin := api.Group("/admin")
		admin.Use(middlewares.AuthMiddleware(), middlewares.AdminMiddleware())
		{
			admin.GET("/users", controllers.GetUsers)
		}
	}
	return r
}
