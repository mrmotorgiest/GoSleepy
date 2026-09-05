package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("GoSleepy")

	// Splash content
	background := canvas.NewRectangle(color.Transparent)

	text := canvas.NewText("Loading GoSleepy...", color.White)
	text.TextSize = 30
	text.Alignment = fyne.TextAlignCenter

	splashContent := container.NewMax(
		background,
		container.NewCenter(text),
	)

	w.SetContent(splashContent)
	w.SetFullScreen(true)
	w.SetPadded(false)
	w.Show()

	go func() {
		// Perform startup work here.
		time.Sleep(2 * time.Second)

		fyne.Do(func() {
			// Replace the splash content in the same window.
			mainContent := container.NewCenter(
				widget.NewLabel("GoSleepy is ready!"),
			)

			w.SetContent(mainContent)
			w.SetFullScreen(false)
			w.Resize(fyne.NewSize(500, 300))
			w.CenterOnScreen()
		})
	}()

	a.Run()
}
