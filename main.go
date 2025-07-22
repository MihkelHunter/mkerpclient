package main

import (
	"mkerpclient/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

/*type CurrentMenuState struct {
	CurrentMenuState string
}

var currentState = &CurrentMenuState{CurrentMenuState: "PurchaseOrders"}

func showPurchase(mainContent *fyne.Container) {
	purchData := [][]string{
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
	}

	purchaseTopBar := container.NewVBox(
		widget.NewButton("Purchase Orders", func() { showPurchaseOrders(mainContent) }),
		widget.NewButton("Suppliers", func() { showSuppliers(mainContent) }),
	)

	var data [][]string
	var title string
	if currentState.CurrentMenuState == "Suppliers" {
		data = supplierData
		title = "Supplier Data"
	} else {
		data = purchData
		title = "Purchase Data"
	}

	mainContent.Objects = []fyne.CanvasObject{
		container.NewBorder(
			createToolbar(),
			nil, nil, nil,
			container.NewBorder(
				purchaseTopBar,
				nil, nil, nil,
				container.NewBorder(
					canvas.NewText(title, color.White),
					nil, nil, nil,
					container.NewVScroll(createTable(&data)),
				),
			),
		),
	}
	mainContent.Refresh()
}

func showInventory(mainContent *fyne.Container) {
	itemData := [][]string{
		{"Item", "Stock"},
		{"Widget A", "50"},
		{"Widget B", "20"},
		{"Widget C", "0"},
	}

	inventoryTopBar := container.NewVBox(
		widget.NewButton("Stock Check", func() { showStockCheck(mainContent) }),
		widget.NewButton("Add Item", func() { showAddItem(mainContent) }),
	)

	var data [][]string
	var title string
	if currentState.CurrentMenuState == "StockCheck" {
		data = itemData
		title = "Stock Check Data"
	} else {
		data = itemData
		title = "Add Item Data"
	}

	mainContent.Objects = []fyne.CanvasObject{
		container.NewBorder(
			createToolbar(),
			nil, nil, nil,
			container.NewBorder(
				inventoryTopBar,
				nil, nil, nil,
				container.NewBorder(
					canvas.NewText(title, color.White),
					nil, nil, nil,
					container.NewVScroll(createTable(&data)),
				),
			),
		),
	}
	mainContent.Refresh()
}

func showSales(mainContent *fyne.Container) {
	salesData := [][]string{
		{"Order #", "Amount"},
		{"12345", "$1000"},
		{"12346", "$500"},
		{"12347", "$750"},
	}

	salesTopBar := container.NewVBox(
		widget.NewButton("Sales Orders", func() { showSalesOrders(mainContent) }),
		widget.NewButton("Client List", func() { showClientList(mainContent) }),
	)

	var data [][]string
	var title string
	if currentState.CurrentMenuState == "SalesOrders" {
		data = salesData
		title = "Sales Orders Data"
	} else {
		data = salesData
		title = "Client List Data"
	}

	mainContent.Objects = []fyne.CanvasObject{
		container.NewBorder(
			createToolbar(),
			nil, nil, nil,
			container.NewBorder(
				salesTopBar,
				nil, nil, nil,
				container.NewBorder(
					canvas.NewText(title, color.White),
					nil, nil, nil,
					container.NewVScroll(createTable(&data)),
				),
			),
		),
	}
	mainContent.Refresh()
}

func createToolbar() *widget.Toolbar {
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

func createTable(data *[][]string) *widget.Table {
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
}*/

func setWindowDefaults(w fyne.Window) {
	w.SetMaster()
	w.Resize(fyne.NewSize(800, 600))
	w.SetFixedSize(false)
	w.CenterOnScreen()
	w.SetPadded(true)
	w.SetIcon(theme.FyneLogo())
}

/*
func showSalesOrders(mainContent *fyne.Container) {
	print("Sales Orders button clicked\n")
	setMenuState("SalesOrders")
	showSales(mainContent)
}

func showClientList(mainContent *fyne.Container) {
	print("Client List button clicked\n")
	setMenuState("ClientList")
	showSales(mainContent)
}

func showPurchaseOrders(mainContent *fyne.Container) {
	print("Purchase Orders button clicked\n")
	setMenuState("PurchaseOrders")
	showPurchase(mainContent)
}

func showSuppliers(mainContent *fyne.Container) {
	print("Suppliers button clicked\n")
	setMenuState("Suppliers")
	showPurchase(mainContent)
}

func showStockCheck(mainContent *fyne.Container) {
	print("Stock Check button clicked\n")
	setMenuState("StockCheck")
	showInventory(mainContent)
}

func showAddItem(mainContent *fyne.Container) {
	print("Add Item button clicked\n")
	setMenuState("AddItem")
	showInventory(mainContent)
}

func getPurchaseData(state string) ([][]string, string) {
	purchData := [][]string{
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
	}
	if state == "Suppliers" {
		return supplierData, "Supplier Data"
	}
	return purchData, "Purchase Data"
}

func setMenuState(state string) {
	currentState.CurrentMenuState = state
}*/

func main() {
	myapp := app.NewWithID("mkerpclient")
	myWindow := myapp.NewWindow("Mkerp Client")
	setWindowDefaults(myWindow)

	mainContent := container.NewStack()

	//showSales(mainContent)
	//showPurchase(mainContent)
	//showInventory(mainContent)

	/*sidebar := container.NewVBox(
		widget.NewButton("Sales", func() { showSales(mainContent) }),
		widget.NewButton("Purchase", func() { showPurchase(mainContent) }),
		widget.NewButton("Inventory", func() { showInventory(mainContent) }),
	)*/
	sidebar := ui.BuildSidebar(mainContent)

	ui.ShowPurchase(mainContent)

	split := container.NewHSplit(sidebar, mainContent)
	split.Offset = 0.2

	myWindow.SetContent(split)
	myWindow.ShowAndRun()
}
