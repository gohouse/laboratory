package main

import (
	"log"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	log.Println("https://localhost/ping")
	//log.Fatal(autotls.Run(r, "localhost"))
	log.Fatal(autotls.Run(r, "ssl.mmsex.ml"))
}