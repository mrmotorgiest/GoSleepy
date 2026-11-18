package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2/app"
	"main/config"
	"main/config/debug"
)

func main() {
	// The app runs in a sandboxed Wayland environment in this workspace, which
	// triggers GTK icon loader failures before the window can open. Forcing the
	// X11 backend avoids that crash while keeping the Fyne UI intact.
	os.Setenv("GDK_BACKEND", "x11")
	os.Setenv("XDG_BACKEND", "x11")

	a := app.New()
	config := config.FetchConfig("DEFAULT")
	debug.PrintConfig(config)
	fmt.Println("int3:", config.Root.SubPackages["something"].Variables["int3"])
	a.Run()
}
