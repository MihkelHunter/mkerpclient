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

func ShowPurchase(mainContent *fyne.Container) {
	purchaseTopBar := container.NewVBox(
		widget.NewButton("Purchase Orders", func() { ShowPurchaseOrders(mainContent) }),
		widget.NewButton("Suppliers", func() { ShowSuppliers(mainContent) }),
	)

	var data [][]string
	var title string

	data, title = datafetch.GetPurchaseData(state.GetMenuState())

	mainContent.Objects = []fyne.CanvasObject{
		container.NewBorder(
			CreateToolbar(),
			nil, nil, nil,
			container.NewBorder(
				purchaseTopBar,
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

func ShowSuppliers(mainContent *fyne.Container) {
	print("Suppliers button clicked\n")
	state.SetMenuState("Suppliers")
	ShowPurchase(mainContent)
}

func ShowPurchaseOrders(mainContent *fyne.Container) {
	print("Purchase Orders button clicked\n")
	state.SetMenuState("PurchaseOrders")
	ShowPurchase(mainContent)
}
