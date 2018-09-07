package main

import (
	"net/http"
	"log"
	"os"
	"io"
	"fmt"
	"strings"
)

func main() {//已经有静态文件了
	http.HandleFunc("/up",up)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func up(w http.ResponseWriter, r *http.Request)  {
	r.ParseMultipartForm(8 << 20)
	//title := r.ParseFormValue["title"]
	fhs := r.MultipartForm.File["file[]"]
	//options := r.MultipartForm.Value["options[]"]
	//answers := r.MultipartForm.Value["answers[]"]

	l := len(fhs)
	optionDirs := make([]string, l)
	//t := time.Now()
	for i := 0; i < l; i++ {
		file, err := fhs[i].Open()
		if err != nil {
			panic(err)
		}
		filename := fhs[i].Filename
		f, err := os.OpenFile("static/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		io.Copy(f, file)
		optionDirs = append(optionDirs, filename)
	}

	fmt.Println(optionDirs)
	w.Write([]byte(strings.Join(optionDirs,"<br>")))
}
