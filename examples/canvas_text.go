package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

func CanvasText() {
	a := app.New()
	w := a.NewWindow("Text")

	text := canvas.NewText("Text Object", color.White)
	text.Alignment = fyne.TextAlignTrailing
	text.TextStyle = fyne.TextStyle{Italic: true}
	w.SetContent(text)

	w.ShowAndRun()
}
