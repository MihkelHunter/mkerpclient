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
	if state.GetMenuState() == "" || state.GetMenuState() == "PurchaseOrders" {
		state.SetMenuState("StockCheck")
	}
	inventoryTopBar := container.NewHBox(
		widget.NewButton("Stock Check", func() { ShowStockCheck(mainContent) }),
		widget.NewButton("Add Item", func() { ShowAddItem(mainContent) }),
	)

	data, title, err := datafetch.GetInventoryData(state.GetMenuState())
	if err != nil {
		println("Error fetching inventory data:", err)
	}

	mainContent.Objects = []fyne.CanvasObject{
		container.NewBorder(
			CreateToolbar(),
			nil, nil, nil,
			container.NewBorder(
				inventoryTopBar,
				nil, nil, nil,
				container.NewBorder(
					canvas.NewText(title, color.Black),
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
	state.SetMenuState("StockCheck")
	ShowInventory(mainContent)
}

func ShowAddItem(mainContent *fyne.Container) {
	print("Add Item button clicked\n")
	state.SetMenuState("AddItem")
	ShowInventory(mainContent)
}
