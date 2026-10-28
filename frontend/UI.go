package frontend
import (
	"fmt"

	"fyne.io/fyne/v2"
	"main/config"
	"main/frontend/layout"
	"main/frontend/style"
)

// UI \\
func ShowMain(app fyne.App, uiConfig config.UIConfig) error {
	if err := style.SetFontTheme(app, style.FontConfig{
		Path: uiConfig.Font.Main.Path,
		Size: uiConfig.Font.Main.Size,
	}); err != nil {
		return fmt.Errorf("set font theme: %w", err)
	}

	BuildUI(app, &uiConfig)
	return nil
}

func BuildUI(app fyne.App, uiConfig *config.UIConfig, content ...fyne.CanvasObject) fyne.Window {
	window := app.NewWindow(uiConfig.Title)
	window.Resize(fyne.NewSize(float32(uiConfig.WindowWidth), float32(uiConfig.WindowHeight)))
	
	window.SetMainMenu(BuildMainMenu(uiConfig.MenuBar.Items))
	window.SetContent(layout.BuildContent(content...))

	window.Show()
	return window
}

// Items \\
func BuildMainMenu(items []config.MenuItem) *fyne.MainMenu {
	menus := make([]*fyne.Menu, 0, len(items))
	for _, item := range items {
		menus = append(menus, fyne.NewMenu(item.Label, layout.BuildMenuItems(item.Subitems)...))
	}
	return fyne.NewMainMenu(menus...)
}