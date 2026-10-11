package main

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
	a.Settings().SetTheme(theme.DarkTheme())

	// Splash screen
	splash := a.NewWindow("Loading")
	splash.Resize(fyne.NewSize(400, 220))
	splash.SetFixedSize(true)

	title := canvas.NewText("MY APPLICATION", theme.ForegroundColor())
	title.TextSize = 24
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	subtitle := canvas.NewText("Initializing...", theme.ForegroundColor())
	subtitle.TextSize = 13
	subtitle.Alignment = fyne.TextAlignCenter

	progress := widget.NewProgressBarInfinite()

	content := container.NewCenter(
		container.NewVBox(
			title,
			subtitle,
			progress,
		),
	)

	splash.SetContent(content)
	splash.CenterOnScreen()

	// Main window
	mainWindow := a.NewWindow("My Application")
	mainWindow.Resize(fyne.NewSize(900, 600))
	mainWindow.SetContent(
		container.NewCenter(
			widget.NewLabel("Welcome to my application!"),
		),
	)

	// Show splash
	splash.Show()

	// Simulate initialization
	go func() {
		time.Sleep(3 * time.Second)

		fyne.Do(func() {
			splash.Close()
			mainWindow.Show()
		})
	}()

	a.Run()
}