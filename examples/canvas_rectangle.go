package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

func CanvasRectangle() {
	a := app.New()
	w := a.NewWindow("Rectangle")

	rect := canvas.NewRectangle(color.White)
	rect.FillColor = color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	w.SetContent(rect)
	w.Resize(fyne.NewSize(150, 100))
	w.ShowAndRun()
}
