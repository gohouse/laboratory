package config

import "errors"

const (
	MYSQL  = "mysql"
	SQLITE = "sqlite"
	JSON = "json"
	TOML = "toml"
)

type ConfigParser struct {
	Driver          string            // 驱动: mysql/sqlite/oracle/mssql/postgres
	Default         string            // 默认 dsn
	QueryLog        bool              // 是否开启sql日志
	SetMaxOpenConns int               // (连接池)最大打开的连接数，默认值为0表示不限制
	SetMaxIdleConns int               // (连接池)闲置的连接数, 默认-1
	Dsns            map[string]string // 数据库链接
}
type IConfig interface {
	Parse(d string) (p *ConfigParser, err error)
}

var ConfigParsers = map[string]IConfig{
	JSON: JsonConfigParser{},
	TOML: JsonConfigParser{},
}

func NewConfigParser(file, fileType string) (*ConfigParser, error) {
	if pr, ok := ConfigParsers[fileType]; ok {
		return pr.Parse(file)
	}
	return &ConfigParser{}, errors.New("不支持的配置类型")
}

// 临时配置, 方便 test 测试
var ConfigFiles = map[string]string{
	JSON: "/Users/fizz/go/src/github.com/gohouse/laboratory/gorose/config/demoConfigFiles/mysql.json",
	TOML: "/Users/fizz/go/src/github.com/gohouse/laboratory/gorose/config/demoConfigFiles/mysql.toml",
}