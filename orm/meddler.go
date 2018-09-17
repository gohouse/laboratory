package main

import "github.com/russross/meddler"

type user struct {
	//Id      int64   `gorm:"primary_key"`
	Age     int64   `meddler:"age"`
	Name    string  `meddler:"name"`
	//Website string  `gorm:"website"`
	//Job     string  `gorm:"job"`
	//Money   float64 `gorm:"money"`
	//Created_at time.Time `gorm:"created_at"`
	//Updated_at time.Time `gorm:"updated_at"`
}
func main() {
	var people []*user
	err := meddler.QueryRow(db, &people, "select * from person")
	err := meddler.QueryAll(db, &people, "select * from person")
}
