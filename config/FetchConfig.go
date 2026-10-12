package config

import (
	"yaml"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var ConfigPath = "/home/user/.config/GoSleepy/main.yaml"
func LoadConfig() (*Config, error) {
	//placeholder
	return nil, nil
}