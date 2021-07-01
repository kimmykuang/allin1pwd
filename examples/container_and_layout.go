package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"image/color"
)

func ContainerAndLayout() {
	a := app.New()
	w := a.NewWindow("Container")

	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText("There", green)
	text2.Move(fyne.NewPos(20, 20))
	//content := container.NewWithoutLayout(text1, text2)
	content := container.New(layout.NewGridLayout(2), text1, text2)

	w.SetContent(content)
	w.ShowAndRun()
}
