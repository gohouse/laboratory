package main

import (
	"github.com/nguyenthenguyen/docx"
	"fmt"
)
const testFile = "./wordtmp2.docx"
const testFileResult = "./TestDocumentResult.docx"

func loadFile(file string) (*docx.Docx,*docx.ReplaceDocx) {
	r, err := docx.ReadDocxFile(file)
	if err != nil {
		panic(err)
	}

	return r.Editable(), r
}

func TestReplace() {
	d,r := loadFile(testFile)
	d.Replace("document.", "line1\r\nline2", 1)
	d.WriteToFile(testFileResult)

	r.Close()
	//d = loadFile(testFileResult)

}

func main() {
	//r, err := docx.ReadDocxFile("./wordtmp.docx")
	//if err != nil {
	//	panic(err)
	//}
	//docx1 := r.Editable()
	//// Replace like https://golang.org/pkg/strings/#Replace
	//docx1.Replace("a101", "new_1_1", -1)
	//docx1.Replace("a102", "new_1_2", -1)
	//docx1.Replace("a103", "new_1_3", -1)
	//docx1.Replace("a104", "new_1_4", -1)
	////docx1.ReplaceLink("http://example.com/", "https://github.com/nguyenthenguyen/docx")
	////docx1.ReplaceHeader("out with the old", "in with the new")
	////docx1.ReplaceFooter("Change This Footer", "new footer")
	//docx1.WriteToFile("./new_result_1.docx")
	//
	////docx2 := r.Editable()
	////docx2.Replace("old_2_1", "new_2_1", -1)
	////docx2.Replace("old_2_2", "new_2_2", -1)
	////docx2.WriteToFile("./new_result_2.docx")

	//r.Close()
	TestReplace()
	fmt.Println("finish...")
}