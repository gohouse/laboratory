package main

import "github.com/go-xorm/xorm"

var engine *xorm.Engine
func init()  {
	engine, _ = xorm.NewEngine(driverName, dataSourceName)
}
func main() {
	var users []User
	err := engine.Where("name = ?", name).And("age > 10").Limit(10, 0).Find(&users)
	// SELECT * FROM user WHERE name = ? AND age > 10 limit 10 offset 0

	type Detail struct {
		Id int64
		UserId int64 `xorm:"index"`
	}

	type UserDetail struct {
		User `xorm:"extends"`
		Detail `xorm:"extends"`
	}

	var users []UserDetail
	err := engine.Table("user").Select("user.*, detail.*").
		Join("INNER", "detail", "detail.user_id = user.id").
		Where("user.name = ?", name).Limit(10, 0).
		Find(&users)

}
