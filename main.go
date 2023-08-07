package main

import (
	"fmt"
	"os"

	"github.com/atticus64/dona/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dona",
	Short: "Dona 🍩 is a very fast dotfile manager",
	Long:  `CLI to manage your dotfiles`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Dona cli to manage your dotfiles")
	},
}

func main() {
	rootCmd.AddCommand(cmd.VersionCmd)
	rootCmd.AddCommand(cmd.SearchCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
