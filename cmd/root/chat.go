package root

import (
	"fmt"
	"log"
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

		usernames := strings.Split(usernamesStr, ",")
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

func init() {
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
}
