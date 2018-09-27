package main

import (
	"github.com/gohouse/gorose"
	"sync"
	"fmt"
)

var m *Manager
var once sync.Once

func GetInstance() *Manager {
	once.Do(func() {
		m = &Manager{}
	})
	return m
}

type Manager struct{}

func (p Manager) Manage() {
	fmt.Println("manage...")
}

func main() {

	// DB Config.(Recommend to use configuration file to import)
	var dbConfig = &gorose.DbConfigSingle{
		Driver:          "mysql", // 驱动: mysql/sqlite/oracle/mssql/postgres
		EnableQueryLog:  true,    // 是否开启sql日志
		SetMaxOpenConns: 0,       // (连接池)最大打开的连接数，默认值为0表示不限制
		SetMaxIdleConns: 0,       // (连接池)闲置的连接数, 默认-1
		Prefix:          "",      // 表前缀
		Dsn:             "root:root@tcp(localhost:3306)/test?charset=utf8", // 数据库链接
	}
}
