package datafetch

func GetPurchaseData(state string) ([][]string, string) {
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

func GetInventoryData(state string) ([][]string, string) {
	itemData := [][]string{
		{"Item", "Stock"},
		{"Widget A", "50"},
		{"Widget B", "20"},
		{"Widget C", "0"},
	}
	if state == "StockCheck" {
		return itemData, "Stock Check data"
	}
	return itemData, "Add Item Data"
}

func GetSalesData(state string) ([][]string, string) {
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
		return salesData, "Sales Orders Data"
	}
	if state == "ClientList" {
		return clientList, "Client List Data"
	}

	return [][]string{}, "Unknown state: " + state
}
