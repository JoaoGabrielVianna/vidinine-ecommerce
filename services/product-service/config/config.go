package config

import (
	"os"
	"time"

	"gorm.io/gorm"
)

var (
	logger       *Logger
	DB           *gorm.DB
	configLogger = GetLogger("CONFIG")
	systemlogger = GetLogger("SYSTEM")

	startTime = time.Now()
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
	ShowWelcomeBanner()

	systemlogger.System("🔧 ETAPA 1: CONFIGURAÇÕES INICIAIS")
	LoadEnv("../..")

	systemlogger.System("🔌 ETAPA 2: CONEXÕES EXTERNAS")
	systemlogger.System("🛢️ Conectando ao banco de dados...")
	ConnectDB()
	checkDatabase()

	systemlogger.System("🚀 ETAPA 3: INICIALIZANDO SERVIÇOS")
	systemlogger.System("🌐 Configurando rotas HTTP...\n")

	systemlogger.Systemf("✅ SISTEMA PRONTO | Porta: :%s | Tempo: %v", os.Getenv("PRODUCT_SERVICE_PORT"), time.Since(startTime).Round(time.Millisecond))
	systemlogger.Systemf("🕒 Iniciado em: %s\n\n", startTime.Format("02/01/2006 15:04:05"))
}
