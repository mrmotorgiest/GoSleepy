package main
import (
	"fyne.io/fyne/v2/app"
	"main/config"
	"main/frontend"
)

func main() {
	a := app.New()
	frontend.ShowSplashScreen(a)

	uiConfig, err := config.DecodeConfig("main.yaml")
	if err != nil {
		panic(err)
	}
	if err := frontend.ShowMain(a, uiConfig); err != nil {
		panic(err)
	}
	a.Run()
}
