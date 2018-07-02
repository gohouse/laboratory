package main

import "fmt"

type Article struct{
	Info map[string]interface{}
	Name string
}

func ArticleInit() *Article {
	return &Article{Name: "sdfa"}
}

func main()  {
	art := ArticleInit()
	fmt.Println(art.Name)

	var b = &Article{}
	fmt.Println(b.Name)
}
