package ui

import (
	"mkerpclient/state"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func BuildSidebar(mainContent *fyne.Container) *fyne.Container {
	sidebar := container.NewVBox()
	sidebar.Resize(fyne.NewSize(200, 600)) // fixed width
	var btn *widget.Button
	// Helper to create buttons with icon and active highlight
	createMenuButton := func(label string, icon fyne.Resource, menuKey string, onTap func()) *widget.Button {
		btn := widget.NewButtonWithIcon(label, icon, func() {
			state.SetMenuState(menuKey)
			onTap()
			updateActiveHighlight(sidebar, btn)
		})
		btn.Alignment = widget.ButtonAlignLeading
		return btn
	}

	salesBtn := createMenuButton("Sales", theme.ContentAddIcon(), "SalesOrders", func() {
		ShowSales(mainContent)
	})

	purchaseBtn := createMenuButton("Purchase", theme.ContentAddIcon(), "PurchaseOrders", func() {
		ShowPurchase(mainContent)
	})

	inventoryBtn := createMenuButton("Inventory", theme.StorageIcon(), "StockCheck", func() {
		ShowInventory(mainContent)
	})

	sidebar.Add(salesBtn)
	sidebar.Add(purchaseBtn)
	sidebar.Add(inventoryBtn)

	// Set initial active button highlight
	updateActiveHighlight(sidebar, salesBtn)

	return sidebar
}

// updateActiveHighlight highlights the selected button and resets others
func updateActiveHighlight(sidebar *fyne.Container, activeBtn *widget.Button) {
	for _, obj := range sidebar.Objects {
		if btn, ok := obj.(*widget.Button); ok {
			if btn == activeBtn {
				btn.Importance = widget.HighImportance
			} else {
				btn.Importance = widget.MediumImportance
			}
			btn.Refresh()
		}
	}
}

/*package ui

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
}*/
