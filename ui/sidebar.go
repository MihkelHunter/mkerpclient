package ui

import (
	"mkerpclient/state"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BuildSidebar(mainContent *fyne.Container) *fyne.Container {
	sidebar := container.NewVBox(
		widget.NewButton("Sales", func() { state.SetMenuState("SalesOrders"); ShowSales(mainContent) }),
		widget.NewButton("Purchase", func() { state.SetMenuState("PurchaseOrders"); ShowPurchase(mainContent) }),
		widget.NewButton("Inventory", func() { state.SetMenuState("StockCheck"); ShowInventory(mainContent) }),
	)

	return sidebar
}
