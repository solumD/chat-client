package config

import (
	"errors"
	"net"
	"os"
)

const (
	authHostEnvName     = "AUTH_SERVER_HOST"
	authPortEnvName     = "AUTH_SERVER_PORT"
	authCertPathEnvName = "AUTH_CERT_PATH"
	chatHostEnvName     = "CHAT_SERVER_HOST"
	chatPortEnvName     = "CHAT_SERVER_PORT"
	chatCertPathEnvName = "CHAT_CERT_PATH"
)

type chatClientConfig struct {
	authHost     string
	authPort     string
	authCertPath string
	chatHost     string
	chatPort     string
	chatCertPath string
}

func NewChatClientConfig() (ClientConfig, error) {
	authHost := os.Getenv(authHostEnvName)
	if len(authHost) == 0 {
		return nil, errors.New("auth server host not found")
	}

	authPort := os.Getenv(authPortEnvName)
	if len(authPort) == 0 {
		return nil, errors.New("auth server port not found")
	}

	authCertPath := os.Getenv(authCertPathEnvName)
	if len(authCertPath) == 0 {
		return nil, errors.New("auth cert path not found")
	}

	chatHost := os.Getenv(chatHostEnvName)
	if len(chatHost) == 0 {
		return nil, errors.New("chat server host not found")
	}

	chatPort := os.Getenv(chatPortEnvName)
	if len(chatPort) == 0 {
		return nil, errors.New("chat server port not found")
	}

	chatCertPath := os.Getenv(chatCertPathEnvName)
	if len(chatCertPath) == 0 {
		return nil, errors.New("chat cert path not found")
	}

	return &chatClientConfig{
		authHost:     authHost,
		authPort:     authPort,
		authCertPath: authCertPath,
		chatHost:     chatHost,
		chatPort:     chatPort,
		chatCertPath: chatCertPath,
	}, nil
}

func (cfg *chatClientConfig) AuthServerAddress() string {
	return net.JoinHostPort(cfg.authHost, cfg.authPort)
}

func (cfg *chatClientConfig) ChatServerAddress() string {
	return net.JoinHostPort(cfg.chatHost, cfg.chatPort)
}

func (cfg *chatClientConfig) AuthCertPath() string {
	return cfg.authCertPath
}

func (cfg *chatClientConfig) ChatCertPath() string {
	return cfg.chatCertPath
}
