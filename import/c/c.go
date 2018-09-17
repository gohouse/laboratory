package c

import (
	"fmt"
	"github.com/gohouse/laboratory/import/b"
)

func C(arg b.B) string {
	fmt.Println(arg.B2)
	return "ccc"
}
