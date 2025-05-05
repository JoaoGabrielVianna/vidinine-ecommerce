package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vidinine-ecommerce/product-service/services"
)

func ListProductsHandler(c *gin.Context) {
	products, err := services.ListProducts()

	controllerLogger.Logf("Iniciando requisição GET /products")
	requestStart := time.Now()

	if err != nil {
		controllerLogger.Errorf("Falha ao processar requisição | Status: %d | Erro: %v | Tempo: %v",
			http.StatusInternalServerError, err.Error(), time.Since(requestStart))

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Falha ao buscar produtos",
			"details": err.Error(),
		})
		return
	}

	controllerLogger.Successf("Requisição concluída com sucesso | Status: %d | Produtos: %d | Tempo: %v",
		http.StatusOK, len(products), time.Since(requestStart))

	c.JSON(http.StatusOK, gin.H{
		"data": products,
		"meta": gin.H{
			"count":     len(products),
			"timestamp": time.Now().Unix(),
		},
	})
}
