package main

import (
	"fmt"
	ui "github.com/VladimirMarkelov/clui"
	"strings"
)

func createButtons(mainFrame *ui.Frame) {
	w, h := ui.ScreenSize()
	frame := ui.CreateFrame(mainFrame, w-1, h/4-1, ui.BorderThin, ui.Fixed)
	startBtn := ui.CreateButton(frame, w/4, ui.AutoSize,"Start", ui.Fixed)
	quitBtn := ui.CreateButton(frame, w/4, ui.AutoSize,"Quit", ui.Fixed)
	InstallBtn := ui.CreateButton(frame, w/4, ui.AutoSize,"Install", ui.Fixed)
	startBtn.OnClick(func(event ui.Event) {
		frame.Destroy()
		go startVer(mainFrame)
	})
	InstallBtn.OnClick(func(event ui.Event) {
		frame.Destroy()
		go selectVer(mainFrame)
	})
	quitBtn.OnClick(func(event ui.Event) {
		go ui.Stop()
	})

}

func selectVer(frame *ui.Frame) {
	w, h := ui.ScreenSize()
	selectFrame := ui.CreateFrame(frame, w-1, h-1, ui.BorderNone, ui.Fixed)
	selectFrame.SetPack(ui.Vertical)
	selectFrame.SetScrollable(true)
	for _, version := range getAvailableVersions(){
		label := fmt.Sprintf("%v (%v)", version.Id, version.Type)
		btn := ui.CreateButton(selectFrame, 40, ui.AutoSize, label, 1)

		btn.OnClick(func(ev ui.Event) {
			go ui.Stop()
		})
	}
}

func startVer(frame *ui.Frame) {
	w, h := ui.ScreenSize()
	selectFrame := ui.CreateFrame(frame, w-1, h-1, ui.BorderNone, ui.Fixed)
	selectFrame.SetPack(ui.Vertical)
	selectFrame.SetScrollable(true)
	for _, version := range getInstalledVersions(){
		label := fmt.Sprintf("%v (%v)", version.Id, version.Type)
		btn := ui.CreateButton(selectFrame, 40, ui.AutoSize, label, 1)

		btn.OnClick(func(ev ui.Event) {
			verID := strings.Split(btn.Title(), " ")[0]
			ui.Logger().Println("Libraries: " + getLibs(verID))
			ui.Logger().Println("Command: "+getArgs(verID))
			ui.Stop()
			go Launch(verID)


		})
	}

}