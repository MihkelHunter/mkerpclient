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

/*func ShowPurchase(mainContent *fyne.Container) {
	//data, title := data.GetPurchaseData(state.GetMenuState())
	print("Purchase Orders button clicked\n")
	//setMenuState("PurchaseOrders")
	state.SetMenuState("PurchaseOrders")
	showPurchase(mainContent)
}*/

func ShowPurchase(mainContent *fyne.Container) {
	/*purchData := [][]string{
		{"Invoice #", "Amount"},
		{"A100", "$300"},
		{"A101", "$450"},
		{"A102", "$200"},
	}
	supplierData := [][]string{
		{"Supplier ID", "Name"},
		{"C001", "Supplier A"},
		{"C002", "Supplier B"},
		{"C003", "Supplier C"},
	}*/
	//state.SetMenuState("PurchaseOrders")

	purchaseTopBar := container.NewVBox(
		widget.NewButton("Purchase Orders", func() { ShowPurchaseOrders(mainContent) }),
		widget.NewButton("Suppliers", func() { ShowSuppliers(mainContent) }),
	)

	var data [][]string
	var title string
	/*if state.GetMenuState() == "Suppliers" {
		data = supplierData
		title = "Supplier Data"
	} else {
		data = purchData
		title = "Purchase Data"
	}*/

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
	//setMenuState("Suppliers")
	state.SetMenuState("Suppliers")
	ShowPurchase(mainContent)
}

func ShowPurchaseOrders(mainContent *fyne.Container) {
	print("Purchase Orders button clicked\n")
	//setMenuState("PurchaseOrders")
	state.SetMenuState("PurchaseOrders")
	ShowPurchase(mainContent)
}
