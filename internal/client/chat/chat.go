package chat

import (
	"context"

	"github.com/solumD/chat-client/internal/client"
	"github.com/solumD/chat-client/internal/model"
	"github.com/solumD/chat-server/pkg/chat_v1"
)

type chatServerClient struct {
	chatClient chat_v1.ChatV1Client
}

func New(chatClient chat_v1.ChatV1Client) client.ChatServerClient {
	return &chatServerClient{
		chatClient: chatClient,
	}
}

func (cl *chatServerClient) CreateChat(ctx context.Context, chat *model.Chat) (int64, error) {
	req := &chat_v1.CreateChatRequest{
		Name:      chat.Name,
		Usernames: chat.Usernames,
	}

	res, err := cl.chatClient.CreateChat(ctx, req)
	if err != nil {
		return 0, err
	}

	return res.GetId(), nil
}

func (cl *chatServerClient) ConnectChat(ctx context.Context, chatID int64, username string) (chat_v1.ChatV1_ConnectChatClient, error) {
	req := &chat_v1.ConnectChatRequest{
		Id:       chatID,
		Username: username,
	}

	stream, err := cl.chatClient.ConnectChat(ctx, req)
	if err != nil {
		return nil, err
	}

	return stream, nil
}

func (cl *chatServerClient) SendMessage(ctx context.Context, message *model.Message) error {
	req := &chat_v1.SendMessageRequest{
		Id:   message.ChatID,
		From: message.From,
		Text: message.Text,
	}

	_, err := cl.chatClient.SendMessage(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
