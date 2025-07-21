package main

import (
    "testing"
    "fyne.io/fyne/v2/container"
)

func TestShowSuppliersSetsState(t *testing.T) {
    mainContent := container.NewVBox()
    currentState.CurrentMenuState = "PurchaseOrders"

    showSuppliers(mainContent)

    if currentState.CurrentMenuState != "Suppliers" {
        t.Errorf("expected state 'Suppliers', got '%s'", currentState.CurrentMenuState)
    }
}

func TestGetPurchaseData(t *testing.T) {
    data, title := getPurchaseData("Suppliers")
    if title != "Supplier Data" || data[1][1] != "Supplier A" {
        t.Error("Supplier data not returned correctly")
    }
    data, title = getPurchaseData("PurchaseOrders")
    if title != "Purchase Data" || data[1][0] != "A100" {
        t.Error("Purchase data not returned correctly")
    }
}