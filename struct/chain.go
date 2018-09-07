package main

import (
	"database/sql"
	"fmt"
)

type Database struct {
	table                string          // table
	fields               []string        // fields
	where                [][]interface{} // where
	order                string          // order
	limit                int             // limit
	offset               int             // offset
	join                 [][]interface{} // join
	distinct             bool            // distinct
	count                string          // count
	sum                  string          // sum
	avg                  string          // avg
	max                  string          // max
	min                  string          // min
	group                string          // group
	having               string          // having
	data                 interface{}     // data
	trans                bool
	LastInsertId         int64           // insert last insert id
	SqlLogs              []string
	LastSql              string
	tx                   *sql.Tx //Dbstruct Database
	beforeParseWhereData [][]interface{}
}

//var db Database

func (db *Database) Fields(str string) *Database {
	db.fields = append(db.fields, str)
	return db
}

func (db *Database) Get() interface{} {
	return db.fields
}


func main()  {
	var dbs *Database
	res := dbs.Fields("a").Get()
	//db = db.Fields("c")
	//res2 := db.Fields("b").Get()
	fmt.Println(res)
	//fmt.Println(res2)
}

