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

var createChatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Создает чат",
	Run: func(cmd *cobra.Command, args []string) {
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

		_, err = a.ServiceProvider.AuthServerClient(ctx).Check(ctx, aToken, "/chat_v1.ChatV1/CreateChat")
		if err != nil {
			fmt.Printf("Ошибка во время проверки доступа\n%s", err.Error())
			return
		}

		chatID, err := a.ServiceProvider.ChatServerClient(ctx).CreateChat(ctx,
			&model.Chat{
				Name:      chatName,
				Usernames: usernames,
			},
		)

		if err != nil {
			fmt.Printf("Не удалось создать чат\n%s\n", err.Error())
			return
		}

		fmt.Printf("Успешно создан чат с id %d\n", chatID)
	},
}

var connectChatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Подключает к чату. Нажмите q/й для выхода",
	Run: func(cmd *cobra.Command, args []string) {
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

		username, err := a.ServiceProvider.AuthServerClient(ctx).Check(ctx, aToken, "/chat_v1.ChatV1/ConnectChat")

		if err != nil {
			fmt.Printf("Ошибка во время проверки доступа\n%s\n", err.Error())
			return
		}

		stream, err := a.ServiceProvider.ChatServerClient(ctx).ConnectChat(ctx, chatID, username)
		if err != nil {
			fmt.Printf("Не удалось подключиться к чату\n%s", err.Error())
			return
		}

		fmt.Println("Успешное подключение к чату")
		fmt.Println("Для отправки сообщения введите его текст, а затем нажмите два раза на Enter")

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
	createCmd.AddCommand(createChatCmd)
	connectCmd.AddCommand(connectChatCmd)

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
