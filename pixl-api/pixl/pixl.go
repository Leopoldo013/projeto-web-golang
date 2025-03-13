// @Title
// @Description
// @Author
// @Update

package main

import (
	"Leopoldo013/pixl/apptype"
	"Leopoldo013/pixl/pxcanvas"
	"Leopoldo013/pixl/swatch"
	"Leopoldo013/pixl/ui"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	pixlApp := app.New()
	pixlWindow := pixlApp.NewWindow("pixl")
	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	pixlCanvasConfig := apptype.PxCanvasConfic{
		DrawingArea:  fyne.NewSize(800, 800),
		CanvasOffset: fyne.NewPos(0, 0),
		PxRow:        50,
		PxCols:       50,
		PxSize:       5,
	}

	pixlCanvas := pxcanvas.NewPxCanvas(&state, pixlCanvasConfig)

	appInit := ui.AppInit{
		PixlCanvas: pixlCanvas,
		PixlWindow: pixlWindow,
		State:      &state,
		Swatches:   make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PixlWindow.ShowAndRun()
}
