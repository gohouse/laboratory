package main

import (
	"fmt"
	"context"
)

type Values struct {
	m map[string]string
}

func (v Values) Get(key string) string {
	return v.m[key]
}

func main() {
	v := Values{map[string]string{
		"1": "one",
		"2": "two",
	}}

	//c := context.Background()
	c2 := context.WithValue(context.Background(), "myvalues", v)

	fmt.Println(c2.Value("myvalues").(Values).Get("2"))
}
