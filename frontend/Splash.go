package frontend

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"time"
)

var ShowSplashScreen = func(a fyne.App) bool {
	var splash = a.NewWindow("GoSleepy Splash")
	splash.Resize(fyne.NewSize(400, 400))

	title := widget.NewLabel("GoSleepy")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	subtitle := widget.NewLabel("Currently in deving")
	subtitle.Alignment = fyne.TextAlignCenter

	content := container.NewPadded(
		container.NewVBox(title, subtitle),
	)
	splash.SetContent(content)
	splash.CenterOnScreen()

	splash.Show()
	time.Sleep(time.Second * 1)
	splash.Close()
	return true
}
