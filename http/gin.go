package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main(){
	gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		go func() {
			//s := time.Now().Second()
			test()
		}()
		c.String(http.StatusOK,"success3")
	})

	router.Run(":8005")
}

func test()  {
	time.Sleep(time.Duration(time.Now().Second()%30)+1 * time.Second)
	log.Println("safds")
}
