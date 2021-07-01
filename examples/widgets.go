package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func Widgets() {
	a := app.New()
	w := a.NewWindow("Widgets")

	w.SetContent(widget.NewEntry())
	w.ShowAndRun()
}