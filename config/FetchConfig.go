package config

import (
	"fmt"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v3"
)

type UIConfig struct {
	Theme        string `yaml:"theme"`
	WindowWidth  int    `yaml:"window_width"`
	WindowHeight int    `yaml:"window_height"`
}

func DefaultConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("get user config directory: %w", err)
	}
	return filepath.Join(configDir, "GoSleepy", "main.yaml"), nil
}

func LoadConfig(path string) (UIConfig, error) {
	var config UIConfig
	if path == "" {
		defaultPath, err := DefaultConfigPath()
		if err != nil {
			return config, err
		}
		path = defaultPath
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("read config %q: %w", path, err)
	}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return config, fmt.Errorf("parse config %q: %w", path, err)
	}
	return config, nil
}
