package gorose

import (
	"github.com/gohouse/laboratory/gorose/builder"
)

type Database struct {
	connection *Connection
	table      string
}

func (d *Database) Get() string {
	return d.table
}

func (d *Database) Table(arg string) *Database {
	d.table = arg
	return d
}

func (d *Database) BuildSql(b string) (string, error) {
	switch b {

	case "select":
		return builder.BuildQuery(d.connection.driver)

	case "insert", "update", "delete":
		return builder.BuildExecute(d.connection.driver)
	}

	return "", nil
}
