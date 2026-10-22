package config

import (
	"fmt"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v3"
)

var DefaultConfigPath = filepath.Join(os.Getenv("HOME"), ".config", "GoSleepy") 
type UIConfig struct {
	WindowWidth  int    `yaml:"window_width"`
	WindowHeight int    `yaml:"window_height"`
	Title        string `yaml:"title"`

	Register map[string]string `yaml:"register"`

	MenuBar     MenuBar
}

func FetchConfig(path string) (UIConfig , error) {
	var config UIConfig
	fmt.Println("Fetching config from path:", path)
	var ConfigPath = DefaultConfigPath + path

	fmt.Println("Fetching config from:", ConfigPath)
	yamlFile, err := os.ReadFile(ConfigPath)
	if err != nil {
		return config, fmt.Errorf("failed to read config file: %w", err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return config, fmt.Errorf("error while unmarshalling config file: %v", err)
	}

	return config, nil
}

func DecodeConfig(path string) (UIConfig, error) {
	var config UIConfig
	var err error
	config, err = FetchConfig("main.yaml")
	fmt.Println(config.Register["menu_bar"])
	menuConfig, err := FetchConfig(string(config.Register["menu_bar"]))
	config.MenuBar = menuConfig.MenuBar

	fmt.Println("MenuBar Config:", config.MenuBar)
	return config, err
}