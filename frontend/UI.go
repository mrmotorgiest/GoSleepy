package frontend

import (
	"fyne.io/fyne/v2"
	"fmt"
	"main/config"
)

func ShowMain(a fyne.App , config config.UIConfig) {
	var splash = a.NewWindow(config.Title)
	fmt.Println("Window Width:", config.WindowWidth)
	fmt.Println("Window Height:", config.WindowHeight)
	fmt.Println("Window Title:", config.Title)
	splash.SetFixedSize(true)
	splash.Resize(fyne.NewSize(float32(config.WindowWidth), float32(config.WindowHeight)))
	splash.Show()
}