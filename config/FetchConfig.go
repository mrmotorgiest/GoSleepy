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

// config \\
func FetchConfig(path string) (UIConfig, error) {
	var config UIConfig
	
	configPath := configFilePath(path)
	yamlFile, err := os.ReadFile(configPath)
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return config, fmt.Errorf("Something went wrong with the YAML file %q: %w", configPath, err)
	}

	return config, err
}

func DecodeConfig(path string) (UIConfig, error) {
	config, err := FetchConfig(path)

	menuPath, ok := config.Register["menu_bar"]
	if !filepath.IsAbs(menuPath) {
		menuPath = filepath.Join(filepath.Dir(configFilePath(path)), menuPath)
	}

	menuConfig, err := FetchConfig(menuPath)
	config.MenuBar = menuConfig.MenuBar

	if err != nil || !ok {
		return config, fmt.Errorf("failed to load config %q: %w", menuPath, err)
	}
	return config, nil
}