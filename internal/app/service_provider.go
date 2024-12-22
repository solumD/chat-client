package app

import (
	"context"
	"log"

	"github.com/solumD/auth/pkg/access_v1"
	"github.com/solumD/auth/pkg/auth_v1"
	"github.com/solumD/auth/pkg/user_v1"
	"github.com/solumD/chat-client/internal/client"
	"github.com/solumD/chat-client/internal/client/auth"
	"github.com/solumD/chat-client/internal/client/chat"
	"github.com/solumD/chat-client/internal/closer"
	"github.com/solumD/chat-client/internal/config"
	"github.com/solumD/chat-server/pkg/chat_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type serviceProvider struct {
	authServerCfg config.AuthServerConfig
	chatServerCfg config.ChatServerConfig

	authServerClient client.AuthServerClient
	chatServerClient client.ChatServerClient
}

func NewServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) AuthServerConfig() config.AuthServerConfig {
	if s.authServerCfg == nil {
		cfg, err := config.NewAuthServerConfig()
		if err != nil {
			log.Fatalf("failed to get auth server config: %v", err)
		}

		s.authServerCfg = cfg
	}

	return s.authServerCfg
}

func (s *serviceProvider) ChatServerConfig() config.ChatServerConfig {
	if s.chatServerCfg == nil {
		cfg, err := config.NewChatServerConfig()
		if err != nil {
			log.Fatalf("failed to get chat server config: %v", err)
		}

		s.chatServerCfg = cfg
	}

	return s.chatServerCfg
}

func (s *serviceProvider) AuthServerClient(ctx context.Context) client.AuthServerClient {
	if s.authServerClient == nil {
		creds, err := credentials.NewClientTLSFromFile(s.AuthServerConfig().AuthCertPath(), "")
		if err != nil {
			log.Fatalf("could not auth server process credentials: %v", err)
		}

		conn, err := grpc.DialContext(ctx, s.AuthServerConfig().AuthServerAddress(), grpc.WithTransportCredentials(creds))
		if err != nil {
			log.Fatalf("failed to connect to %s: %v", s.AuthServerConfig().AuthServerAddress(), err)
		}

		closer.Add(conn.Close)

		userClient := user_v1.NewUserV1Client(conn)
		authClient := auth_v1.NewAuthV1Client(conn)
		accessClient := access_v1.NewAccessV1Client(conn)

		s.authServerClient = auth.New(userClient, authClient, accessClient)
	}

	return s.authServerClient
}

func (s *serviceProvider) ChatServerClient(ctx context.Context) client.ChatServerClient {
	if s.chatServerClient == nil {
		creds, err := credentials.NewClientTLSFromFile(s.ChatServerConfig().ChatCertPath(), "")
		if err != nil {
			log.Fatalf("could not chat server process credentials: %v", err)
		}

		conn, err := grpc.DialContext(ctx, s.ChatServerConfig().ChatServerAddress(), grpc.WithTransportCredentials(creds))
		if err != nil {
			log.Fatalf("failed to connect to %s: %v", s.ChatServerConfig().ChatServerAddress(), err)
		}

		closer.Add(conn.Close)

		chatClient := chat_v1.NewChatV1Client(conn)

		s.chatServerClient = chat.New(chatClient)
	}

	return s.chatServerClient
}
