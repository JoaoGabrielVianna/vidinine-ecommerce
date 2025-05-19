package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vidinine-ecommerce/auth-service/config"
	"github.com/vidinine-ecommerce/auth-service/routers"
)

func main() {
	config.Init()
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	r := gin.New()
	// Rotas
	routers.SetupRouter(r)
	r.Run(":3000")
}
