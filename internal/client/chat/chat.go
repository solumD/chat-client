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

// New возвращает новый объект chatServerClient
func New(chatClient chat_v1.ChatV1Client) client.ChatServerClient {
	return &chatServerClient{
		chatClient: chatClient,
	}
}

// CreateChat отправляет запрос на создание чата и
// в случае успеха возвращает его id
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

// ConnectChat отправляет запрос на подключение к чату и
// в случае успеха возвращает stream чата
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

// GetUserChats отправляет запрос на получение всех чатов пользователя и информации о них
func (cl *chatServerClient) GetUserChats(ctx context.Context, username string) ([]*chat_v1.ChatInfo, error) {
	req := &chat_v1.GetUserChatsRequest{
		Username: username,
	}

	res, err := cl.chatClient.GetUserChats(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.GetChats(), nil
}

// SendMessage отправляет запрос на отправку сообщения в чат
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
