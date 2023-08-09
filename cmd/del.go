package cmd

import (
	"fmt"
	"os/exec"
	"os"
	"strings"
	"github.com/spf13/cobra"
)

var DelCmd = &cobra.Command{
	Use:   "del [Git Repo URL/Name]",
	Short: "Delete a third party dotfile in dots directory",
	Args:  cobra.MinimumNArgs(1),
	Example: `
	dona del user/dotfiles # uri postfix
	dona del user-dotfiles # folder name
	`,
	Run: func(cmd *cobra.Command, args []string) {
		param := args[0]

		home, fsErr := GetHome()

		if fsErr != nil {
			fmt.Println(fsErr)
			return
		}

		name := strings.Join(strings.Split(param, "/"), "-")
		path := home + "/.dona/dots/" + name
		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Println("🤷 No dotfiles found")
			return
		}
		err := exec.Command("rm", "-rf", path).Run()

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("🧨 Deleted!")

	},
}



