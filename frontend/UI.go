package frontend

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func ShowSplashScreen(a fyne.App) { // sum generated placeholder
	splash := a.NewWindow("Loading")
	splash.Resize(fyne.NewSize(400, 220))
	splash.SetFixedSize(true)

	title := canvas.NewText("GoSleepy", theme.ForegroundColor())
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
	splash.Show()
}