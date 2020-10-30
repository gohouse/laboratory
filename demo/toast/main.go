package main

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/gen2brain/dlgs"
)

func main() {
	dialog3()
}
func main3() {
	var err error
	err = beeep.Alert("Title", "Message body", "https://raw.githubusercontent.com/gen2brain/beeep/master/assets/warning.png")
	if err != nil {
		panic(err)
	}
	//err = beeep.Notify("Title", "Message body", "assets/information.png")
	//if err != nil {
	//	panic(err)
	//}
	//err = beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	//if err != nil {
	//	panic(err)
	//}
}

func dialog()  {
	item, _, err := dlgs.List("List", "Select item from list:", []string{"Bug", "New Feature", "Improvement"})
	if err != nil {
		panic(err)
	}
	fmt.Println(item)
}
func dialog2()  {
	passwd, _, err := dlgs.Password("Password", "Enter your API key:")
	if err != nil {
		panic(err)
	}
	fmt.Println(passwd)
}
func dialog3()  {
	//dlgs.Date("test","for test date", time.Now())
	//dlgs.File("test","for test date", false)
	////dlgs.Error("aa","bb")
	//entry, b, err := dlgs.Entry("aa", "bb", "cc")
	//fmt.Println(entry, b, err)

	yes, err := dlgs.Question("Question", "Are you sure you want to format this media?", false)
	if err != nil {
		panic(err)
	}
	fmt.Println(yes)
}