package ui

import (
	"image/color"
	"mkerpclient/datafetch"
	"mkerpclient/state"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ShowSales(mainContent *fyne.Container) {

	if state.GetMenuState() == "" || state.GetMenuState() == "PurchaseOrders" {
		state.SetMenuState("SalesOrders")
	}

	salesTopBar := container.NewHBox(
		widget.NewButton("Sales Orders", func() { ShowSalesOrders(mainContent) }),
		widget.NewButton("Client List", func() { ShowClientList(mainContent) }),
		layout.NewSpacer(),
	)

	data, title, err := datafetch.GetSalesData(state.GetMenuState())
	if err != nil {
		println("Error fetching sales data:", err)
	}

	mainContent.Objects = []fyne.CanvasObject{
		container.NewBorder(
			CreateToolbar(),
			nil, nil, nil,
			container.NewBorder(
				salesTopBar,
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

func ShowSalesOrders(mainContent *fyne.Container) {
	print("Sales Orders button clicked\n")
	state.SetMenuState("SalesOrders")
	ShowSales(mainContent)
}

func ShowClientList(mainContent *fyne.Container) {
	print("Client List button clicked\n")
	state.SetMenuState("ClientList")
	ShowSales(mainContent)
}
