package state

type MenuState struct {
    Current string
}

var currentState = &MenuState{Current: "PurchaseOrders"}

func GetMenuState() string {
    return currentState.Current
}

func SetMenuState(state string) {
    currentState.Current = state
}