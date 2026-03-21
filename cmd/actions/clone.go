package actions

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/atticus64/dona/cmd/util"
	"github.com/spf13/cobra"
)

func cloneRepo(url string) error {

	home, fsErr := util.GetHome()
	if fsErr != nil {
		return fsErr
	}
	if e := exec.Command("rm", "-rf", home+"/.dona/dotfiles").Run(); e != nil {
		fmt.Println(e)
		return e
	}

	out, err := exec.Command("git", "clone", url, home+"/.dona/dotfiles").Output()

	if err != nil {
		return err
	}

	fmt.Print(string(out))

	return nil

}

var CloneCmd = &cobra.Command{
	Use:     "clone [Git Repo URL]",
	Short:   "Clone dotfiles from Repository",
	Args:    cobra.MinimumNArgs(1),
	Example: "dona clone https://github.com/user/dotfiles",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cloneRepo(args[0]); err != nil {
			fmt.Println("Failed to clone")
			os.Exit(1)
		}
		fmt.Println("Dotfiles cloned successfully")
	},
}
