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

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
