package main

import (
	"main/frontend"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	frontend.ShowSplashScreen(a)
	frontend.ShowMain(a)
}
