package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	f := app.New()
	window := f.NewWindow("title")

	tab1 := widget.NewTabItemWithIcon("text", theme.HomeIcon(), widget.NewLabel("text"))
	container := widget.NewTabContainer(
		tab1,
		widget.NewTabItemWithIcon("text2", theme.SearchIcon(), widget.NewLabel("text2")),
		)

	button := widget.NewButton("aaa", func() {
		container.SelectTabIndex(1)
	})
	button2 := widget.NewButton("bbb", func() {
		container.SelectTabIndex(0)
	})
	button3 := widget.NewButton("ccc", func() {
		container.SelectTab(tab1)
	})

	window.SetContent(widget.NewVBox(button,button2,button3,container))
	window.ShowAndRun()
}
