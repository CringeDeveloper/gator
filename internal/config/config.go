package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name,omitempty"`
}

func (c *Config) SetUser(name string) {
	path, _ := getConfigFilePath()

	c.CurrentUserName = name
	f, _ := json.Marshal(c)

	_ = os.WriteFile(path, f, 0644)
}
