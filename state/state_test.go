package state

import "testing"

func TestSetMenuState(t *testing.T) {
	//setMenuState("Suppliers")
	SetMenuState("Suppliers")
	if GetMenuState() != "Suppliers" {
		t.Errorf("expected state 'Suppliers', got '%s'", GetMenuState())
	}
}
