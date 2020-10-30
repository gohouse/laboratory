package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
	"time"
)

var t time.Duration
var content string

func init() {
	flag.DurationVar(&t, "t", 600, "提醒周期, 默认600, 单位是秒")
	flag.StringVar(&content, "content", "记得喝水活动一下", "提醒内容")
}
func main() {
	flag.Parse()
	WithTicker(t*time.Second, notice)
}

func notice() {
	Notice()
}

func Notice() error {
	var shellfile string
	var cmd []string
	switch runtime.GOOS {
	case "windows":
		//home,_:=user.Current()
		//shellfile = fmt.Sprintf("c:/Users/notice_file.bat")
		//content := `mshta vbscript:msgbox("do sports now",64,"notice")(window.close())`
		shellfile = "c:\\notice_file.vbs"
		content := `CreateObject("SAPI.SpVoice").Speak "notice for rest"
		createobject("wscript.shell").popup "notice for rest",3,"notice",4096+64 '3秒后自bai动关闭`
		cmd = append(cmd, `cmd`, `/c`)
		// content 写入文件
		ioutil.WriteFile(shellfile, []byte(content), 777)
	default:
		home, _ := user.Current()
		shellfile = fmt.Sprintf("%s/tmp/notice_file.sh", home.HomeDir)
		content := `#!/bin/env sh
title="notice"
content="`+content+`"
subtitle="just a notice"
sound="Pon"
cmd=$(printf 'display notification "%s" with title "%s" subtitle "%s" sound name "%s"' "$content" "$title" "$subtitle" "$sound")
osascript -e "$cmd"
say -v Ting-ting $content`
		cmd = append(cmd, `sh`)
		// content 写入文件
		ioutil.WriteFile(shellfile, []byte(content), 777)
	}
	log.Println("notice of:", content)
	return exec.Command(strings.Join(cmd, " "), shellfile).Run()
	//return exec.Command("sh", shellfile).Start()
}

func WithTicker(duration time.Duration, fn func()) {
	fn()
	ticker := time.NewTicker(duration)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fn()
		}
	}
}
