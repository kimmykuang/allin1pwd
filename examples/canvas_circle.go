package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

func CanvasCircle() {
	a := app.New()
	w := a.NewWindow("Circle")

	c := canvas.NewCircle(color.White)
	c.StrokeColor = color.Gray{Y: 0x99}
	c.StrokeWidth = 5
	w.SetContent(c)

	w.Resize(fyne.NewSize(100, 100))
	w.ShowAndRun()
}
