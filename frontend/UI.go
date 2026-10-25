package frontend

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"main/config"
	"main/frontend/layout"
	"main/frontend/style"
)

func ShowMain(a fyne.App, config config.UIConfig) error {
	if err := style.SetFontTheme(a, style.FontConfig{
		Path: config.Font.Main.Path,
		Size: config.Font.Main.Size,
	}); err != nil {
		return fmt.Errorf("set font theme: %w", err)
	}

	window := a.NewWindow(config.Title)
	window.Resize(fyne.NewSize(float32(config.WindowWidth), float32(config.WindowHeight)))

	window.SetMainMenu(BuildMain(config.MenuBar.Items))
	window.Show()
	return nil
}

func BuildMain(items []config.MenuItem) *fyne.MainMenu {
	menus := make([]*fyne.Menu, 0, len(items))
	for _, item := range items {
		menus = append(menus, fyne.NewMenu(item.Label, layout.BuildMenuItems(item.Subitems)...))
	}
	return fyne.NewMainMenu(menus...)
}
