package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	var a = 0Xf3abcdef
	var b = 0o12
	var c = 0B1100
	fmt.Printf("%x,%x,%x",a,b,c)
	return
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "123456"})
	client.Set("aid", 1000000000000,0).Val()
	for i:=0;i<1000;i++{
		client.Incr("aid").Val()
	}
}
