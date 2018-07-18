package main

import (
	"github.com/go-xorm/xorm"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"github.com/go-xorm/core"
)

func main()  {
	engine, err := xorm.NewEngine("mysql",
		"gcore:gcore@(192.168.200.248:3306)/test?charset=utf8")

	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "")
	engine.SetTableMapper(tbMapper)

	if err!=nil{
		fmt.Println(err)
		return
	}

	type Task struct {
		Id int
		Money string `xorm:"money"`
		PlatformName time.Time `xorm:"platformName"`
	}

	type Area struct {
		Uid int
		Province string `xorm:"province"`
		City string `xorm:"city"`
	}

	//type Tasks struct {
	//	rows []Task
	//}

	//var area Area
	//everyone := make([]Area, 0)
	peveryone := make(map[int]Area)

	engine.Find(&peveryone)


	fmt.Println(peveryone)
}
