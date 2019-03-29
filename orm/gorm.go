package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

func main() {
	db, err := gorm.Open("mysql", "gcore:gcore@tcp(192.168.200.248:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println(err,11111)
		return
	}
	db.AutoMigrate()
	fmt.Println(db,222222)
	defer db.Close()
	aaa,err2 := db.DB().Query("select * from users limit 1")
	fmt.Println(aaa,err2)

	type User struct {
		//Id      int64   `gorm:"primary_key"`
		Age     int64   `gorm:"age"`
		Name    string  `gorm:"name"`
		//Website string  `gorm:"website"`
		//Job     string  `gorm:"job"`
		//Money   float64 `gorm:"money"`
		//Created_at time.Time `gorm:"created_at"`
		//Updated_at time.Time `gorm:"updated_at"`
	}

	var user User

	res := db.First(&user, 47)
	fmt.Println(res)
}
