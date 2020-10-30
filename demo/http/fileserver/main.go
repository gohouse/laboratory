package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer()
}

func fileServer()  {
	//文件浏览
	http.Handle("/", http.FileServer(http.Dir("./")))
	log.Println("http://localhost")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
