package main

import (
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

func main() {
	i := 0
	c := cron.New()
	// 每分钟的第 44s 执行
	//spec := "44 */1 * * * ?"
	//每3分钟执行一次
	spec2 := "1 */3 * * * ?"
	c.AddFunc(spec2, func() {
		i++
		logrus.Info("cron running:, ", i)
	})
	c.Start()

	select{}
}
