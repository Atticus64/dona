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
		ShowVersion, error := cmd.Flags().GetBool("version")
		if error != nil {
			fmt.Println(error)
		}

		if ShowVersion {
			fmt.Println("Dona version 0.0.1 beta ultra alpha radioactive")
			return
		}

		fmt.Println("Dona cli to manage your dotfiles")
	},
}

var ShowVersion bool
var page int

func main() {
	rootCmd.PersistentFlags().BoolVarP(&ShowVersion, "version", "v", false, "Show Dona version")

	cmd.SearchCmd.PersistentFlags().IntVarP(&page, "page", "p", 1, "Number of page")
	rootCmd.AddCommand(cmd.VersionCmd)
	rootCmd.AddCommand(cmd.SearchCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
