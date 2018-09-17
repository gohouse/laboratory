package builder

import (
	"github.com/gohouse/laboratory/gorose/config"
)

// 检查解析器是否实现了接口
var mysqlBuilder IBuilder = &MysqlBuilder{}
var sqliteBuilder IBuilder = &SqliteBuilder{}

// 注册解析器
var builders = map[string]IBuilder{
	config.MYSQL:  mysqlBuilder,
	config.SQLITE: sqliteBuilder,
}

func BuildQuery(d string) (string, error) {
	return builders[d].BuildQuery()
}

func BuildExecute(d string) (string, error) {
	return builders[d].BuildExecute()
}

//func buildDriver(d string) IBuilder {
//	var bdi IBuilder
//
//	switch strings.ToLower(d) {
//	case driver.MYSQL:
//		bdi = new(MysqlBuilder)
//
//	case driver.SQLITE:
//		bdi = new(SqliteBuilder)
//	}
//
//	return bdi
//}
