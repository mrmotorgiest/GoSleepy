package config

import (
	"fmt"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v3"
)

var DefaultConfigPath = filepath.Join(os.Getenv("HOME"), ".config", "GoSleepy", "main.yaml") 
type UIConfig struct {
	WindowWidth  int    `yaml:"window_width"`
	WindowHeight int    `yaml:"window_height"`
	Title        string `yaml:"title"`
}

func FetchConfig(path string) (UIConfig, error) {
	var config UIConfig
	if path == "" {
		path = DefaultConfigPath
	}

	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("failed to read config file: %w", err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return config, fmt.Errorf("failed to unmarshal config file: %w", err)
	}
	fmt.Println("Config loaded successfully from", path)
	return config, nil
}
