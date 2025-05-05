package services

import (
	"time"

	"github.com/vidinine-ecommerce/product-service/config"
	"github.com/vidinine-ecommerce/product-service/models"
)

func ListPublicProducts() ([]models.PublicProductResponse, error) {
	var products []models.Product

	serviceLogger.Logf("Iniciando busca de produtos no banco de dados")
	startTime := time.Now()

	if err := config.DB.Select("id, name, description, price, stock").Find(&products).Error; err != nil {
		serviceLogger.Errorf("Falha na query de produtos | Erro: %v | Tempo: %v",
			err.Error(), time.Since(startTime))
		return nil, err
	}

	response := make([]models.PublicProductResponse, len(products))
	for i, p := range products {
		response[i] = models.PublicProductResponse{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Available:   p.Stock > 0,
		}
	}

	serviceLogger.Successf("Produtos recuperados com sucesso | Total: %d | Tempo: %v",
		len(products), time.Since(startTime))

	return response, nil
}
