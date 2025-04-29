package utils

import (
	"errors"
	"regexp"
	"strings"

	"github.com/vidinine-ecommerce/aut-service/config"
)

var (
	utilsLogger = config.GetLogger("Validators")
)

func ValidateEmail(email string) error {
	if strings.TrimSpace(email) == "" {
		utilsLogger.Error("Email vazio")
		return errors.New("Email é obrigatório")
	}

	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, email)

	if !matched {
		utilsLogger.Errorf("Email inválido: %s", email)
		return errors.New("Email inválido")
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 6 {
		utilsLogger.Error("Senha muito curta")
		return errors.New("senha deve ter no mínimo 6 caracteres")
	}

	return nil
}
