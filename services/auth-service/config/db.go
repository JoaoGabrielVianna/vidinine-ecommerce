package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	dblogger := GetLogger("Database")

	LoadEnv("../..")

	dblogger.Log("Iniciando conexão com o banco de dados...")

	// Set up the database connection string
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		// os.Getenv("DB_HOST"),
		"localhost",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		dblogger.Errorf("Erro ao conectar ao banco de dados: %v", err)
		return err
	}
	// Set the global DB variable
	DB = db
	dblogger.Successf("Conexão com o banco de dados estabelecida com sucesso: %s", os.Getenv("DB_NAME"))
	return nil
}
