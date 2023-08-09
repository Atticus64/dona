package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Dona",
	Long:  `All software has versions. This is Dona's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Dona version 0.0.2 beta")
	},
}
