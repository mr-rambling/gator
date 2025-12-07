package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DBURL          string `json:"db_url"`
	CurrenUserName string `json:"current_user_name"`
}

const configFileName = "/.gatorconfig.json"

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func SetUser(cfg Config, username string) error {
	cfg.CurrenUserName = username
	if err := write(cfg); err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path += configFileName

	return path, nil
}

func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	permissions := os.FileMode(0644)
	if err := os.WriteFile(path, jsonData, permissions); err != nil {
		return err
	}
	return nil
}
