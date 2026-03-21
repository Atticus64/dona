package util

import "os"

func GetHome() (string, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return dirname, nil
}

func CheckDir(dirname string) error {
	if _, err := os.Stat(dirname); err != nil {
		if err := os.Mkdir(dirname, 0755); err != nil {
			panic(err)
		}
	}

	return nil
}
