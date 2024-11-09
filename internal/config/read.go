package config

import (
	"encoding/json"
	"os"
)

func Read() (Config, error) {
	cfg := Config{}

	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	res, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(res, &cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
