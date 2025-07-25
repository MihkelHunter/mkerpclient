package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type CustomTheme struct{}

var _ fyne.Theme = (*CustomTheme)(nil)

// Colors inspired by Dynamics 365
var (
	white      = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	lightGray  = color.NRGBA{R: 245, G: 245, B: 245, A: 255}
	midGray    = color.NRGBA{R: 160, G: 160, B: 160, A: 255}
	darkGray   = color.NRGBA{R: 51, G: 51, B: 51, A: 255}
	accentBlue = color.NRGBA{R: 0, G: 120, B: 212, A: 255} // Microsoft blue
	hoverBlue  = color.NRGBA{R: 229, G: 243, B: 255, A: 255}
)

func (c *CustomTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNameBackground:
		return lightGray
	case theme.ColorNameButton:
		return hoverBlue
	case theme.ColorNameDisabled:
		return midGray
	case theme.ColorNameForeground:
		return darkGray
	case theme.ColorNamePlaceHolder:
		return midGray
	case theme.ColorNamePrimary:
		return accentBlue
	case theme.ColorNameHover:
		return color.NRGBA{R: 180, G: 210, B: 250, A: 255}
	case theme.ColorNameInputBackground:
		return white
	case theme.ColorNameScrollBar:
		return midGray
	default:
		return theme.DefaultTheme().Color(n, v)
	}
}

func (c *CustomTheme) Font(s fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(s)
}

func (c *CustomTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (c *CustomTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}
