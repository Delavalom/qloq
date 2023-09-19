package config

import (
	"os"
	"path/filepath"
	"strings"
)

func Create(apiKey string) (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir = filepath.Join(dir, ".qloq")
	file, err := os.Create(dir)

	if err != nil {
		return "", err
	}

	defer file.Close()
	_, err = file.WriteString("API_KEY=" + apiKey)

	if err != nil {
		return "", err
	}

	return "Config file created succesfully", nil
}

func Retrieve() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir = filepath.Join(dir, ".qloq")

	content, err := os.ReadFile(dir)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(strings.Split(string(content), "=")[1]), nil
}
