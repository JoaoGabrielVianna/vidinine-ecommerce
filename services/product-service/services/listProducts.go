package services

import (
	"time"

	"github.com/vidinine-ecommerce/product-service/config"
	"github.com/vidinine-ecommerce/product-service/models"
)

func ListProducts() ([]models.Product, error) {
	var products []models.Product

	serviceLogger.Logf("Iniciando busca de produtos no banco de dados")
	startTime := time.Now()

	if err := config.DB.Find(&products).Error; err != nil {
		serviceLogger.Errorf("Falha na query de produtos | Erro: %v | Tempo: %v",
			err.Error(), time.Since(startTime))
		return nil, err
	}

	serviceLogger.Successf("Produtos recuperados com sucesso | Total: %d | Tempo: %v",
		len(products), time.Since(startTime))

	return products, nil
}
