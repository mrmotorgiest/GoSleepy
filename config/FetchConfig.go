package config
import (
	"io/ioutil"
	"os"
	"path/filepath"
	"fmt"
)

// Path \\

// config \\
func FetchConfig(path string) SConfig {
	var config SConfig
	if path == "DEFAULT" {
		path = filepath.Join(os.Getenv("HOME"), ".config", "GoSleepy")
	}

	data, err := ioutil.ReadFile(path + "/main.gsc")
	if err != nil {
		panic(err)
	}
	//pring out data
	fmt.Println(string(data))

	return config
}
