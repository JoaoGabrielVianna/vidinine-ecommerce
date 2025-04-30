package config

import (
	"gorm.io/gorm"
)

var (
	logger       *Logger
	DB           *gorm.DB
	configLogger = GetLogger("config")
)

// Init inicializa as configurações do serviço de autenticação. 🚀✨
//
// Esta função realiza as seguintes etapas:
//
// 1. Exibe uma mensagem de log indicando o início das configurações. 🛠️
//
// 2. Conecta ao banco de dados utilizando a função ConnectDB. 🗄️
//
// 3. Verifica o estado do banco de dados com a função checkDatabase. ✅
//
// 4. Exibe uma mensagem de sucesso ao concluir as configurações. 🎉
//
// É essencial chamar esta função no início da aplicação para garantir que
// todas as dependências e configurações estejam prontas para uso. 💡
func Init() {
	configLogger.Log("Iniciando configurações...")
	ConnectDB()
	checkDatabase()
	configLogger.Success("Configuração completa!")
}
