package config

import (
	"fmt"
	"os"

	"github.com/vidinine-ecommerce/product-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB 🔌
// Função responsável por conectar ao banco de dados PostgreSQL.
// Carrega as variáveis de ambiente, configura a string de conexão e inicializa o GORM.
// Retorna um erro caso a conexão falhe.
func ConnectDB() error {
	dblogger := GetLogger("Database")

	dblogger.Log("🔄 Iniciando conexão com o banco de dados...")

	// Configura a string de conexão com o banco de dados 🛠️
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		"localhost",              // Host do banco de dados (atualmente fixado como localhost)
		os.Getenv("DB_USER"),     // Usuário do banco de dados
		os.Getenv("DB_PASSWORD"), // Senha do banco de dados
		os.Getenv("DB_NAME"),     // Nome do banco de dados
		os.Getenv("DB_PORT"),     // Porta do banco de dados
		os.Getenv("DB_SSLMODE"),  // Modo SSL
	)

	// Tenta abrir a conexão com o banco de dados 🚀
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		dblogger.Errorf("❌ Erro ao conectar ao banco de dados: %v", err)
		return err
	}

	// Define a variável global DB 🌍
	DB = db
	dblogger.Successf("✅ Conexão estabelecida: %s", os.Getenv("DB_NAME"))
	return nil
}

// checkDatabase 🛡️
// Função para verificar e garantir que a tabela 'products' existe no banco de dados.
// Caso não exista, ela será criada automaticamente.
func checkDatabase() {
	// Verifica se a tabela 'products' existe no banco de dados 🔍
	if !DB.Migrator().HasTable(&models.Product{}) {
		configLogger.Error("⚠️ Tabela 'products' em falta, criando...")

		// Tenta criar a tabela 'products' 🏗️
		if err := DB.AutoMigrate(&models.Product{}); err != nil {
			configLogger.Errorf("❌ Falha ao criar tabela: %v", err)
			return
		}

		configLogger.Success("✅ Tabela 'products' criada com sucesso\n")
	} else {
		configLogger.Success("✅ Tabela 'products' encontrada e validada\n")
	}
}
