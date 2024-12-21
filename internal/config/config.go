package config

import "github.com/joho/godotenv"

type ClientConfig interface {
	AuthServerAddress() string
	ChatServerAddress() string
	AuthCertPath() string
	ChatCertPath() string
}

// LoggerConfig интерфейс конфига логгера
type LoggerConfig interface {
	Level() string
}

// Load читает .env файл по указанному пути
// и загружает переменные в проект
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
