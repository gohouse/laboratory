package main

import (
	"fmt"
	"github.com/gohouse/gorose"
)

type db struct {
	Table string
}
func (d *db) find(id interface{}) interface{} {
	return d.Table
}

type userModel struct {
	db
	Table string
}

func UserModel() *userModel {
	return &userModel{db:db{Table:"test"}}
}

func main() {
	//fmt.Println(UserModel().find(2))

	var (
		Um = &gorose.Model{Table:"test"}
	)
	//
	fmt.Println(Um.Test())
}
