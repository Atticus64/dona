package actions

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/atticus64/dona/cmd/models"
	"github.com/atticus64/dona/cmd/util"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var Tag string

func addPin(home string, args []string) error {
	path := home + "/.dona/pins.json"

	// check if exist pins.json
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println(color.RedString("Error: ") + "No pins.json found")
		fmt.Println(color.YellowString("Suggestion: ") + "Run `dona init` to create that file")
		return err
	}

	newPin := models.Pin{
		Name: args[1],
		Tag:  Tag,
	}

	file, err := os.ReadFile(home + "/.dona/pins.json")
	pins := []models.Pin{}
	if err != nil {
		fmt.Println(err)
		return err
	}

	if e := json.Unmarshal([]byte(file), &pins); e != nil {
		fmt.Println(e)
		return err
	}

	pins = append(pins, newPin)

	fileData, e := json.MarshalIndent(pins, "", " ")

	if e != nil {
		fmt.Println(e)
		return err
	}

	if err := os.WriteFile(path, fileData, 0644); err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("🎉 Pin created")

	return nil
}

func byName(pins []models.Pin, key string) bool {
	for _, pin := range pins {
		if pin.Name == key {
			return true
		}
	}

	return false

}

func delPin(home string, args []string) error {
	path := home + "/.dona/pins.json"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println(color.RedString("Error: ") + "No pins.json found")
		fmt.Println(color.YellowString("Suggestion: ") + "Run `dona init` to create that file")
		return fmt.Errorf("No pins.json file found")
	}

	key := args[1]

	file, err := os.ReadFile(home + "/.dona/pins.json")

	if err != nil {
		return err
	}

	pins := []models.Pin{}

	if e := json.Unmarshal([]byte(file), &pins); e != nil {
		return err
	}

	newPins := []models.Pin{}
	if byName(pins, key) {
		newPins = FilterPins(pins, func(pin models.Pin) bool {
			return pin.Name != key
		})

	} else {
		newPins = FilterPins(pins, func(pin models.Pin) bool {
			return pin.Tag != key
		})
	}

	fileData, e := json.MarshalIndent(newPins, "", " ")

	if e != nil {
		return err
	}

	if err := os.WriteFile(path, fileData, 0644); err != nil {
		return err
	}

	fmt.Println("🗑️ Pin deleted")

	return nil
}

var PinCmd = &cobra.Command{
	Use:   "pin add/del [Name/url ofo a dotfile]",
	Short: "Manage your pins (add or delete)",
	Args:  cobra.MinimumNArgs(2),
	Example: `
	dona pin add user/dotfiles -t fedora
	dona pin del user/dotfiles
	dona pin del fedora # delete all pins with tag fedora
	`,
	Run: func(cmd *cobra.Command, args []string) {

		tag, error := cmd.Flags().GetString("tag")
		if error != nil || tag == "" && args[0] == "add" {
			fmt.Println(color.RedString("Error: ") + "Value for a tag is required")
			cmd.Help()
			return
		}

		if args[0] != "add" && args[0] != "del" {
			fmt.Println(color.RedString("Error: ") + "Command not recognized")
			cmd.Help()
			return
		}

		home, error := util.GetHome()

		if error != nil {
			fmt.Println(error)
			return
		}

		if args[0] == "add" {
			if error := addPin(home, args); error != nil {
				fmt.Println(error)
			}
		} else {
			if err := delPin(home, args); err != nil {
				fmt.Println(err)
			}
		}
	},
}
