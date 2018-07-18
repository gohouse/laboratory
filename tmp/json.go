package main

import (
	"fmt"
	"encoding/json"
	"reflect"
)

//type Confer struct {
//	Default string	`json:"Default"`
//	SetMaxOpenConns int	`json:"SetMaxOpenConns"`
//	SetMaxIdleConns int	`json:"SetMaxIdleConns"`
//	Connections map[string]map[string]string	`json:"Connections"`
//}
type Configer struct {
	Default string
	SetMaxOpenConns int
	SetMaxIdleConns int
	Connections map[string]map[string]string
}

func main() {
	jsons := `{"Default":"mysql_dev","SetMaxOpenConns":300,"SetMaxIdleConns":10,"Connections":{"mysql_dev":{"charset":"utf8","database":"test","driver":"mysql","host":"192.168.200.248","password":"gcore","port":"3306","prefix":"","protocol":"tcp","username":"gcore"}}}`

	var conf Configer

	var confReal = map[string]interface{}{}
	json.Unmarshal([]byte(jsons), &conf)

	confReal["Default"] = conf.Default
	confReal["SetMaxOpenConns"] = conf.SetMaxOpenConns
	confReal["SetMaxIdleConns"] = conf.SetMaxIdleConns
	confReal["Connections"] = conf.Connections
	//for _,data := range maps {
	//	fmt.Println(data)
	//}
	fmt.Println(reflect.TypeOf(conf))

	fmt.Println(conf.Connections["mysql_dev"]["host"])

	parseConfig(confReal)
}

func parseConfig(args interface{})  {
	fmt.Println(args)
	res,_ := json.Marshal(args)

	fmt.Println(string(res))
}