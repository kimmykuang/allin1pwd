package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

func CanvasLine() {
	a := app.New()
	w := a.NewWindow("Line")

	line := canvas.NewLine(color.White)
	line.StrokeWidth = 5
	println(line.Position().X, line.Position().Y)
	w.SetContent(line)
	w.Resize(fyne.NewSize(100, 100))
	//line.Move(fyne.NewPos(50, 50))

	w.ShowAndRun()
}
