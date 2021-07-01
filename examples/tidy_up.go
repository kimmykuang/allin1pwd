package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"log"
)

func TidyUp() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.SetContent(widget.NewLabel("Hello"))

	w.Show()
	a.Run()
	tidyUp()
}

func tidyUp() {
	log.Print("tidy up after app quit.")
}
