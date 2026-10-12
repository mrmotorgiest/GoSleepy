package main
import (
	"GoSleepy/config"
	"GoSleepy/frontend/UI"
)

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	frontend.ShowSplashScreen(a)

	// Simulate loading time
	time.Sleep(3 * time.Second)

	// Load configuration
	config, err := config.LoadConfig()
	if err != nil {
		// Handle error (e.g., show an error dialog)
	}

	// Proceed with the rest of the application using the loaded config
	// For example, create the main window and set its content
	mainWindow := a.NewWindow("GoSleepy")
	mainWindow.Resize(fyne.NewSize(800, 600))
	mainWindow.SetContent(canvas.NewText("Main Application Window", theme.ForegroundColor()))
	mainWindow.ShowAndRun()
}