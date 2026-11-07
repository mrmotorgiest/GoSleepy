package config

import (
	"path/filepath"
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
