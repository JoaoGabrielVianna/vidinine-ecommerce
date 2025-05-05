package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vidinine-ecommerce/product-service/services"
)

func ListPublicProductsHandler(c *gin.Context) {
	controllerLogger.Logf("Iniciando requisição GET /products")
	requestStart := time.Now()

	// Obter produtos do service
	products, err := services.ListPublicProducts()
	if err != nil {
		controllerLogger.Errorf("Falha ao processar requisição | Status: %d | Erro: %v | Tempo: %v",
			http.StatusInternalServerError, err.Error(), time.Since(requestStart))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch products",
			"details": err.Error(),
		})
		return
	}

	controllerLogger.Successf("Requisição concluída com sucesso | Status: %d | Produtos: %d | Tempo: %v",
		http.StatusOK, len(products), time.Since(requestStart))

	// Construir resposta final
	c.JSON(http.StatusOK, gin.H{
		"data": products,
		"meta": gin.H{
			"count":     len(products),
			"timestamp": time.Now().Unix(),
		},
	})
}
