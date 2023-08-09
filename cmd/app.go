package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

func GetHome () (string, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return dirname, nil
}

func checkDir (dirname string) error {
	if _, err := os.Stat(dirname); err != nil {
		if err := os.Mkdir(dirname, 0755); err != nil {
			panic(err)
		}
	} 
	
	return nil
}

var ShowVersion bool
var page int
var tag string

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

func CreateDona() error {
	configure()
	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}

func configure() *cobra.Command {

	rootCmd.PersistentFlags().BoolVarP(&ShowVersion, "version", "v", false, "Show Dona version")

	SearchCmd.PersistentFlags().IntVarP(&page, "page", "p", 1, "Number of page")
	PinCmd.PersistentFlags().StringVarP(&tag, "tag", "t", "", "Tag to save pin")
	GitCmd.DisableFlagParsing = true
	rootCmd.AddCommand(VersionCmd, SearchCmd, GitCmd, DelCmd)
	rootCmd.AddCommand(InitCmd, CloneCmd, SaveCmd, ListCmd)
	rootCmd.AddCommand(PinCmd)

	return rootCmd
}


