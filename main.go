package main

import (
	"fmt"
	"os"
	"github.com/atticus64/dona/cmd"
)

func main() {

	if err := cmd.CreateDona(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
