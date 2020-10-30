package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"github.com/gohouse/go4rdm/uitheme"
	"image/color"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")


	width := 500
	height := 400

	window := myWindow
	window.Resize(fyne.NewSize(int(width), int(height)))
	window.SetContent(buildBoxWithCanvas())
	window.ShowAndRun()

	//newCanvas.EndRun()
	//fc.AbsEndRun(w,canvas,width,height)
}

func buildBoxWithCanvas() fyne.CanvasObject {
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 111}
	blue := color.RGBA{44, 77, 232, 255}
	//blue = uitheme.Dark.Info

	gradient := canvas.NewHorizontalGradient(blue,black)
	text := canvas.NewText("test",white)
	text.Alignment = fyne.TextAlignCenter

	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, nil), gradient, text)
}
