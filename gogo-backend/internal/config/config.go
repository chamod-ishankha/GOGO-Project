package config

import (
	"fmt"
	"os"

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

	// Read the file from the current working directory
	data, err := os.ReadFile(filename)
	if err != nil {
		// Get current working directory to show in the error message
		cwd, _ := os.Getwd()
		return nil, fmt.Errorf("could not find %s in %s: %w", filename, cwd, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("error parsing yaml: %w", err)
	}

	return &cfg, nil
}
