package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func update_list(list *widget.Box, category string) {
	switch category {
	case "numbers":
		fmt.Println("numbers")
		list.Children = []fyne.CanvasObject{widget.NewLabel("1"),widget.NewLabel("2"),widget.NewLabel("3")}
	case "letters":
		fmt.Println("letters")
		list.Children = []fyne.CanvasObject{widget.NewLabel("A"),widget.NewLabel("B"),widget.NewLabel("C")}
	}
	list.Refresh()
}

func create_dynamic_box() *widget.Box {
	mylist := widget.NewVBox()
	var get_numbers_btn, get_letters_btn *widget.Button
	get_numbers_btn = widget.NewButton("Get numbers", func(){
		update_list(mylist, "numbers")
		get_numbers_btn.Disable()
		get_letters_btn.Enable()
	})
	get_letters_btn = widget.NewButton("Get letters", func(){
		update_list(mylist, "letters")
		get_numbers_btn.Enable()
		get_letters_btn.Disable()
	})

	update_list(mylist, "numbers")
	get_numbers_btn.Disable()
	return widget.NewVBox(
		get_numbers_btn,
		get_letters_btn,
		mylist,
	)
}

func main() {
	app := app.New()
	w := app.NewWindow("Cache issue?")
	var modal widget.PopUp
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Tested on v1.2.3 works fine, but on v1.2.4 it does not!"),
		widget.NewButton("Modal", func() {
			modal = *widget.NewModalPopUp(
				widget.NewVBox(
					create_dynamic_box(),
					widget.NewButton("Close", func(){
						modal.Hide()
					}),
				),
				w.Canvas(),
			)
		}),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	w.Resize(fyne.Size{400, 500})
	w.ShowAndRun()
}