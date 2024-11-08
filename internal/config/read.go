package config

import (
	"encoding/json"
	"os"
)

func Read() (Config, error) {
	c := Config{}

	path, err := getConfigFilePath()

	if err != nil {
		return Config{}, err
	}

	res, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(res, &c)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}
