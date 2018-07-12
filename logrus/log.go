package main

import (
	"github.com/sirupsen/logrus"
	"time"
)

type Animal struct {
	Name string
	age int
}

func main() {
	//log.SetFormatter(&log.JSONFormatter{})
	a := Animal{"dog", 22}
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:true})
	logrus.WithFields(logrus.Fields{
		"event": "ne",
		"topic": "title",
		"key": "my key",
	}).Info("hello", a)

	log.Error("hello world")
	//for {
		time.Sleep(time.Second)
		log.Printf("i am ok %s", "dock")
	//}
	log.Fatal("kill ")
}
