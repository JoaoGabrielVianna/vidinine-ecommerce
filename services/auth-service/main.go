package main

import (
	"github.com/vidinine-ecommerce/auth-service/config"
	"github.com/vidinine-ecommerce/auth-service/routers"
)

func main() {
	config.Init()

	// Rotas
	r := routers.SetupRouter()
	r.Run(":3000")
}
