package config

import (
	"encoding/json"
	"os"
)

const permission = 0644

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name,omitempty"`
}

func (cfg *Config) SetUser(name string) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	cfg.CurrentUserName = name
	dat, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, dat, permission)
	if err != nil {
		return err
	}

	return nil
}
