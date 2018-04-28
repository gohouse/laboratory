package main

import (
	"fmt"
	"github.com/gin-gonic/gin/json"
)

var DbConfig = map[string]interface{}{
	"default":         "mysql_dev", // 默认数据库配置
	"SetMaxOpenConns": 300,         // (连接池)最大打开的连接数，默认值为0表示不限制
	"SetMaxIdleConns": 10,          // (连接池)闲置的连接数, 默认-1
	"mysql_dev": map[string]string{// 定义名为 mysql_dev 的数据库配置
		"host": "192.168.200.248", // 数据库地址
		"username": "gcore",       // 数据库用户名
		"password": "gcore",       // 数据库密码
		"port": "3306",            // 端口
		"database": "test",        // 链接的数据库名字
		"charset": "utf8",         // 字符集
		"protocol": "tcp",         // 链接协议
		"prefix": "",              // 表前缀
		"driver": "mysql",         // 数据库驱动(mysql,sqlite,postgres,oracle,mssql)
	},
}

type Conf struct {
	Default string
	SetMaxOpenConns int
	SetMaxIdleConns int
	Connections map[string]map[string]string
}

func main() {
	var conf = Conf{
		"mysql_dev",
		300,
		10,
		map[string]map[string]string{
			"mysql_dev": map[string]string{// 定义名为 mysql_dev 的数据库配置
				"host": "192.168.200.248", // 数据库地址
				"username": "gcore",       // 数据库用户名
				"password": "gcore",       // 数据库密码
				"port": "3306",            // 端口
				"database": "test",        // 链接的数据库名字
				"charset": "utf8",         // 字符集
				"protocol": "tcp",         // 链接协议
				"prefix": "",              // 表前缀
				"driver": "mysql",         // 数据库驱动(mysql,sqlite,postgres,oracle,mssql)
			},
		},
	}

	res,_ := json.Marshal(conf)
	fmt.Println(string(res))
}
