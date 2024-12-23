package client

import (
	"context"

	"github.com/solumD/chat-client/internal/model"
	"github.com/solumD/chat-server/pkg/chat_v1"
)

// AuthServerClient интерфейс клиента auth сервера
type AuthServerClient interface {
	CreateUser(ctx context.Context, user *model.UserToCreate) (int64, error)
	Login(ctx context.Context, user *model.UserToLogin) (string, string, error)
	GetRefreshToken(ctx context.Context, refreshToken string) (string, error)
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
	Check(ctx context.Context, accessToken string, endpoint string) (string, error)
}

// ChatServerClient интерфейс клиента chat сервера
type ChatServerClient interface {
	CreateChat(ctx context.Context, chat *model.Chat) (int64, error)
	ConnectChat(ctx context.Context, chatID int64, username string) (chat_v1.ChatV1_ConnectChatClient, error)
	SendMessage(ctx context.Context, message *model.Message) error
}
