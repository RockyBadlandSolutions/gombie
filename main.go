package main

import (
	"fmt"
	ui "github.com/VladimirMarkelov/clui"
)

var minecraftDir = "minecraft"

func main()  {
	ui.InitLibrary()
	defer ui.DeinitLibrary()
	ui.SetThemePath("themes")
	ui.SetCurrentTheme(ui.ThemeNames()[1])
	w, h := ui.ScreenSize()
	window := ui.AddWindow(0, 0, w, h, "MineGO v0")
	window.SetSizable(false)
	mainFrame := ui.CreateFrame(window, w-1, h-1, ui.BorderNone, ui.AutoSize)

	createButtons(mainFrame)

	ui.MainLoop()

	fmt.Println(getMinecraftDirectory())
	fmt.Println(getInstalledVersions())
	fmt.Println(getJAVAExecutable())
	fmt.Println(getAvailableVersions()[6].Url)

}


