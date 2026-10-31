package style

import (
	"fmt"
	"os"
)

func LoadCSS(path string) error {
	if path == "" {
		return nil
	}
	if _, err := os.ReadFile(path); err != nil {
		return fmt.Errorf("read stylesheet %q: %w", path, err)
	}
	return nil
}
