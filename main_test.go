package main

import (
	"mkerpclient/ui"
	"testing"

	"fyne.io/fyne/v2/container"
)

func TestSidebarBuild(t *testing.T) {
	mainContent := container.NewStack()
	sidebar := ui.BuildSidebar(mainContent)

	if sidebar == nil {
		t.Fatal("Sidebar should not be nil")
	}

	if len(sidebar.Objects) == 0 {
		t.Error("Sidebar should contain UI elements")
	}
}
