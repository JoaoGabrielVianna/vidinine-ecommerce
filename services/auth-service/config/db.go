package config

import (
	"fmt"
	"os"

	"github.com/vidinine-ecommerce/auth-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	l "gorm.io/gorm/logger"
)

// ConnectDB 🔌
// Função responsável por conectar ao banco de dados PostgreSQL.
// Carrega as variáveis de ambiente, configura a string de conexão e inicializa o GORM.
// Retorna um erro caso a conexão falhe.
func ConnectDB() error {
	dblogger := GetLogger("Database")

	// Carrega as variáveis de ambiente 🌱
	LoadEnv("../..")

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
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: l.Default.LogMode(l.Silent),
	})
	if err != nil {
		dblogger.Errorf("❌ Erro ao conectar ao banco de dados: %v", err)
		return err
	}

	// Define a variável global DB 🌍
	DB = db
	dblogger.Successf("✅ Conexão com o banco de dados estabelecida com sucesso: %s", os.Getenv("DB_NAME"))
	return nil
}

// checkDatabase 🛡️
// Função para verificar e garantir que a tabela 'users' existe no banco de dados.
// Caso não exista, ela será criada automaticamente.
func checkDatabase() {
	// Verifica se a tabela 'users' existe no banco de dados 🔍
	if !DB.Migrator().HasTable(&models.User{}) {
		configLogger.Error("⚠️ Tabela 'users' em falta, criando...")

		// Tenta criar a tabela 'users' 🏗️
		if err := DB.AutoMigrate(&models.User{}); err != nil {
			configLogger.Errorf("❌ Falha ao criar tabela: %v", err)
			return
		}

		configLogger.Success("✅ Tabela 'users' criada com sucesso\n")
	} else {
		configLogger.Success("✅ Tabela 'users' encontrada e validada\n")
	}
}
