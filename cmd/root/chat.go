package root

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/solumD/chat-client/internal/app"
	"github.com/solumD/chat-client/internal/model"
	"github.com/spf13/cobra"
)

const (
	createChatEP  = "/chat_v1.ChatV1/CreateChat"
	connectChatEP = "/chat_v1.ChatV1/ConnectChat"
)

// команда создания чата
var createChatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Создает чат",
	Run: func(cmd *cobra.Command, _ []string) {
		ctx := cmd.Context()

		chatName, err := cmd.Flags().GetString("chat-name")
		if err != nil {
			log.Fatalf("failed to get chat name: %s\n", err.Error())
		}

		usernamesStr, err := cmd.Flags().GetString("usernames")
		if err != nil {
			log.Fatalf("failed to get usernames: %s\n", err.Error())
		}

		usernames := strings.Split(usernamesStr, " ")
		if len(usernames) == 0 {
			fmt.Println("Список имен пользователей не может быть пустым")
			return
		}

		aToken, err := cmd.Flags().GetString("access-token")
		if err != nil {
			log.Fatalf("failed to get access token: %s\n", err.Error())
		}

		a, err := app.NewApp(ctx)
		if err != nil {
			log.Fatalf("failed to get new app: %s\n", err.Error())
		}

		_, err = a.ServiceProvider.AuthServerClient(ctx).Check(ctx, aToken, createChatEP)
		if err != nil {
			fmt.Printf("\nОшибка во время проверки доступа\n%s", err.Error())
			return
		}

		chatID, err := a.ServiceProvider.ChatServerClient(ctx).CreateChat(ctx,
			&model.Chat{
				Name:      chatName,
				Usernames: usernames,
			},
		)

		if err != nil {
			fmt.Printf("\nНе удалось создать чат\n%s\n", err.Error())
			return
		}

		fmt.Printf("\nУспешно создан чат с id %d\n", chatID)
	},
}

// команда подключения чата и отправки в него сообщения
var connectChatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Подключает к чату дает возможность писать в него сообщения. Нажмите q/й для выхода",
	Run: func(cmd *cobra.Command, _ []string) {
		ctx := cmd.Context()

		id, err := cmd.Flags().GetString("chat-id")
		if err != nil {
			log.Fatalf("failed to get chat id: %s\n", err.Error())
		}

		cID, err := strconv.Atoi(id)
		if err != nil {
			log.Fatalf("invalid chat id: %s\n", err.Error())
		}
		chatID := int64(cID)

		aToken, err := cmd.Flags().GetString("access-token")
		if err != nil {
			log.Fatalf("failed to get access token: %s\n", err.Error())
		}

		a, err := app.NewApp(ctx)
		if err != nil {
			log.Fatalf("failed to get new app: %s\n", err.Error())
		}

		username, err := a.ServiceProvider.AuthServerClient(ctx).Check(ctx, aToken, connectChatEP)

		if err != nil {
			fmt.Printf("\nОшибка во время проверки доступа\n%s\n", err.Error())
			return
		}

		stream, err := a.ServiceProvider.ChatServerClient(ctx).ConnectChat(ctx, chatID, username)
		if err != nil {
			fmt.Printf("\nНе удалось подключиться к чату\n%s", err.Error())
			return
		}

		fmt.Println("\nУспешное подключение к чату")
		fmt.Println("Для отправки сообщения введите его текст, а затем нажмите два раза на Enter")

		// читаем входящие сообщения
		go func() {
			for {
				message, errRecv := stream.Recv()
				if errRecv == io.EOF {
					return
				}
				if errRecv != nil {
					log.Printf("Не удалось получить сообщения из чата\n%s\n", errRecv.Error())
					return
				}

				log.Printf("[from: %s]: %s\n", message.GetFrom(), message.GetText())
			}
		}()

		// считываем сообщение
		for {
			scanner := bufio.NewScanner(os.Stdin)
			var lines strings.Builder
			for {
				scanner.Scan()
				line := scanner.Text()
				if len(line) == 0 {
					break
				}

				if line == "q" || line == "й" {
					fmt.Println("Выход из чата")
					return
				}

				lines.WriteString(line)
				lines.WriteString("\n")
			}

			err = scanner.Err()
			if err != nil {
				log.Printf("Не удалось просканировать сообщение\n%s", err.Error())
				return
			}

			// отправляем сообщение
			err = a.ServiceProvider.ChatServerClient(ctx).SendMessage(ctx, &model.Message{
				ChatID: chatID,
				From:   username,
				Text:   lines.String(),
			})

			if err != nil {
				log.Printf("Не удалось отправить сообщение\n%s", err.Error())
				return
			}
		}
	},
}

func init() {
	// инициализируем команду создания чата
	createCmd.AddCommand(createChatCmd)

	createChatCmd.Flags().StringP("access-token", "t", "", "Access token")
	err := createChatCmd.MarkFlagRequired("access-token")
	if err != nil {
		log.Fatalf("failed to mark access-token flag as required: %s\n", err.Error())
	}

	createChatCmd.Flags().StringP("chat-name", "n", "", "Имя чата")
	err = createChatCmd.MarkFlagRequired("chat-name")
	if err != nil {
		log.Fatalf("failed to mark chat-name flag as required: %s\n", err.Error())
	}

	createChatCmd.Flags().StringP("usernames", "u", "", "Пользователи чата")
	err = createChatCmd.MarkFlagRequired("usernames")
	if err != nil {
		log.Fatalf("failed to mark usernames flag as required: %s\n", err.Error())
	}

	// инициализируем команду подключения к чату и отправки в него сообщения
	connectCmd.AddCommand(connectChatCmd)
	connectChatCmd.Flags().StringP("chat-id", "i", "", "ID чата")
	err = connectChatCmd.MarkFlagRequired("chat-id")
	if err != nil {
		log.Fatalf("failed to mark chat-id flag as required: %s\n", err.Error())
	}

	connectChatCmd.Flags().StringP("access-token", "t", "", "Access token")
	err = connectChatCmd.MarkFlagRequired("access-token")
	if err != nil {
		log.Fatalf("failed to mark access-token flag as required: %s\n", err.Error())
	}
}
