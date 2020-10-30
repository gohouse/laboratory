package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ThomasRooney/gexpect"
)

func main() {
	cmd := "rsync -avz -e 'ssh -p 22' --delete --timeout=60 --bwlimit=3000 /home/video root@192.168.1.208:/root/"
	pwd := "yourpassword"

	child, err := gexpect.Spawn(cmd)
	if err != nil {
		log.Fatal("Spawn cmd error ", err)
	}

	if err := child.ExpectTimeout("password: ", 10*time.Second); err != nil {
		log.Fatal("Expect timieout error ", err)
	}

	if err := child.SendLine(pwd); err != nil {
		log.Fatal("SendLine password error ", err)
	}

	if err := child.Wait(); err != nil {
		log.Fatal("Wait error: ", err)
	}

	fmt.Println("Success")
}