package frontend

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"time"
)

func ShowSplashScreen(a fyne.App) {
	splash := a.NewWindow("Initializing...")
	splash.Resize(fyne.NewSize(220, 220))
	splash.SetFixedSize(true)
	title := canvas.NewText("GoSleepy", theme.ForegroundColor())
	title.TextStyle = fyne.TextStyle{Bold: true}
	subtitle := canvas.NewText("Initializing...", theme.ForegroundColor())
	progress := widget.NewProgressBarInfinite()
	content := container.NewCenter(container.NewVBox(title, subtitle, progress))
	splash.SetContent(content)
	splash.CenterOnScreen()
	splash.Show()
	time.Sleep(2 * time.Second)
	splash.Close()
}
