package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func ShowAbout(mainContent *fyne.Container) {
	title := canvas.NewText("About MKERP Client", color.Black)
	title.Alignment = fyne.TextAlignCenter

	mainContent.Objects = []fyne.CanvasObject{
		container.NewBorder(
			title,
			nil, nil, nil,
			container.NewBorder(
				canvas.NewText("Creating a hobby ERP client in Go", color.Black),
				canvas.NewText("MKERP Client version 0.1 © 2025 ", color.Black),
				nil, nil, nil,
			),
		),
	}
	//mainContent.Refresh()
}
