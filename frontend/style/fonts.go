package style

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type FontConfig struct {
	Path string
	Size float32
}

type fontTheme struct {
	fyne.Theme
	font fyne.Resource
	size float32
}

func (t fontTheme) Font(fyne.TextStyle) fyne.Resource { return t.font }

func (t fontTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText && t.size > 0 {
		return t.size
	}
	return t.Theme.Size(name)
}

func SetFontTheme(a fyne.App, fontConfig FontConfig) error {
	if fontConfig.Path != "" {
		fontData, err := os.ReadFile(fontConfig.Path)
		if err != nil {
			return fmt.Errorf("load main font %q: %w", fontConfig.Path, err)
		}
		a.Settings().SetTheme(fontTheme{
			Theme: theme.DarkTheme(),
			font:  fyne.NewStaticResource(fontConfig.Path, fontData),
			size:  fontConfig.Size,
		})
	} else {
		a.Settings().SetTheme(theme.DarkTheme())
	}
	return nil
}