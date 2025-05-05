package config

import (
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv(baseDir string) error {
	configLogger := GetLogger("ENV")

	absPath, err := filepath.Abs(baseDir)
	if err != nil {
		configLogger.Errorf("Erro ao resolver caminho: %v", err)
		return nil
	}

	envPath := filepath.Join(absPath, ".env")
	configLogger.Logf("Buscando .env em: %s", envPath)

	if err := godotenv.Load(envPath); err != nil {
		configLogger.Errorf("Falha ao carregar .env: %v", err)
		configLogger.Errorf("arquivo .env não encontrado em %s", envPath)
		return nil
	}

	configLogger.Success("Variáveis de ambiente carregadas\n")
	return nil
}
