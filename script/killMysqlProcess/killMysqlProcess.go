package main

import (
	"github.com/gohouse/gorose"
	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/cron"
	"time"
	"github.com/sirupsen/logrus"
	"github.com/gohouse/laboratory/bootService"
	"fmt"
	"github.com/gohouse/laboratory/utils"
)

var config = map[string]string{ // 定义名为 mysql_dev 的数据库配置
	"host": "180.97.188.201", // 数据库地址
	"username": "idfas",      // 数据库用户名
	"password": "Mysql777",   // 数据库密码
	"port": "3306",           // 端口
	"database": "idfas",      // 链接的数据库名字
	"charset": "utf8",        // 字符集
	"protocol": "tcp",        // 链接协议
	"prefix": "idfa_",        // 表前缀
	"driver": "mysql",        // 数据库驱动(mysql,sqlite,postgres,oracle,mssql)
}

var connection gorose.Connection
var log *logrus.Logger

func init() {
	var err error
	connection, err = gorose.Open(config)
	if err != nil {
		panic(err)
	}
	log = bootService.NewLogrus("/tmp/logrus.log")
}

func main() {
	defer connection.Close()
	fmt.Println("cron running: ", time.Now().Format("2006-01-02 15:04:05"))
	log.Info("cron running: ", time.Now().Format("2006-01-02 15:04:05"))
	// 开始迁移
	//KillMysqlProcess()
	Crontab()
}

func KillMysqlProcess() {
	log.Info("KillMysqlProcess")
	// 查出当前运行进程数
	processListCount,_ := NewDBInstance().
		Query("select count(*) as count from information_schema.processlist where user='idfas'")
	processListCountSleep,_ := NewDBInstance().
		Query("select count(*) as count from information_schema.processlist where user='idfas' and COMMAND='Sleep'")
	fmt.Println(processListCount, processListCountSleep)
	if processListCount[0]["count"].(int64)>30 {
		processList,_ := NewDBInstance().
			Query("select * from information_schema.processlist where user='idfas'")
		fmt.Println(processList)
		if (processListCount[0]["count"].(int64)-processListCountSleep[0]["count"].(int64))>10{
			for _,item := range processList{
				NewDBInstance().Query("kill "+utils.ParseStr(item["ID"]))
			}
		} else {
			for _,item := range processList{
				if item["COMMAND"]=="Sleep"{
					res,err := NewDBInstance().Query("kill "+utils.ParseStr(item["ID"]))
					fmt.Println(res,err)
				}
			}
		}
	}

}

func Crontab() {
	i := 0
	c := cron.New()
	// 每分钟的第 44s 执行
	//spec := "44 */1 * * * ?"
	//每1分钟的第一秒执行一次
	spec2 := "*/30 * * * * ?"
	c.AddFunc(spec2, func() {
		i++
		fmt.Println("cron running: ", time.Now().Format("2006-01-02 15:04:05"), i)
		log.Println("cron running: ", time.Now().Format("2006-01-02 15:04:05"), i)
		KillMysqlProcess()
	})
	c.Start()

	select {}
}

func M(table string) *gorose.Database {
	return NewDBInstance().Table(table)
}

func NewDBInstance() *gorose.Database {
	return connection.GetInstance()
}
