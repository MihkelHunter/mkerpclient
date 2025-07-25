package main

import (
	_ "embed"
	"mkerpclient/ui"

	"mkerpclient/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

//go:embed ui/images/MKERP_icon_256x256.png
var iconData []byte

func setWindowDefaults(w fyne.Window) {
	w.SetMaster()
	w.Resize(fyne.NewSize(800, 600))
	w.SetFixedSize(false)
	w.CenterOnScreen()
	w.SetPadded(true)

	icon := fyne.NewStaticResource("MKERP_icon_256x256.png", iconData)
	w.SetIcon(icon)
}

func main() {
	myapp := app.NewWithID("mkerpclient")
	myapp.Settings().SetTheme(&theme.CustomTheme{})
	myWindow := myapp.NewWindow("Mkerp Client")
	setWindowDefaults(myWindow)

	mainContent := container.NewStack()
	sidebar := ui.BuildSidebar(mainContent)

	ui.ShowAbout(mainContent)

	split := container.NewHSplit(sidebar, mainContent)
	split.Offset = 0.2

	myWindow.SetContent(split)
	myWindow.ShowAndRun()
}
