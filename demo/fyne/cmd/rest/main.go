package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")
	button := widget.NewButton("Save1", func() {

	})

	content := widget.NewVBox(input,button)
	content.Append(widget.NewButton("Save", func() {
		content.Append(widget.NewLabel(input.Text))
	}))

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
