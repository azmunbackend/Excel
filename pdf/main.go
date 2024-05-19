package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main() {
    pdf := gopdf.GoPdf{}
    pdf.Start(gopdf.Config{ PageSize: *gopdf.PageSizeA4 })

    err := pdf.AddTTFFont("Times New Roman", "Times New Roman.ttf")
    if err != nil {
        log.Print(err.Error())
        return
    }

    err = pdf.SetFont("Times New Roman", "", 30)
    if err != nil {
        log.Print(err.Error())
        return
    }
	pdf.SetTextColor(0, 0, 0)

    pdf.AddHeader(func() {
		pdf.SetXY(100,5)
		pdf.SetFontSize(40)
        pdf.Cell(nil, "Linux barada taze dusunjeler!")
    })
	pdf.AddFooter(func() {
		pdf.SetX(250)
		pdf.SetY(825)
		pdf.SetFontSize(10)
		pdf.Cell(nil, "2024 golangPdf")
	})
	
    pdf.AddPage()
	pdf.Line(20, 40, 570, 40)
    pdf.SetY(50)
    pdf.Text("Ilkinji owrenmeli zatlar")
	pdf.Rotate(0, 0, 0)
	pdf.Image("../img.jpeg", 10, 110, nil)

    pdf.WritePdf("test.pdf")
}

