package main

import (
	"context"
	"fmt"
)

func main() {
	InitValues()
	var s *MyContext
	fmt.Println(s.C)
}

type MyValues struct {
	m map[string]string
}

func (v MyValues) Get(key string) string {
	return v.m[key]
}

type MyContext struct {
	V MyValues
	C context.Context
}

func InitValues()  {
	var v *MyValues
	v.m = map[string]string{
		"a":"aa",
		"b":"bb",
	}

	var c *MyContext
	c.C = context.WithValue(context.Background(), "myvalues", v)
}

func (m *MyContext) GetItem() string {
	var ctx context.Context

	return ctx.Value("myvalues").(MyValues).Get("a")
}
