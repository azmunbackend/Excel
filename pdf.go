package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func main() {
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()

    pdf.SetFont("Arial", "", 16)
    pdf.Cell(40, 10, "Hello, World!")

	pdf.Ln(30)

	opt := gofpdf.ImageOptions{
        ImageType: "JPG",
    }

    pdf.ImageOptions("img.jpeg", 10, 10, 0, 0, false, opt, 0, "")
	
	pdf.AddPage()
	pdf.Cell(10, 30, "Amanyaz Muhammetamanow")

    err := pdf.OutputFileAndClose("test2.pdf")
    if err != nil {
       fmt.Println(err)
    }
}