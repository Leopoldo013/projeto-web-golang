// @Title
// @Description
// @Author
// @Update

package ui

import (
	"Leopoldo013/pixl/apptype"
	"Leopoldo013/pixl/pxcanvas"
	"Leopoldo013/pixl/swatch"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
