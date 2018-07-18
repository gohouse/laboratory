package main

import (
	"github.com/knadh/jsonconfig"
)

func main() {
	// setup the structure
	config := struct {
		Url string `json:"url"`

		Methods []string `json:"methods"`

		AlwaysLoad  bool `json:"always_load"`

		Module struct{
			Name string `json:"name"`
			Route string `json:"route"`
			Port int `json:"port"`
		} `json:"module"`
	}{}

	// parse and load json config
	err := jsonconfig.Load("/Users/fizz/go/src/github.com/gohouse/study/config.json", &config)

	if err == nil {
		println("The url is", config.Url)
		println("Supported methods are", config.Methods[0], config.Methods[1])

		println("The module is", config.Module.Name, "on route", config.Module.Route)
	}
}