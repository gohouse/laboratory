package main

import "fmt"

type User struct {
	Name string
	Age int
}

func main() {

	var b *User = UserInit()
	fmt.Println(b)
	fmt.Println(getInfo())
}

func UserInit() *User {
	//var a *User
	//a.Name = "fizz"
	//a.Age = 18

	return  &User{
		"fizz",
		18,
	}
}

func getInfo() interface{} {
	return GetUserAge()
}

func GetUserName(u *User) string {
	return u.Name
}

func GetUserAge(u *User) int {
	return u.Age
}
