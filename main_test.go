package main

import (
	"mkerpclient/datafetch"
	"mkerpclient/state"
	"testing"
)

func TestGetPurchaseData(t *testing.T) {
	data, title := datafetch.GetPurchaseData("Suppliers")
	if title != "Supplier Data" || data[1][1] != "Supplier A" {
		t.Error("Supplier data not returned correctly")
	}
	data, title = datafetch.GetPurchaseData("PurchaseOrders")
	if title != "Purchase Data" || data[1][0] != "A100" {
		t.Error("Purchase data not returned correctly")
	}
}

func TestSetMenuState(t *testing.T) {
	//setMenuState("Suppliers")
	state.SetMenuState("Suppliers")
	if state.GetMenuState() != "Suppliers" {
		t.Errorf("expected state 'Suppliers', got '%s'", state.GetMenuState())
	}
}
