package frontend

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"time"
)

var ShowSplashScreen = func(a fyne.App) bool {
	var splash = a.NewWindow("GoSleepy Splash")
	splash.SetFixedSize(true)
	splash.Resize(fyne.NewSize(400, 400))

	title := widget.NewLabel("GoSleepy")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	subtitle := widget.NewLabel("Preparing your workspace")
	subtitle.Alignment = fyne.TextAlignCenter

	content := container.NewPadded(
		container.NewVBox(title, subtitle),
	)
	splash.SetContent(content)
	splash.CenterOnScreen()

	splash.Show()
	time.Sleep(time.Second * 3)
	splash.Close()
	return true
}
