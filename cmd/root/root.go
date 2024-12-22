package root

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chat-client",
	Short: "Клиент чата",
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Создает что-то",
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Подключается к чему-то",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(connectCmd)

	rootCmd.AddCommand(createUserCmd)
	rootCmd.AddCommand(loginCmd)

	rootCmd.AddCommand(createChatCmd)
}
