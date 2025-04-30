package main

import (
	"github.com/vidinine-ecommerce/auth-service/config"
	"github.com/vidinine-ecommerce/auth-service/models"
	"github.com/vidinine-ecommerce/auth-service/routers"
)

func main() {
	// Configurações
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})

	// Rotas
	r := routers.SetupRouter()
	r.Run(":3000")
}
