package services

import (
	"fmt"
	"strings"

	"github.com/vidinine-ecommerce/product-service/config"
	"github.com/vidinine-ecommerce/product-service/models"
)

func CreateProduct(product *models.Product) error {
	if err := config.DB.Create(product).Error; err != nil {
		if strings.Contains(err.Error(), "valor de chave duplicada viola a restrição de unicidade") {
			serviceLogger.Errorf("Produto com nome duplicado: %s", product.Name)
			return fmt.Errorf("já existe um produto com este nome")
		}
		serviceLogger.Errorf("Erro ao registrar: %v", err)
		return err
	}
	serviceLogger.Success("Produto registrado com sucesso")
	return nil
}
