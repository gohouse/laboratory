package main

import (
	"github.com/go-redis/redis"
	"log"
)

var (
	rds      *redis.Client
	chatKey  = "z:chat"
	idKey = "a:id"
)

func redisInit(opt *redis.Options) {
	rds = redis.NewClient(&redis.Options{
		Addr:     "redis-17876.c1.ap-southeast-1-1.ec2.cloud.redislabs.com:17876",
		Password: "DOiWgwDu7RaWdb5HIqh4TkmJbjoXQnt886",
	})
}
func main()  {
	redisInit(nil)
	channel := rds.Subscribe("p:test").Channel()
	for {
		select {
		case msg:=<-channel:
			log.Printf("receive msg from channel %s of msg %s \n",msg.Channel, msg.Payload)
		}
	}
}
