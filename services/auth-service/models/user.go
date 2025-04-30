package models

import "gorm.io/gorm"

type User struct {
	// 🆔 ID
	// 🕒 CreatedAt
	// 🕒 UpdatedAt
	// 🗑️ DeletedAt.
	gorm.Model

	Name     string `json:"name" gorm:"not null"`              // 📝 Name of the user
	Email    string `json:"email" gorm:"not null;unique"`      // 📧 Email of the user
	Password string `json:"password" gorm:"not null;size:255"` // 🔒 Password of the user
}

type UpdateUser struct {
	Name  string `json:"name" binding:"omitempty,min=2"`  // Obrigatório se enviado
	Email string `json:"email" binding:"omitempty,email"` // Valida formato de email
}
