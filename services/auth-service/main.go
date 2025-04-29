package main

import (
	"github.com/vidinine-ecommerce/aut-service/config"
	"github.com/vidinine-ecommerce/aut-service/models"
	"github.com/vidinine-ecommerce/aut-service/routers"
)

func main() {
	// Configurações
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})

	// Rotas
	r := routers.SetupRouter()
	r.Run(":3000")
}
