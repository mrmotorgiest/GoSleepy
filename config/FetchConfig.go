package config

import (
	"fmt"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v3"
)

// Path \\
func configFilePath(path string) string {
	if filepath.IsAbs(path) {
		return filepath.Clean(path)
	}
	return filepath.Join(DefaultConfigPath, path)
}

func resolveConfigRelativePath(configPath, targetPath string) string {
	if filepath.IsAbs(targetPath) {
		return filepath.Clean(targetPath)
	}
	return filepath.Join(filepath.Dir(configPath), targetPath)
}

// config \\
func FetchConfig(path string) (UIConfig, error) {
	var config UIConfig

	configPath := configFilePath(path)
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		return config, fmt.Errorf("read config %q: %w", configPath, err)
	}
	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		return config, fmt.Errorf("decode config %q: %w", configPath, err)
	}

	return config, nil
}

func DecodeConfig(path string) (UIConfig, error) {
	config, err := FetchConfig(path)
	config.Components = make(map[string]Component)
	mainConfigPath := configFilePath(path)
	if config.Font.Main.Path != "" {
		config.Font.Main.Path = resolveConfigRelativePath(mainConfigPath, config.Font.Main.Path)
	}

	for _, registeredPath := range config.Register {
		registeredConfigPath := resolveConfigRelativePath(mainConfigPath, registeredPath)
		registeredFile, err := os.ReadFile(registeredConfigPath)
		var registered map[string]yaml.Node
		for name, node := range registered {
			if name == "menu_bar" {
				continue
			}

			var component Component
			if component.CSSPath != "" {
				component.CSSPath = resolveConfigRelativePath(mainConfigPath, component.CSSPath)
			}
			config.Components[name] = component
		}
	}
	if err != nil {
		return config, err
	}
	return config, nil
}
