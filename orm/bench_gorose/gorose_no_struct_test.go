package bench_gorose

import (
	"fmt"
	_ "github.com/gohouse/gorose/driver/mysql"
	"testing"
)

func TestNoStruct_First(test *testing.T)  {
	db := Connection.NewSession()
	res,err2 := db.Table("users").Fields("name, age, job").First()
	if err2 != nil {
		test.Error("FAIL: open failed.", err2)
		return
	}
	test.Log(fmt.Sprintf("PASS: open: %v", res))
}


func BenchmarkNoStruct_First(bmtest *testing.B) {
	db := Connection.NewSession()
	for cnt := 0; cnt < bmtest.N; cnt++ {
		db.Table("users").Fields("name, age, job").First() // 200	   6567155 ns/op
	}
}