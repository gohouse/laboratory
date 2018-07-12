package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	var c *gin.Context

	c.Set("bbb", "234324")

	k,v := c.Get("bbb")

	fmt.Println(k, v)
}
