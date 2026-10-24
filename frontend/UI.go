package frontend

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"main/config"
)

func ShowMain(a fyne.App, config config.UIConfig) {
	window := a.NewWindow(config.Title)
	
	a.Settings().SetTheme(theme.DarkTheme())
	window.Resize(fyne.NewSize(float32(config.WindowWidth), float32(config.WindowHeight)))
	
	window.SetContent(container.NewWithoutLayout())
	window.SetMainMenu(BuildMain(config.MenuBar.Items))
	window.Show()
}

func BuildMain(items []config.MenuItem) *fyne.MainMenu {
	menus := make([]*fyne.Menu, 0, len(items))
	for _, item := range items {
		menus = append(menus, fyne.NewMenu(item.Label, buildMenuItems(item.Subitems)...))
	}
	return fyne.NewMainMenu(menus...)
}

func buildMenuItems(items []config.MenuItem) []*fyne.MenuItem {
	menuItems := make([]*fyne.MenuItem, 0, len(items))
	for _, item := range items {
		menuItem := fyne.NewMenuItem(item.Label, nil)
		if len(item.Subitems) > 0 {
			menuItem.ChildMenu = fyne.NewMenu(item.Label, buildMenuItems(item.Subitems)...)
		}
		menuItems = append(menuItems, menuItem)
	}
	return menuItems
}
