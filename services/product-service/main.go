package main

import (
	"github.com/vidinine-ecommerce/product-service/config"
	"github.com/vidinine-ecommerce/product-service/routers"
)

func main() {
	config.Init()

	// Rotas
	r := routers.SetupRouter()
	r.Run(":3001")
}
