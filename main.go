package main

import (
	"mkerpclient/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func setWindowDefaults(w fyne.Window) {
	w.SetMaster()
	w.Resize(fyne.NewSize(800, 600))
	w.SetFixedSize(false)
	w.CenterOnScreen()
	w.SetPadded(true)
	w.SetIcon(theme.FyneLogo())
}

func main() {
	myapp := app.NewWithID("mkerpclient")
	myWindow := myapp.NewWindow("Mkerp Client")
	setWindowDefaults(myWindow)

	mainContent := container.NewStack()
	sidebar := ui.BuildSidebar(mainContent)

	ui.ShowPurchase(mainContent)

	split := container.NewHSplit(sidebar, mainContent)
	split.Offset = 0.2

	myWindow.SetContent(split)
	myWindow.ShowAndRun()
}
