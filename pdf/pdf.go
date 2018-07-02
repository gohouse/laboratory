package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 11)
	//pdf.Image("test.png", 10, 10, 30, 0, false, "", 0, "")
	//pdf.Text(50, 20, "test.png")
	//pdf.Image("test.gif", 10, 40, 30, 0, false, "", 0, "")
	//pdf.Text(50, 50, "test.gif")
	pdf.Image("1.jpg", 10, 130, 30, 0, false, "", 0, "")
	pdf.Text(50, 140, "1.jpg")

	err := pdf.OutputFileAndClose("write_pdf_with_image.pdf")
	if err != nil {
		fmt.Println(err)
	}
}