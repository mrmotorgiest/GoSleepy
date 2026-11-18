package config
import (
	"main/config/decode"
	"os"
	"path/filepath"
)

func FetchConfig(path string) decode.DConfig {
	if path == "DEFAULT" {
		path = filepath.Join(os.Getenv("HOME"), ".config", "GoSleepy")
	}

	mainPath := filepath.Join(path, "main.gsc")
	_, err := os.Stat(mainPath)
	if err != nil {
		panic(err)
	}

	return decode.DecodeFile(mainPath)
}