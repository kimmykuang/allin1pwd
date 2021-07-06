// this package contains example or test codes
package main

import (
	"flag"
	"log"
)

type exampleRunFunc func()

var exampleFlag string

// enabledExamples all examples' entry function registered here
var enabledExamples = map[string]exampleRunFunc{
	"new_window":           NewWindow,
	"tidy_up":              TidyUp,
	"show_another":         ShowAnother,
	"canvas_object":        CanvasObject,
	"container_and_layout": ContainerAndLayout,
	"widgets":              Widgets,
	"canvas_rectangle":     CanvasRectangle,
	"canvas_text":          CanvasText,
	"canvas_line":          CanvasLine,
	"canvas_circle":        CanvasCircle,
	"canvas_image":         CanvasImage,
	"canvas_raster":        CanvasRaster,
	"canvas_gradient":      CanvasGradient,
	"layout_box": LayoutBox,
	"layout_grid":LayoutGrid,
}

func init() {
	flag.StringVar(&exampleFlag, "example", "", "this flag controls which example will execute.")
}

// main run examples as `go run [-tags mobile] *.go -example=[xxx]`
// mobile tags wil run as mobile simulation mode
// example flag always equals to the example file name, eg: -example=new_window, means run new_window.go
func main() {
	flag.Parse()
	if exampleFlag == "" {
		log.Fatalf("example flag cannot be ampty, please run as -example=xxx")
	}
	f, ok := enabledExamples[exampleFlag]
	if !ok {
		log.Fatalf("unsupported example flag.")
	}
	f()
}
