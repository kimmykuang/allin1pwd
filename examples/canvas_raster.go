package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

func CanvasRaster() {
	a := app.New()
	w := a.NewWindow("Raster")

	//raster := canvas.NewRasterWithPixels(func (_, _, w, h int) color.Color {
	//	return color.RGBA{
	//		R: uint8(rand.Intn(255)),
	//		G: uint8(rand.Intn(255)),
	//		B: uint8(rand.Intn(255)),
	//		A: 0xff,
	//	}
	//})
	raster := canvas.NewRasterFromImage(canvas.NewImageFromResource(theme.FyneLogo()).Image)
	w.SetContent(raster)
	w.Resize(fyne.NewSize(100, 100))
	w.ShowAndRun()
}
