package cmd

import (
	"fmt"
	"os"

	"encoding/json"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Returns a new slice of Pins containing all strings in the
// slice that satisfy the predicate `f`.
func FilterPins(vs []Pin, f func(Pin) bool) []Pin {
    vsf := make([]Pin, 0)
    for _, v := range vs {
        if f(v) {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

func printDots(home string) {
	path := home + "/.dona/dots"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("No dotfiles found")
		return
	}

	files, err := os.ReadDir(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(color.YellowString(path))
	items := len(files)
	for i, entry := range files {
		fmt.Println(color.CyanString("│"))
		if i == items - 1 {
			fmt.Println(color.CyanString("└── ") + entry.Name())
		} else {
			fmt.Println(color.CyanString("├── ") + entry.Name())
		}
	}

}

func removeDuplicateValues(slice []string) []string {
    keys := make(map[string]bool)
    list := []string{}
 
    // If the key(values of the slice) is not equal
    // to the already present value in new slice (list)
    // then we append it. else we jump on another element.
    for _, entry := range slice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }

    return list
}

var ListCmd = &cobra.Command{
	Use:   "list dots/pins",
	Short: "List dotfiles you saved or pins",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]

		home, fsErr := GetHome()

		if fsErr != nil {
			fmt.Println(fsErr)
			return
		}

		if arg == "dots" {
			printDots(home)
		} else if arg == "pins" {
			file, err  := os.ReadFile(home + "/.dona/pins.json")
			pins := []Pin{}
			if err != nil {
				fmt.Println(err)
				return
			}

			if e := json.Unmarshal([]byte(file), &pins); e != nil {
				fmt.Println(e)
				return
			}

			tags := []string{}
			for _, pin := range pins {
				tags = append(tags, pin.Tag)
			}

			tags = removeDuplicateValues(tags)

			fmt.Println(color.YellowString("Pins"))
			fmt.Println(color.CyanString("│"))
			for _, tag := range tags {
				pins := FilterPins(pins, func(pin Pin) bool {
					return pin.Tag == tag
				})

				fmt.Println(color.CyanString("└── ") + tag)
				for _, pin := range pins {
					fmt.Println(color.CyanString("\t└── ") + pin.Name)
				}

			}


		}

	},
}

