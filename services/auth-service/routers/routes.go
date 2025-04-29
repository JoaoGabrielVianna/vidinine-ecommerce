package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vidinine-ecommerce/aut-service/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.GET("/register", controllers.RegisterHandler)
	}
	return r
}
