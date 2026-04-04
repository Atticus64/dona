package util

import (
	"fmt"

	"github.com/spf13/cobra"
)

func GetVersion() string {
	return "0.0.6"
}

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Dona",
	Long:  `All software has versions. This is Dona's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Dona version " + GetVersion())
	},
}
