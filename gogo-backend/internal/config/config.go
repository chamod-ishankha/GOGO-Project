package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port   string `yaml:"port"`
		Prefix string `yaml:"prefix"` // e.g., /api/v1/gogo
	} `yaml:"server"`
	Database struct {
		DSN string `yaml:"dsn"` // Data Source Name
	} `yaml:"database"`
}

func LoadConfig(serviceName string) (*Config, error) {
	filename := "config." + serviceName + ".yaml"

	configDir := os.Getenv("CONFIG_DIR")
	if configDir == "" {
		return nil, fmt.Errorf("CONFIG_DIR environment variable not set")
	}

	fullPath := filepath.Join(configDir, filename)

	data, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, fmt.Errorf("could not find %s at %s: %w", filename, fullPath, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
