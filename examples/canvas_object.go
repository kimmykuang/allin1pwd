package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"image/color"
	"time"
)

func CanvasObject() {
	a := app.New()
	w := a.NewWindow("Canvas")
	c := w.Canvas()

	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text := canvas.NewText("Text", green)
	text.TextStyle.Bold = true
	c.SetContent(text)

	go changeContent(c)

	w.Resize(fyne.NewSize(100, 100))
	w.ShowAndRun()
}

func changeContent(c fyne.Canvas) {
	time.Sleep(2 * time.Second)

	blue := color.NRGBA{R: 0, G: 0, B: 100, A: 255}
	c.SetContent(canvas.NewRectangle(blue))

	time.Sleep(2 * time.Second)
	c.SetContent(canvas.NewLine(color.Gray{Y: 180}))

	time.Sleep(2 * time.Second)
	red := color.NRGBA{R: 0xff, G: 0x33, B: 0x33, A: 0xff}
	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 4
	circle.StrokeColor = red
	c.SetContent(circle)

	time.Sleep(2 * time.Second)
	c.SetContent(canvas.NewImageFromResource(theme.FyneLogo()))
}
