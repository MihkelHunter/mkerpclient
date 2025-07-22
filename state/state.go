package state

type MenuState struct {
	Current string
}

var currentState = &MenuState{Current: "PurchaseOrders"} //Default state

func GetMenuState() string {
	return currentState.Current
}

func SetMenuState(state string) {
	currentState.Current = state
}
