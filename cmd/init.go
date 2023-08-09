package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"os"
	"fmt"
)

type Pin struct {
	Name string
	Tag string
}
func CreateDotsFolders () error {
	dirname, err := GetHome()
	if err != nil {
		return err
	}

	path := dirname + "/.dona"
	checkDir(path)
	checkDir(path + "/dots") 
    checkDir(path + "/dotfiles") 

	data := []Pin{}

	files, err := os.ReadDir(path)

	file, _ := json.MarshalIndent(data, "", " ")
	os.WriteFile(path + "/pins.json", file, 0644)

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


