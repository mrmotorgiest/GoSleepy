package main
import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"main/config"
	"main/config/debug"
)

func main() {
	a := app.New()
	config := config.FetchConfig("DEFAULT")
	debug.PrintConfig(config)
	fmt.Println("int3:", config.Root.SubPackages["something"].SubPackages["class"].Variables["int2"])
	a.Run()
}
