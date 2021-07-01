package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"time"
)

func ShowAnother() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.SetContent(widget.NewLabel("Hello"))

	go showAnother(a)

	w.ShowAndRun()

}

func showAnother(a fyne.App) {
	time.Sleep(5 * time.Second)

	w := a.NewWindow("Shown later")
	w.SetContent(widget.NewLabel("5 seconds later"))
	w.Resize(fyne.NewSize(200, 200))
	w.Show()

	time.Sleep(2 * time.Second)
	w.Close()
}
