package main

import (
	"fyne.io/fyne/v2/app"
	"main/frontend"
	"main/config"
)

func main() {
	a := app.New()
	var splashShown = frontend.ShowSplashScreen(a)
	if splashShown {
		var config, err = config.DecodeConfig("main.yaml")
		if err != nil {
			panic(err)
		}
		frontend.ShowMain(a, config)
	}
	a.Run()
}
