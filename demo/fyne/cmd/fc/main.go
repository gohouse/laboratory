package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"github.com/ajstarks/fc"
	"image/color"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")

	//input := widget.NewEntry()
	//input.SetPlaceHolder("Enter text...")
	//button := widget.NewButton("Save1", func() {
	//
	//})
	//
	//content := widget.NewVBox(input,button)
	//content.Append(widget.NewButton("Save", func() {
	//	content.Append(widget.NewLabel(input.Text))
	//}))


	width := 500
	height := 400

	white := color.RGBA{255, 255, 255, 255}
	blue := color.RGBA{44, 77, 232, 255}
	midx := width / 2
	iy := height / 5
	ty := 3 * height / 4
	newCanvas := fc.NewCanvas("abcd", width, height)
	newCanvas.Window = myWindow
	canvas := newCanvas.Container
	//w, canvas := fc.AbsStart("hello", width, height)
	fc.AbsCircle(canvas, midx, height, midx, blue)
	fc.AbsTextMid(canvas, midx, ty, "hello, world", width/10, white)
	fc.AbsImage(canvas, midx, iy, 200, 200, "earth.jpg")
	//content.Append(canvas)

	//myWindow.SetContent(canvas)
	//myWindow.ShowAndRun()
	window := myWindow
	window.Resize(fyne.NewSize(int(width), int(height)))
	//window.SetFixedSize(true)
	//window.SetPadded(false)
	window.SetContent(canvas)
	window.ShowAndRun()

	//newCanvas.EndRun()
	//fc.AbsEndRun(w,canvas,width,height)
}
