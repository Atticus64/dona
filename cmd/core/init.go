package core

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/atticus64/dona/cmd/models"
	"github.com/atticus64/dona/cmd/util"
	"github.com/spf13/cobra"
)

func CreateDotsFolders() error {
	dirname, err := util.GetHome()
	if err != nil {
		return err
	}

	path := dirname + "/.dona"
	util.CheckDir(path)
	util.CheckDir(path + "/dots")
	util.CheckDir(path + "/dotfiles")

	data := []models.Pin{}

	files, err := os.ReadDir(path)

	file, _ := json.MarshalIndent(data, "", " ")
	os.WriteFile(path+"/pins.json", file, 0644)

	if err != nil {
		return err
	}

	if len(files) <= 2 {
		return fmt.Errorf("No files found in %s", path)
	}

	return nil
}

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize your dotfiles",
	Run: func(cmd *cobra.Command, args []string) {
		CreateDotsFolders()
		fmt.Println("Dotfiles initialized")
	},
}
