package ui

import (
	"image/color"
	"mkerpclient/datafetch"
	"mkerpclient/state"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowInventory(mainContent *fyne.Container) {
	/*itemData := [][]string{
		{"Item", "Stock"},
		{"Widget A", "50"},
		{"Widget B", "20"},
		{"Widget C", "0"},
	}*/

	//state.SetMenuState("StockCheck")

	inventoryTopBar := container.NewVBox(
		widget.NewButton("Stock Check", func() { ShowStockCheck(mainContent) }),
		widget.NewButton("Add Item", func() { ShowAddItem(mainContent) }),
	)

	var data [][]string
	var title string
	/*if currentState.CurrentMenuState == "StockCheck" {
		data = itemData
		title = "Stock Check Data"
	} else {
		data = itemData
		title = "Add Item Data"
	}*/

	data, title = datafetch.GetInventoryData(state.GetMenuState())

	mainContent.Objects = []fyne.CanvasObject{
		container.NewBorder(
			CreateToolbar(),
			nil, nil, nil,
			container.NewBorder(
				inventoryTopBar,
				nil, nil, nil,
				container.NewBorder(
					canvas.NewText(title, color.White),
					nil, nil, nil,
					container.NewVScroll(CreateTable(&data)),
				),
			),
		),
	}
	mainContent.Refresh()
}

func ShowStockCheck(mainContent *fyne.Container) {
	print("Stock Check button clicked\n")
	//setMenuState("StockCheck")
	state.SetMenuState("StockCheck")
	ShowInventory(mainContent)
}

func ShowAddItem(mainContent *fyne.Container) {
	print("Add Item button clicked\n")
	//setMenuState("AddItem")
	state.SetMenuState("AddItem")
	ShowInventory(mainContent)
}
