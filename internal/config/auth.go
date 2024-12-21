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
)

type authServerConfig struct {
	authHost     string
	authPort     string
	authCertPath string
}

func NewAuthServerConfig() (AuthServerConfig, error) {
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

	return &authServerConfig{
		authHost:     authHost,
		authPort:     authPort,
		authCertPath: authCertPath,
	}, nil
}

func (cfg *authServerConfig) AuthServerAddress() string {
	return net.JoinHostPort(cfg.authHost, cfg.authPort)
}

func (cfg *authServerConfig) AuthCertPath() string {
	return cfg.authCertPath
}
