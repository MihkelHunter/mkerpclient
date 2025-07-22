package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BuildSidebar(mainContent *fyne.Container) *fyne.Container {

	sidebar := container.NewVBox(
		widget.NewButton("Sales", func() { ShowSales(mainContent) }),
		widget.NewButton("Purchase", func() { ShowPurchase(mainContent) }),
		widget.NewButton("Inventory", func() { ShowInventory(mainContent) }),
	)

	return sidebar
}
