package config

import (
	"encoding/json"
	"os"
	"path"
)

const (
	configFileName = ".rss-feed-aggregator.json"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := path.Join(homeDir, configFileName)
	return path, nil
}

func write(cfg Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(configPath, data, 0666) // 0666 = perms, everyone can r,w
	if err != nil {
		return err
	}
	return nil
}

func Read() (Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	fileData, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}
	var config Config
	err = json.Unmarshal(fileData, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func (c *Config) SetUser(user string) error {
	c.CurrentUserName = user
	err := write(*c)
	if err != nil {
		return err
	}
	return nil
}


