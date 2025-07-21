package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func(c *Config) SetUser(username string) error {
	c.Current_user_name = username
	if err := write(*c); err != nil {
		return fmt.Errorf("write function error")
	}
	return nil
}


func write(cfg Config) error {
	file, err := GetConfigFilePath()
	if err != nil {
		return fmt.Errorf("file not found %w", err)
	}
	
	updated, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling json: %w", err)
	}

	if err := os.WriteFile(file, updated, 0644); err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}