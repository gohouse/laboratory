package a

type AA struct {
	AAA string
}

func A() string {
	return "aaa"
}

func CallC() interface{} {
	//var tmpB b.B
	var aaT = &AA{}
	aaT.AAA = "bbbbbb"
	//res := c.C(aaT.bT)
	//fmt.Println(res)

	return aaT
}