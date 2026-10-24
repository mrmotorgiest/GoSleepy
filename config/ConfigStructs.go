package config
import (
	"os"
	"path/filepath"
)

// Config \\
var DefaultConfigPath = filepath.Join(os.Getenv("HOME"), ".config", "GoSleepy") 
type UIConfig struct {
	WindowWidth  int    `yaml:"window_width"`
	WindowHeight int    `yaml:"window_height"`
	Title        string `yaml:"title"`

	Register map[string]string `yaml:"register"`
	MenuBar  MenuBar           `yaml:"menu_bar"`
}

// Menu bar \\
type MenuBar struct {
	Padding string `yaml:"padding"`
	Spacing string `yaml:"spacing"`

	Orientation string     `yaml:"orientation"`
	Items       []MenuItem `yaml:"items"`
}

type MenuItem struct {
	Label    string     `yaml:"label"`
	Action   string     `yaml:"action"`
	Subitems []MenuItem `yaml:"subitems"`
}
