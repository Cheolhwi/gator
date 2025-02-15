package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

type State struct {
	Config *Config
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, configFileName), nil
}

func Read() (Config, error) {
	var cfg Config
	filepath, err := getConfigFilePath()
	if err != nil {
		return cfg, fmt.Errorf("read config file failed: %w", err)
	}
	file, err := os.Open(filepath)
	if err != nil {
		return cfg, fmt.Errorf("open config file failed: %w", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return cfg, fmt.Errorf("decode json file failed: %w", err)
	}
	return cfg, nil
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	return write(*cfg)
}

func write(cfg Config) error {
	filepath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("read config file failed: %w", err)
	}

	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("create config file failed: %w", err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	if err := encoder.Encode(cfg); err != nil {
		return fmt.Errorf("encode json file failed: %w", err)
	}
	return nil
}
