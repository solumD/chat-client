package root

import (
	"fmt"
	"log"

	"github.com/solumD/chat-client/internal/app"
	"github.com/solumD/chat-client/internal/model"
	"github.com/spf13/cobra"
)

var createUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Создает пользователя",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		username, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatalf("failed to get username: %v", err)
		}

		email, err := cmd.Flags().GetString("email")
		if err != nil {
			log.Fatalf("failed to get email: %v", err)
		}

		password, err := cmd.Flags().GetString("password")
		if err != nil {
			log.Fatalf("failed to get password: %v", err)
		}

		a, err := app.NewApp(ctx)
		if err != nil {
			log.Fatalf("failed to get new app: %v", err)
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
			fmt.Printf("Не удалось создать пользователя\n%v\n", err)
			return
		}

		fmt.Printf("Успешно создан пользователь с id %d\n", userID)
	},
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Авторизует пользователя на сервере",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		username, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatalf("failed to get username: %v", err)
		}

		password, err := cmd.Flags().GetString("password")
		if err != nil {
			log.Fatalf("failed to get password: %v", err)
		}

		a, err := app.NewApp(ctx)
		if err != nil {
			log.Fatalf("failed to get new app: %v", err)
		}

		rToken, aToken, err := a.ServiceProvider.AuthServerClient(ctx).Login(ctx,
			&model.UserToLogin{
				Name:     username,
				Password: password,
			},
		)

		if err != nil {
			fmt.Printf("Не удалось выполниить вход\n%v\n", err)
			return
		}

		fmt.Printf("Успешный вход\nrefresh_token: %s\naccess_token: %s\n", rToken, aToken)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(createUserCmd)
	rootCmd.AddCommand(loginCmd)

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
}