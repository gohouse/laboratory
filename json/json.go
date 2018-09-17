package main
import (
	"encoding/json"
	"fmt"
	"github.com/gohouse/gorose/examples"
)

// Product 商品信息
type Product struct {
	Name      string
	test *examples.TestT
}

func main() {
	p := &Product{}
	p.Name = "Xiao mi 6"
	data, _ := json.Marshal(p)
	fmt.Println(string(data))

	var tmp *examples.TestT
	json.Unmarshal(data,&tmp)
	fmt.Println(tmp.Name)
	p.test = tmp

	fmt.Println(p)
}