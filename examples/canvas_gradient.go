package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

func CanvasGradient() {
	a := app.New()
	w := a.NewWindow("Gradient")

	//gradient := canvas.NewHorizontalGradient(color.White, color.Transparent)
	gradient := canvas.NewRadialGradient(color.White, color.Transparent)
	w.SetContent(gradient)
	w.Resize(fyne.NewSize(100, 100))
	w.ShowAndRun()
}
