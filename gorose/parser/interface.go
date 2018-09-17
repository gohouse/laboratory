package parser

type IParser interface {
	Parse(d string) (p *FileParser, err error)
}

type FileParser struct {
	Driver          string            // 驱动: mysql/sqlite/oracle/mssql/postgres
	Default         string            // 默认 dsn
	QueryLog        bool              // 是否开启sql日志
	SetMaxOpenConns int               // (连接池)最大打开的连接数，默认值为0表示不限制
	SetMaxIdleConns int               // (连接池)闲置的连接数, 默认-1
	Dsns            map[string]string // 数据库链接
}
