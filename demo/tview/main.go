// Demo code for the Table primitive.
package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func main() {

	var data = [][]string{
		{"d", "d", "d"},
		{"x", "x"},
		{"d", "d", "d", "d"},
		{"x"},
	}
	ShowLuzhuInTerminal(data)
}
func ShowLuzhuInTerminal(data [][]string) {
	app := tview.NewApplication()
	table := tview.NewTable()
	for c := 0; c < len(data)+5; c++ {
		if c >= len(data) {
			table.SetCell(0, c, tview.NewTableCell(" "))
		} else {
			for r := 0; r < len(data[c]); r++ {
				table.SetCell(r, c,
					tview.NewTableCell(fmt.Sprintf("-%s", data[c][r])).
						//SetTextColor(color).
						SetAlign(tview.AlignCenter))
			}
		}
	}
	table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
		if key == tcell.KeyEnter {
			table.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row int, column int) {
		table.GetCell(row, column).SetTextColor(tcell.ColorRed)
		table.SetSelectable(false, false)
	})
	if err := app.SetRoot(table, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
