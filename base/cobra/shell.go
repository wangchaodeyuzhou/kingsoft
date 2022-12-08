package cobra

import (
	"fmt"
	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "A brief description of your command",
	Long:  "A longer description that spans multiple lines and likely contains examples and usage of using your command. For example:",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("shell called")
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)
}
