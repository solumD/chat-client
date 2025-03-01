package root

import (
	"fmt"
	"log"

	"github.com/solumD/chat-client/internal/app"
	"github.com/solumD/chat-client/internal/model"
	"github.com/spf13/cobra"
)

// команда создания юзера
var createUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Создает пользователя",
	Run: func(cmd *cobra.Command, _ []string) {
		ctx := cmd.Context()

		username, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatalf("failed to get username: %s\n", err.Error())
		}

		email, err := cmd.Flags().GetString("email")
		if err != nil {
			log.Fatalf("failed to get email: %s\n", err.Error())
		}

		password, err := cmd.Flags().GetString("password")
		if err != nil {
			log.Fatalf("failed to get password: %s\n", err.Error())
		}

		a, err := app.NewApp(ctx)
		if err != nil {
			log.Fatalf("failed to get new app: %s\n", err.Error())
		}

		userID, err := a.ServiceProvider.AuthServerClient(ctx).CreateUser(ctx,
			&model.UserToCreate{
				Name:            username,
				Email:           email,
				Password:        password,
				PasswordConfirm: password,
			},
		)

		if err != nil {
			fmt.Printf("\nНе удалось создать пользователя\n%s\n", err.Error())
			return
		}

		fmt.Printf("\nУспешно создан пользователь с id %d\n", userID)
	},
}

// команда авторизации юзера
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Авторизует пользователя на сервере",
	Run: func(cmd *cobra.Command, _ []string) {
		ctx := cmd.Context()

		username, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatalf("failed to get username: %s\n", err.Error())
		}

		password, err := cmd.Flags().GetString("password")
		if err != nil {
			log.Fatalf("failed to get password: %s\n", err.Error())
		}

		a, err := app.NewApp(ctx)
		if err != nil {
			log.Fatalf("failed to get new app: %s\n", err.Error())
		}

		rToken, aToken, err := a.ServiceProvider.AuthServerClient(ctx).Login(ctx,
			&model.UserToLogin{
				Name:     username,
				Password: password,
			},
		)

		if err != nil {
			fmt.Printf("\nНе удалось выполниить вход\n%v\n", err.Error())
			return
		}

		chats, err := a.ServiceProvider.ChatServerClient(ctx).GetUserChats(ctx, username)
		if err != nil {
			fmt.Printf("\nНе удалось получить доступные чаты\n%v\n", err.Error())
			return
		}

		fmt.Printf("Добро пожаловать, %s!\n", username)
		fmt.Printf("\nВаши токены доступа:\nrefresh_token: %s\naccess_token: %s\n", rToken, aToken)
		fmt.Println("\nВаши чаты:")
		for _, c := range chats {
			fmt.Printf("ID: %d | Название: %s | Пользователи: %v\n", c.GetId(), c.GetName(), c.GetUsernames())
		}
		fmt.Println()
	},
}

// команда получения refresh токена
var getRefreshTokenCmd = &cobra.Command{
	Use:   "refresh-token",
	Short: "Получает новый refresh токен после отправки старого",
	Run: func(cmd *cobra.Command, _ []string) {
		ctx := cmd.Context()

		refreshToken, err := cmd.Flags().GetString("refresh-token")
		if err != nil {
			log.Fatalf("failed to get refresh token: %s\n", err.Error())
		}

		a, err := app.NewApp(ctx)
		if err != nil {
			log.Fatalf("failed to get new app: %s\n", err.Error())
		}

		newRefreshToken, err := a.ServiceProvider.AuthServerClient(ctx).GetRefreshToken(ctx, refreshToken)
		if err != nil {
			fmt.Printf("\nНе удалось получить новый refresh токен\n%s\n", err.Error())
			return
		}

		fmt.Printf("\nУспешное получение нового refresh токена\n%s\n", newRefreshToken)
	},
}

// команда получения access токена
var getAccessTokenCmd = &cobra.Command{
	Use:   "access-token",
	Short: "Получает новый access токен после отправки refresh токена",
	Run: func(cmd *cobra.Command, _ []string) {
		ctx := cmd.Context()

		refreshToken, err := cmd.Flags().GetString("refresh-token")
		if err != nil {
			log.Fatalf("failed to get refresh token: %s\n", err.Error())
		}

		a, err := app.NewApp(ctx)
		if err != nil {
			log.Fatalf("failed to get new app: %s\n", err.Error())
		}

		newAccessToken, err := a.ServiceProvider.AuthServerClient(ctx).GetAccessToken(ctx, refreshToken)
		if err != nil {
			fmt.Printf("\nНе удалось получить новый access токен\n%s\n", err.Error())
			return
		}

		fmt.Printf("\nУспешное получение нового access токена\n%s\n", newAccessToken)
	},
}

func init() {
	// инициализируем команду создания юзера
	createCmd.AddCommand(createUserCmd)

	createUserCmd.Flags().StringP("username", "u", "", "Имя пользователя")
	err := createUserCmd.MarkFlagRequired("username")
	if err != nil {
		log.Fatalf("failed to mark username flag as required: %s\n", err.Error())
	}

	createUserCmd.Flags().StringP("email", "e", "", "Почта")
	err = createUserCmd.MarkFlagRequired("email")
	if err != nil {
		log.Fatalf("failed to mark email flag as required: %s\n", err.Error())
	}

	createUserCmd.Flags().StringP("password", "p", "", "Пароль")
	err = createUserCmd.MarkFlagRequired("password")
	if err != nil {
		log.Fatalf("failed to mark password flag as required: %s\n", err.Error())
	}

	// инициализируем команду авторизации
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("username", "u", "", "Имя пользователя")
	err = loginCmd.MarkFlagRequired("username")
	if err != nil {
		log.Fatalf("failed to mark username flag as required: %s\n", err.Error())
	}

	loginCmd.Flags().StringP("password", "p", "", "Пароль")
	err = loginCmd.MarkFlagRequired("password")
	if err != nil {
		log.Fatalf("failed to mark password flag as required: %s\n", err.Error())
	}

	// инициализируем команду получения refresh токена
	getCmd.AddCommand(getRefreshTokenCmd)
	getRefreshTokenCmd.Flags().StringP("refresh-token", "t", "", "Refresh токен")
	err = getRefreshTokenCmd.MarkFlagRequired("refresh-token")
	if err != nil {
		log.Fatalf("failed to mark refresh-token flag as required: %s\n", err.Error())
	}

	// инициализируем команду получения access токена
	getCmd.AddCommand(getAccessTokenCmd)
	getAccessTokenCmd.Flags().StringP("refresh-token", "t", "", "Refresh токен")
	err = getAccessTokenCmd.MarkFlagRequired("refresh-token")
	if err != nil {
		log.Fatalf("failed to mark refresh-token flag as required: %s\n", err.Error())
	}
}
