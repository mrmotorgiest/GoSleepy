package frontend

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowMain(a fyne.App) {
	var splash = a.NewWindow("GoSleepy")
	splash.SetFixedSize(true)
	splash.Resize(fyne.NewSize(1920, 1080))

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
}