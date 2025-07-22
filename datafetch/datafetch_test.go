package datafetch

import "testing"

func TestGetPurchaseData(t *testing.T) {
	data, title, err := GetPurchaseData("PurchaseOrders")
	if err != nil {
		t.Error("Error fetching purchase orders:", err)
	}
	if title != "Purchase Data" || data[1][0] != "A100" {
		t.Error("Purchase data not returned correctly")
	}
}

func TestGetSuppliersData(t *testing.T) {
	data, title, err := GetPurchaseData("Suppliers")
	if err != nil {
		t.Error("Error fetching supplier data:", err)
	}
	if title != "Supplier Data" || data[1][1] != "Supplier A" {
		t.Error("Supplier data not returned correctly")
	}
}
