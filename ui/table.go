package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func CreateTable(data *[][]string) *widget.Table {
	table := widget.NewTable(
		func() (int, int) {
			if len(*data) == 0 {
				return 0, 0
			}
			return len(*data), len((*data)[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText((*data)[id.Row][id.Col])
		},
	)
	table.SetColumnWidth(0, 120)
	table.SetColumnWidth(1, 120)
	return table
}
