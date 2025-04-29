package config

import (
	"io"
	"log"
	"os"
)

const (
	colorReset = "\033[0m"
	colorGreen = "\033[32m"
	colorRed   = "\033[31m"
)

type Logger struct {
	errorLogger   *log.Logger // ❌ Error logger
	successLogger *log.Logger // ✅ Success logger
	neutralLogger *log.Logger // ⚪ Neutral logger
	writer        io.Writer   // Writer para o log
}

// NewLogger cria uma nova instância de Logger com o prefixo fornecido `p`.
func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		successLogger: log.New(writer, colorGreen+"["+p+"] "+"SUCCESS: "+colorReset, logger.Flags()), // ✅
		errorLogger:   log.New(writer, colorRed+"["+p+"] "+"ERROR: "+colorReset, logger.Flags()),     // ❌
		neutralLogger: log.New(writer, "["+p+"] ", logger.Flags()),                                   // ⚪
		writer:        writer,
	}
}

// Success registra uma mensagem de sucesso ✅.
func (l *Logger) Success(v ...interface{}) {
	l.successLogger.Println(v...)
}

// Error registra uma mensagem de erro ❌.
func (l *Logger) Error(v ...interface{}) {
	l.errorLogger.Println(v...)
}

// Log registra uma mensagem neutra ⚪.
func (l *Logger) Log(v ...interface{}) {
	l.neutralLogger.Println(v...)
}

// Successf registra uma mensagem de sucesso formatada ✅.
func (l *Logger) Successf(format string, v ...interface{}) {
	l.successLogger.Printf(format, v...)
}

// Errorf registra uma mensagem de erro formatada ❌.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.errorLogger.Printf(format, v...)
}

// Logf registra uma mensagem neutra formatada ⚪.
func (l *Logger) Logf(format string, v ...interface{}) {
	l.neutralLogger.Printf(format, v...)
}
