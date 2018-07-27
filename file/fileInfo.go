package main

import (
	"fmt"
	"log"
	"syscall"
	"time"
)

func main() {
	file := "/Users/fizz/Downloads/7.26/1/1530089099204.jpg"
	var st syscall.Stat_t
	if err := syscall.Stat(file, &st); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("mtime: %d\n", st.Mtimespec.Sec)
	fmt.Printf("ctime: %d\n", st.Ctimespec.Sec)

	tm := time.Unix(st.Mtimespec.Sec, 0)

	fmt.Println(tm.Format("2006-01-02 15:04:05"))

	fmt.Println(tm.Format("02/01/2006 15:04:05 PM"))
}
