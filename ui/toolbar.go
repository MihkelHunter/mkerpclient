package ui

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateToolbar() *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			print("Add action triggered \n")
		}),
		widget.NewToolbarAction(theme.CheckButtonIcon(), func() {
			print("Check action triggered \n")
		}),
		widget.NewToolbarAction(theme.DeleteIcon(), func() {
			print("Delete action triggered \n")
		}),
	)
}
