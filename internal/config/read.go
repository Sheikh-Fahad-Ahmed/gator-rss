package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read() (Config, error) {
	file, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("%w", err)
	}
	content, err := os.ReadFile(file)
	if err != nil {
		return Config{}, fmt.Errorf("error opening file: %w", err)
	}

	var cfg Config
	err = json.Unmarshal(content, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("error decoding json: %w", err)
	}
	return cfg, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("home directory not found")
	}
	filePath := fmt.Sprintf("%s/%s", homeDir, configFilename)
	return filePath,nil
}