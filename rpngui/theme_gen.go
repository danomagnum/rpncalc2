// Code generated by fyne-theme-generator

package main

import (
	_ "embed"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

//go:embed FiraCode-Regular.ttf
var firaCode []byte

var fontFiraCodeRegularTtf = &fyne.StaticResource{
	StaticName:    "FiraCode-Regular.ttf",
	StaticContent: firaCode,
}

type myTheme struct{}

func (myTheme) Color(c fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(c, v)
}

func (myTheme) Font(s fyne.TextStyle) fyne.Resource {
	return fontFiraCodeRegularTtf
}

func (myTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (myTheme) Size(s fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(s)
}
