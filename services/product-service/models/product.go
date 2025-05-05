package models

import (
	"gorm.io/gorm"
)

type Product struct {
	// 🆔 ID
	// 🕒 CreatedAt
	// 🕒 UpdatedAt
	// 🗑️ DeletedAt.
	gorm.Model

	Name        string  `json:"name" gorm:"not null;unique"`     // 📦 Nome do produto
	Description string  `json:"description" gorm:"type:text"`    // 📝 Descrição do produto
	Price       float64 `json:"price" gorm:"not null"`           // 💰 Preço do produto
	Stock       int     `json:"stock" gorm:"not null;default:0"` // 📦 Quantidade em estoque
	ImageURL    string  `json:"image_url" gorm:"type:text"`      // 🖼️ URL da imagem do produto
	// CategoryID  uint    `json:"category_id"`                     // 🏷️ Categoria (FK futura, opcional por enquanto)
}

// ProductPublicResponse - Estrutura para resposta pública
type PublicProductResponse struct {
	ID          uint    `json:"id"`          // 🆔 ID do produto
	Name        string  `json:"name"`        // 📦 Nome do produto
	Description string  `json:"description"` // 📝 Descrição do produto
	Price       float64 `json:"price"`       // 💰 Preço do produto
	Available   bool    `json:"available"`   // ✅ Disponibilidade (true/false)
}
