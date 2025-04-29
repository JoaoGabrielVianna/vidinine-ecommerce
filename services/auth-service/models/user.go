package models

import "gorm.io/gorm"

type User struct {
	// 🆔 ID
	// 🕒 CreatedAt
	// 🕒 UpdatedAt
	// 🗑️ DeletedAt.
	gorm.Model

	Name     string `json:"name" gorm:"not null"`         // 📝 Name of the user
	Email    string `json:"email" gorm:"not null;unique"` // 📧 Email of the user
	Password string `json:"password" gorm:"not null"`     // 🔒 Password of the user
}
