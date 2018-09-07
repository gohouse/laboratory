package gorose

import "github.com/gohouse/laboratory/dp/driver"

type Connection struct {
	driver string
	dsn    string
}

func Open(dr string) *Connection {
	var conn *Connection
	return conn.drive(dr)
}

func (c *Connection) drive(dr string) *Connection {
	dsns := driver.NewDriver(dr)
	return &Connection{driver: dr, dsn: dsns}
}

func (c *Connection) NewDB() *Database {
	return &Database{connection:c}
}

func (c *Connection) Table(arg string) *Database {
	return c.NewDB().Table(arg)
}
