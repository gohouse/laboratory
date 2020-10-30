package chatcli

import (
	"github.com/atotto/clipboard"
	"github.com/gdamore/tcell"
	"github.com/gen2brain/dlgs"
	"github.com/rivo/tview"
)

func ui(c IChat) {
	app := tview.NewApplication()
	var leftFlex = tview.NewFlex().SetDirection(tview.FlexRow)
	var midFlex = tview.NewFlex().SetDirection(tview.FlexRow)
	var rightFlex = tview.NewFlex().SetDirection(tview.FlexRow)

	// 构建左边
	leftTitle := tview.NewTextView().SetTextAlign(tview.AlignLeft).SetText("群成员")
	leftList := tview.NewList()
	//leftList.
	//	AddItem("wukong", "abcd\ndef", '1', nil).
	//	AddItem("houzi", "abcd", '2', nil).
	//	AddItem("lucian", "", '3', nil).
	//	AddItem("roths", "", '4', nil)
	leftFlex.AddItem(leftTitle, 2, 0, false).
		AddItem(leftList, 0, 1, false)

	// 构建中间=================start
	// 中上展示消息部分===========start
	var middlePage = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetScrollable(true).
		SetChangedFunc(func() {
			app.Draw()
		})

	middlePage.SetBorder(true).SetTitle("chat").SetBorderPadding(0, 0, 1, 0).SetBorderAttributes(tcell.AttrDim)
	// 先展示一下旧消息
	c.ShowHistory(middlePage)
	// 滚动到最底部
	middlePage.ScrollToEnd().SetRegions(true)
	// 不断获取新消息
	go c.ReceiveMsg(middlePage)
	// 中上展示消息部分===========end
	// 中下发送消息部分===========start
	var middlebottom = tview.NewInputField().SetPlaceholder(" please type input text")
	middlebottom.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			c.SendMsgString(middlebottom.GetText())
			middlebottom.SetText("")
		}
	})
	middlebottom.SetFieldTextColor(tcell.ColorYellow).SetLabelWidth(1).
		SetFieldBackgroundColor(tcell.ColorBlack).SetBorder(true).SetTitle("chat").SetBorderAttributes(tcell.AttrDim)
	midFlex.
		AddItem(middlePage, 0, 1, false).
		AddItem(middlebottom, 3, 0, false)
	// 中下发送消息部分===========end
	// 构建中间=================end

	// 构建右边
	rightTitle := tview.NewTextView().SetTextAlign(tview.AlignLeft).SetText("高级操作")
	//rightList := tview.NewPages()
	//rightList.
	//	AddItem("聊天", "hello world", 'a', listHandleFunc(rightList, middlebottom)).
	//	AddItem("股票", "/gp sh000001", 'b', listHandleFunc(rightList, middlebottom)).
	//	AddItem("开车", "/kfc", 'c', listHandleFunc(rightList, middlebottom)).
	//	AddItem("娱乐", "/joke", 'd', listHandleFunc(rightList, middlebottom))
	button := tview.NewButton("复制内容").SetSelectedFunc(func() {
		var tmpArr []string
		var history = c.History()
		if len(history) > 0 {
			for i := len(history) - 1; i >= len(history)-11; i-- {
				if i < 0 {
					break
				}
				tmpArr = append(tmpArr, history[i].Msg)
			}
		}
		//list, _, _ := dlgs.List("复制内容", "选择要复制的内容,点击确认即可复制到粘贴板", []string{"a","b"})
		list, _, _ := dlgs.List("复制内容", "选择内容,点击确认复制到粘贴板", tmpArr)
		clipboard.WriteAll(list)
	})
	button.SetRect(0, 0, 0, 0)
	rightFlex.AddItem(rightTitle, 2, 0, false).
		AddItem(button, 3, 0, false)

	// 构建整体
	var flex = tview.NewFlex().
		AddItem(leftFlex, 16, 1, false).
		AddItem(midFlex, 0, 3, false).
		AddItem(rightFlex, 16, 1, false)

	if err := app.SetRoot(flex, true).EnableMouse(true).SetFocus(middlebottom).Run(); err != nil {
		panic(err)
	}
}

func listHandleFunc(l *tview.List, i *tview.InputField) func() {
	return func() {
		text, secondary := l.GetItemText(l.GetCurrentItem())
		i.SetText(secondary)
		i.SetTitle(text)
	}
}
