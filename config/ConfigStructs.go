package config

import (
	"os"
	"path/filepath"
)

// Config \\
var DefaultConfigPath = filepath.Join(os.Getenv("HOME"), ".config", "GoSleepy")

type UIConfig struct {
	WindowWidth  int        
	WindowHeight int        
	Title        string  
}

type Component struct {
	Type   string

}