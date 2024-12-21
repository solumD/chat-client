package client

import (
	"context"

	"github.com/solumD/chat-client/internal/model"
)

type AuthClient interface {
	CreateUser(ctx context.Context, user *model.UserToCreate) error
	Login(ctx context.Context, user *model.UserToLogin) error
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
	Check(ctx context.Context, accessToken string) (string, error)
}

type ChatClient interface {
	CreateChat(ctx context.Context, chat *model.Chat) (int64, error)
	ConnectChat(ctx context.Context, chatID int64, username string) error
	SendMessage(ctx context.Context, chatID int64, from string, text string) error
}
