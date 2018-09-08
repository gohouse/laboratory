package driver

import (
	"errors"
	"github.com/gohouse/laboratory/gorose/config"
)

var drivers = map[string]IDriver{
	config.MYSQL:  MysqlDriver{},
	config.SQLITE: SqliteDriver{},
}

func NewDriver(d string) (string, error) {
	if dr, ok := drivers[d]; ok {
		return dr.Drive(d)
	}
	return "",errors.New("no driver matched")
}

//func Drive(d string) string {
//	var dri IDriver
//
//	switch strings.ToLower(d) {
//
//	case MYSQL:
//		dri = new(MysqlDriver)
//
//	case SQLITE:
//		dri = new(SqliteDriver)
//	}
//
//	return dri.Drive()
//}
