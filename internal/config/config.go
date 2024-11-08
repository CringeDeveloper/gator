package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name,omitempty"`
}

func (c *Config) SetUser(name string) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	c.CurrentUserName = name
	f, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, f, 0644)
	if err != nil {
		return err
	}

	return nil
}
