package config

import (
	"os"
	"path/filepath"
)

// Config \\
var DefaultConfigPath = filepath.Join(os.Getenv("HOME"), ".config", "GoSleepy")

type UIConfig struct {
	WindowWidth  int        `yaml:"window_width"`
	WindowHeight int        `yaml:"window_height"`
	Title        string     `yaml:"title"`
	Font         FontConfig `yaml:"font"`

	Register   map[string]string    `yaml:"register"`
	MenuBar    MenuBar              `yaml:"menu_bar"`
	Components map[string]Component `yaml:"-"`
}

type Component struct {
	Type     string          `yaml:"type"`
	Position string          `yaml:"position"`
	Priority int             `yaml:"priority"`
	Items    []ComponentItem `yaml:"items"`
	Font     string          `yaml:"font"`
	Spacing  string          `yaml:"spacing"`
	Padding  string          `yaml:"padding"`
	CSSPath  string          `yaml:"css"`
}

type ComponentItem struct {
	Label  string `yaml:"label"`
	Type   string `yaml:"type"`
	Action string `yaml:"action"`
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

// Font \\
type FontConfig struct {
	Main FontFaceConfig `yaml:"main"`
}

type FontFaceConfig struct {
	Path string  `yaml:"path"`
	Size float32 `yaml:"size"`
}
