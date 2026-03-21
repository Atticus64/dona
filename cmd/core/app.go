package core

import (
	"fmt"

	"github.com/atticus64/dona/cmd/actions"
	"github.com/atticus64/dona/cmd/util"
	"github.com/spf13/cobra"
)

var ShowVersion bool
var page int

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
			fmt.Println("Dona version " + util.GetVersion())
			return
		}

		fmt.Println("Dona -> Dotfiles natural manager")
		cmd.Help()
	},
}

func CreateDona() error {
	configure()
	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}

func configure() *cobra.Command {

	rootCmd.PersistentFlags().BoolVarP(&ShowVersion, "version", "v", false, "Show Dona version")

	actions.SearchCmd.PersistentFlags().IntVarP(&page, "page", "p", 1, "Number of page")
	actions.PinCmd.PersistentFlags().StringVarP(&actions.Tag, "tag", "t", "", "Tag to save pin")
	actions.GitCmd.DisableFlagParsing = true
	rootCmd.AddCommand(util.VersionCmd, actions.SearchCmd, actions.GitCmd, actions.DelCmd)
	rootCmd.AddCommand(InitCmd, actions.CloneCmd, actions.SaveCmd, actions.ListCmd)
	rootCmd.AddCommand(actions.PinCmd)

	return rootCmd
}
