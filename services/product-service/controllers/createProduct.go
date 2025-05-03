package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vidinine-ecommerce/product-service/models"
	"github.com/vidinine-ecommerce/product-service/services"
)

func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindBodyWithJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados Inválidos"})
		controllerLogger.Errorf("Erro ao fazer bind: %v", err)
		return
	}

	if err := services.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		controllerLogger.Errorf("Erro ao registrar produto: %v", err)
		return
	}

	response := gin.H{
		"msg":    "Produto criado com sucesso",
		"status": 200,
		"data": gin.H{
			"id":          product.ID,
			"name":        product.Name,
			"description": product.Description,
			"price":       product.Price,
			"stock":       product.Stock,
			"image_url":   product.ImageURL,
			"created_at":  product.CreatedAt,
			"updated_at":  product.UpdatedAt,
		},
	}

	c.JSON(http.StatusOK, response)
	controllerLogger.Successf("Produto registrado com sucesso: %s", product.Name)

}
