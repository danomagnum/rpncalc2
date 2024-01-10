package main

import (
	_ "embed"

	fyne "fyne.io/fyne/v2"
)

//go:embed FiraCode-Regular.ttf
var firaCode []byte

var fontFiraCodeRegularTtf = &fyne.StaticResource{
	StaticName:    "FiraCode-Regular.ttf",
	StaticContent: firaCode,
}
