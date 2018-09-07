package builder

import (
	"github.com/gohouse/laboratory/gorose/config"
)

var builders = map[string]IBuilder{
	config.MYSQL:  MysqlBuilder{},
	config.SQLITE: SqliteBuilder{},
}

type IBuilder interface {
	BuildQuery() (string,error)
	BuildExecute() (string,error)
}

func BuildQuery(d string) (string,error) {
	return builders[d].BuildQuery()
}

func BuildExecute(d string) (string,error) {
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
