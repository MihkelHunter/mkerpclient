package datafetch

import "errors"

func GetPurchaseData(state string) ([][]string, string, error) {
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
		return supplierData, "Supplier Data", nil
	}
	if state == "PurchaseOrders" {
		return purchData, "Purchase Data", nil
	}
	return [][]string{}, "Unknown state: " + state, errors.New("error: unknown state")
}

func GetInventoryData(state string) ([][]string, string, error) {
	itemData := [][]string{
		{"Item", "Stock"},
		{"Widget A", "50"},
		{"Widget B", "20"},
		{"Widget C", "0"},
	}
	if state == "StockCheck" {
		return itemData, "Stock Check data", nil
	}
	if state == "AddItem" {
		return itemData, "Add Item data", nil
	}

	return [][]string{}, "Unknown state: " + state, errors.New("error: unknown state")
}

func GetSalesData(state string) ([][]string, string, error) {
	salesData := [][]string{
		{"Order #", "Amount"},
		{"12345", "$1000"},
		{"12346", "$500"},
		{"12347", "$750"},
	}
	clientList := [][]string{
		{"Client ID", "Client name"},
		{"1", "Google"},
		{"2", "Alphabet"},
		{"3", "Facebook"},
		{"4", "Meta"},
	}
	if state == "SalesOrders" {
		return salesData, "Sales Orders Data", nil
	}
	if state == "ClientList" {
		return clientList, "Client List Data", nil
	}

	return [][]string{}, "Unknown state: " + state, errors.New("error: unknown state")
}
