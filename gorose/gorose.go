package gorose

import (
	"errors"
	"github.com/gohouse/laboratory/gorose/parser"
	"encoding/json"
	"github.com/gohouse/laboratory/gorose/util"
	"database/sql"
)

type Connection struct {
	dbConf *parser.FileParser
	dbs map[string]*sql.DB
}

func Open(fileOrDriverType,dsnOrFile string) (*Connection, error) {
	var c *Connection
	return c.connectionReal(fileOrDriverType,dsnOrFile)
}

func (c *Connection) AddConnection(fileOrDriverTypeOrDsnKey string,dsnOrFile ...string) (*Connection, error) {
	switch len(dsnOrFile){
	case 0:
		return c.connectionReal(c.dbConf.Default,c.dbConf.Dsns[fileOrDriverTypeOrDsnKey])
	case 1:
		return c.connectionReal(fileOrDriverTypeOrDsnKey,dsnOrFile[0])
	default:

	}
}

func (c *Connection) connectionReal(fileOrDriverType,dsnOrFile string) (*Connection, error) {
	var dbConf *parser.FileParser
	//var dsn string
	var err error
	//var db *sql.DB
	// 解析配置
	if dbConf,err = c.parse(fileOrDriverType,dsnOrFile); err!=nil {
		return c,err
	}

	// 驱动数据库
	c.dbs[dbConf.Default],err = c.bootDb(dbConf)

	// 返回链接
	return c, nil
}

func (c *Connection) parse(args ...interface{}) (dbConf *parser.FileParser, err error) {
	argLen := len(args)

	switch argLen{
	case 0:
		return dbConf,errors.New("Open参数不能为空")
	case 1:
		switch args[0].(type) {
		case map[string]interface{}:
			jsonMarshal,err := json.Marshal(args[0])
			if err!=nil {
				return
			}
			err = json.Unmarshal(jsonMarshal, &dbConf)
			if err!=nil {
				return
			}
		default:
			return dbConf,errors.New("Open参数为一个时, 需要传递map[string]interface{}格式的数据库配置")
		}
	case 2:
		switch args[0].(type) {
		case map[string]interface{}:
			return c.parse(args[0])
		case string:
			// 配置文件, 读取配置文件
			if !util.FileExists(args[0].(string)) {
				return dbConf,errors.New("文件不存在:第一个参数为字符串时,需要传入文件路径")
			}
			switch args[1].(type) {
			case string:
				// 解析文件
				dbConf,err = parser.NewFileParser(args[0].(string), args[1].(string))
				if err!=nil{
					return
				}
			default:
				return dbConf,errors.New("第二个参数只能为string")
			}
		}
	}
	return
}

//func (c *Connection) drive(d string) (dsn string, err error) {
//	return driver.NewDriver(d)
//}



// boot sql driver
func (c *Connection) bootDb(dbConf *parser.FileParser) (db *sql.DB, err error) {
	//db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/test?charset=utf8")
	// 开始驱动
	db, err = sql.Open(dbConf.Driver, dbConf.Dsns[dbConf.Default])
	if err != nil {
		return
	}
	if dbConf.SetMaxOpenConns>0 {
		db.SetMaxOpenConns(dbConf.SetMaxOpenConns)
	}
	if dbConf.SetMaxIdleConns>0 {
		db.SetMaxIdleConns(dbConf.SetMaxIdleConns)
	}

	// 检查是否可以ping通
	err = db.Ping()

	return
}

func (c *Connection) NewDB() *Database {
	return &Database{connection: c}
}

func (c *Connection) Table(arg string) *Database {
	return c.NewDB().Table(arg)
}
