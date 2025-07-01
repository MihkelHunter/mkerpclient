package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myapp := app.New()
	myWindow := myapp.NewWindow("Mkerp Client")
	//myWindow.Resize(fyne.NewSize(800, 600))

	mainContent := container.NewMax() // Use NewMax to allow child to fill all space

	// Helper to create a table with minimum cell size
	createTable := func(data [][]string) *widget.Table {
		table := widget.NewTable(
			func() (int, int) { return len(data), len(data[0]) },
			func() fyne.CanvasObject { return widget.NewLabel("") },
			func(id widget.TableCellID, o fyne.CanvasObject) {
				o.(*widget.Label).SetText(data[id.Row][id.Col])
			},
		)
		table.SetColumnWidth(0, 120)
		table.SetColumnWidth(1, 120)
		return table
	}

	showSales := func() {
		data := [][]string{
			{"Order #", "Amount"},
			{"12345", "$1000"},
			{"12346", "$500"},
			{"12347", "$750"},
		}
		mainContent.Objects = []fyne.CanvasObject{
			container.NewBorder(
				canvas.NewText("Sales Data", color.White), // top
				nil, nil, nil,
				createTable(data),
			),
		}
		mainContent.Refresh()
	}
	showPurchase := func() {
		data := [][]string{
			{"Invoice #", "Amount"},
			{"A100", "$300"},
			{"A101", "$450"},
			{"A102", "$200"},
		}
		mainContent.Objects = []fyne.CanvasObject{
			container.NewBorder(
				canvas.NewText("Purchase Data", color.White),
				nil, nil, nil,
				createTable(data),
			),
		}
		mainContent.Refresh()
	}
	showInventory := func() {
		data := [][]string{
			{"Item", "Stock"},
			{"Widget A", "50"},
			{"Widget B", "20"},
			{"Widget C", "0"},
		}
		mainContent.Objects = []fyne.CanvasObject{
			container.NewBorder(
				canvas.NewText("Inventory Data", color.White),
				nil, nil, nil,
				createTable(data),
			),
		}
		mainContent.Refresh()
	}

	// Sidebar with buttons
	sidebar := container.NewVBox(
		widget.NewButton("Sales", func() { showSales() }),
		widget.NewButton("Purchase", func() { showPurchase() }),
		widget.NewButton("Inventory", func() { showInventory() }),
	)

	// Set initial content
	showSales()

	// Use HSplit for resizable sidebar
	split := container.NewHSplit(sidebar, mainContent)
	split.Offset = 0.2 // Sidebar takes 20% of the width

	myWindow.SetContent(split)
	myWindow.ShowAndRun()
}
