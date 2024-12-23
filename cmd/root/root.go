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

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Получает что-то",
}

// Execute выполняет переданную команду
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// инициализируем команды создания
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(createChatCmd)
	rootCmd.AddCommand(createUserCmd)

	// инициализируем команду подключения
	rootCmd.AddCommand(connectCmd)

	// инициализируем команду получения
	rootCmd.AddCommand(getCmd)
}
