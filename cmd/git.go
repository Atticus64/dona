package cmd

import (
	"os/exec"
	"fmt"
	"github.com/spf13/cobra"
)

var GitCmd = &cobra.Command{
	Use:   "git [args]",
	Short: "Execute Git commands in your dotfiles",
	Example: `
	dona git status -s 
	dona git revert HEAD
	dona git remote add origin https://github.com/user/dots
	`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.DisableFlagParsing = true
		runGit(args)
	},
}

func GetDotfilesPath()	(string, error) {
	dirname, err := GetHome()
	if err != nil {
		return "", err
	}
	return dirname + "/.dona/dotfiles", nil
}


func runGit(args []string) {

	path, e := GetDotfilesPath()
	commandArgs := []string{"-C", path}
	if e != nil {
		panic(e)
	}

	args = append(commandArgs, args...)

    out, err := exec.Command("git", args...).Output()	

	if err != nil {
		fmt.Println("Git command failed")
		return
	}

	fmt.Print(string(out))
}
