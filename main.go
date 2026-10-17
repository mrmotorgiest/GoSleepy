package main

import (
	"fyne.io/fyne/v2/app"
	"main/frontend"
)

func main() {
	a := app.New()
	var splashShown = frontend.ShowSplashScreen(a)
	if splashShown {
		frontend.ShowMain(a)
	}
	a.Run()
}
