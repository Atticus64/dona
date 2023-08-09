package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"strings"
	"os/exec"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
)


func isUrl(url string) bool {
	if strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "git@") {
		return true
	}
	return false
}

func getName(url string) string {
	
	if strings.HasPrefix(url, "git@") {
		data := strings.Split(url, ":")
		name := strings.Split(data[1], ".")[0]
		name = strings.Join(strings.Split(name, "/"), "-")
		return name 
	} else {
		data := strings.Split(url, ".com/")
		name := data[1]
		name = strings.Join(strings.Split(name, "/"), "-")
		return name 
	}
}

var SaveCmd = &cobra.Command{
	Use:   "save [Git Repo URL/Name]",
	Short: "Save third party dotfiles in dots directory",
	Args:  cobra.MinimumNArgs(1),
	Example: `
	dona save user/dotfiles # uri postfix
	dona save https://github.com/user/dots  
	`,
	Run: func(cmd *cobra.Command, args []string) {
		param := args[0]
		w := wow.New(os.Stdout, spin.Get(spin.Earth), " Searching in github")
		w.Start()

		home, fsErr := GetHome()

		if fsErr != nil {
			fmt.Println(fsErr)
			return
		}

		if isUrl(param) {
			name := getName(param)
			fmt.Println(name)
			err := exec.Command("git", "clone", param, home + "/.dona/dots/" + name).Run()
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			url := fmt.Sprintf("https://github.com/%s", param)
			name := getName(url)
			err := exec.Command("git", "clone", url, home + "/.dona/dots/" + name).Run()

			if err != nil {
				fmt.Println(err)
				return
			}
		}
		
		w.PersistWith(spin.Spinner{Frames: []string{"☕"}}, " Enjoy!")

	},
}


