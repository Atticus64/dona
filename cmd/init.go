package cmd

import (
	"os"
)

func getDir () (string, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return dirname, nil
}


