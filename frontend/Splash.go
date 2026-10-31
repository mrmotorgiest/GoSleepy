package frontend
import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// still an placeholder
func ShowSplashScreen(app fyne.App) {
	splash := app.NewWindow("GoSleepy Splash")
	splash.Resize(fyne.NewSize(400, 400))

	title := widget.NewLabel("GoSleepy")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	subtitle := widget.NewLabel("In development")
	subtitle.Alignment = fyne.TextAlignCenter

	content := container.NewPadded(
		container.NewVBox(title, subtitle),
	)
	splash.SetContent(content)
	splash.CenterOnScreen()

	splash.Show()
	time.AfterFunc(time.Second, func() {
		fyne.Do(splash.Close)
	})
}
