package main

import (
	"github.com/gohouse/gorose"
	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/cron"
	"time"
	"github.com/BurntSushi/toml"
	"flag"
	"github.com/gohouse/zhuan8/helper"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/gohouse/laboratory/bootService"
	"fmt"
	"github.com/gohouse/gorose/utils"
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
	//ReleaseTimeoutTask()
	Crontab()
}

func parseFlag(key string) string {
	return *flag.String("conf",
		"src/github.com/gohouse/laboratory/script/db.toml", "配置文件地址")
}

func decodeDbToml() {
	var cg conf
	var cpath string = parseFlag("conf")

	if _, err := toml.DecodeFile(cpath, &cg); err != nil {
		log.Fatal(err)
	}
}

type mysql_items struct {
	Host     string `toml:"host"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Port     string `toml:"port"`
	Database string `toml:"database"`
	Charset  string `toml:"charset"`
	Protocol string `toml:"protocol"`
	Prefix   string `toml:"prefix"`
	Driver   string `toml:"driver"`
}
type conf struct {
	Default         string
	SetMaxOpenConns int
	SetMaxIdleConns int
	Connections     map[string]mysql_items
}

func ReleaseTimeoutTask() {
	log.Info("ReleaseTimeoutTask")
	// 查出当天超时数据
	var datetime = helper.GetDate()
	t := M("user_task").Where("status", 1).Where("created_at", ">=", datetime.TodayStart).
		Where("created_at", "<", time.Now().Add(-time.Minute * 15).Format("2006-01-02 15:04:05"))
	//fmt.Println(time.Now().Format("2006-01-02 03:04:05"))
	//res222, _ := t.First()
	//fmt.Println(t.LastSql)
	//fmt.Println(res222)
	//os.Exit(1)
	t.Chunk(100, func(datas []map[string]interface{}) {
		// 开始处理
		if len(datas) > 0 {
			// 释放开始任务
			for _, item := range datas {
				if len(item) == 0 {
					break
				}
				db := NewDBInstance()
				res, err := db.Transaction(func() error {
					// 标记超时
					res, err := db.Table("user_task").Data(map[string]interface{}{
						"status": 3,
					}).Where("id", item["id"]).Update()
					if err != nil {
						return err
					}
					if res == 0 {
						return errors.New("标记超时失败")
					}

					// 任务剩余数量+1
					res2, err2 := db.Table("tasks").Where("id", item["task_id"]).Increment("task_num")
					//fmt.Println(db.LastSql)
					if err2 != nil {
						return err2
					}
					if res2 == 0 {
						return errors.New("剩余数量返还失败")
					}

					// 状态改为进行中
					db.Table("tasks").Where("id", item["task_id"]).Data(map[string]interface{}{
						"status":1,
					}).Update()

					return nil
				})
				//fmt.Println(item["task_name"], res, err, time.Now().Format("2006-01-02 15:04:05"))
				M("task_timeout_log").Data(map[string]interface{}{
					"task_name":   item["task_name"],
					"appstore_id": item["appstore_id"],
					"mobile":      item["mobile"],
					"idfa":        item["idfa"],
					"task_id":     item["task_id"],
					"result":      utils.If(res == true, true, fmt.Sprint(res, err)),
				}).Insert()
			}
		} else {
			//fmt.Println("无任务: ", time.Now().Format("2006-01-02 15:04:05"))
		}
	})
}

func Crontab() {
	i := 0
	c := cron.New()
	// 每分钟的第 44s 执行
	//spec := "44 */1 * * * ?"
	//每1分钟的第一秒执行一次
	spec2 := "1 */3 * * * ?"
	c.AddFunc(spec2, func() {
		i++
		fmt.Println("cron running: ", time.Now().Format("2006-01-02 15:04:05"), i)
		log.Println("cron running: ", time.Now().Format("2006-01-02 15:04:05"), i)
		ReleaseTimeoutTask()
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

func MapValueNilToNull(res map[string]interface{}) map[string]interface{} {
	if len(res) > 0 {
		for k, v := range res {
			switch v.(type) {
			case nil:
				res[k] = "null"
			}
		}
	}
	return res
}
