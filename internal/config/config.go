package config

import "github.com/joho/godotenv"

// AuthServerConfig интерфейс конфига для клиента
// сервера Auth
type AuthServerConfig interface {
	AuthServerAddress() string
	AuthCertPath() string
}

// ChatServerConfig интерфейс конфига для клиента
// сервера Chat
type ChatServerConfig interface {
	ChatServerAddress() string
	ChatCertPath() string
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
