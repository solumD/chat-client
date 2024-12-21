package config

import (
	"errors"
	"net"
	"os"
)

const (
	chatHostEnvName     = "CHAT_SERVER_HOST"
	chatPortEnvName     = "CHAT_SERVER_PORT"
	chatCertPathEnvName = "CHAT_CERT_PATH"
)

type chatServerConfig struct {
	chatHost     string
	chatPort     string
	chatCertPath string
}

func NewChatServerConfig() (ChatServerConfig, error) {
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

	return &chatServerConfig{
		chatHost:     chatHost,
		chatPort:     chatPort,
		chatCertPath: chatCertPath,
	}, nil
}

func (cfg *chatServerConfig) ChatServerAddress() string {
	return net.JoinHostPort(cfg.chatHost, cfg.chatPort)
}

func (cfg *chatServerConfig) ChatCertPath() string {
	return cfg.chatCertPath
}
