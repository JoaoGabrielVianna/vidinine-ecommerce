package services

import (
	"fmt"
	"strings"

	"github.com/vidinine-ecommerce/product-service/config"
	"github.com/vidinine-ecommerce/product-service/models"
)

func CreateProduct(product *models.Product) error {
	if err := config.DB.Create(product); err != nil {

		if err.Error != nil {
			if isDuplicateKeyError(err.Error) {
				serviceLogger.Errorf("Tentativa de criar produto duplicado: %s", product.Name)
				return fmt.Errorf("já existe um produto com o nome '%s'", product.Name)
			}

			serviceLogger.Errorf("Falha no banco de dados: %v", err.Error)
			return fmt.Errorf("erro interno ao criar produto")
		}
	}
	serviceLogger.Successf("Produto criado (ID: %d, Nome: %s)", product.ID, product.Name)
	return nil
}

// Função auxiliar para detectar erros de chave única
func isDuplicateKeyError(err error) bool {
	return strings.Contains(err.Error(), "duplicate key") ||
		strings.Contains(err.Error(), "violates unique constraint")
}
