package driver

import (
	"errors"
	"github.com/gohouse/laboratory/gorose/config"
	"database/sql"
)

// 检查解析器是否实现了接口
var mysqlDriver IDriver = &MysqlDriver{}
var sqliteDriver IDriver = &SqliteDriver{}

// 注册解析器
var drivers = map[string]IDriver{
	config.MYSQL:  mysqlDriver,
	config.SQLITE: sqliteDriver,
}

func NewDriver(d string) (dsn string, err error) {
	var ok bool
	var dr IDriver
	if dr, ok = drivers[d]; !ok {
		return dsn, errors.New("no driver matched")
	}
	dsn,err = dr.GetDsn(d)
	if err!=nil {
		return
	}
}
