package layout
import (
	"fyne.io/fyne/v2"
	"main/config"
)

func BuildMenuItems(items []config.MenuItem) []*fyne.MenuItem {
	menuItems := make([]*fyne.MenuItem, 0, len(items))
	for _, item := range items {
		menuItem := fyne.NewMenuItem(item.Label, nil)
		if len(item.Subitems) > 0 {
			menuItem.ChildMenu = fyne.NewMenu(item.Label, BuildMenuItems(item.Subitems)...)
		}
		menuItems = append(menuItems, menuItem)
	}
	return menuItems
}
