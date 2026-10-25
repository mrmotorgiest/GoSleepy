package main

import (
	"fyne.io/fyne/v2/app"
	"main/config"
	"main/frontend"
)

func main() {
	a := app.New()
	var splashShown = frontend.ShowSplashScreen(a)
	if splashShown {
		var config, err = config.DecodeConfig("main.yaml")
		if err != nil {
			panic(err)
		}
		if err := frontend.ShowMain(a, config); err != nil {
			panic(err)
		}
	}
	a.Run()
}
