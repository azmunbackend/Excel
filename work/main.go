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

    err = pdf.SetFont("Times New Roman", "", 40)
    if err != nil {
        log.Print(err.Error())
        return
    }
	pdf.AddPage()
	pdf.Image("download.png", 20, 10, &gopdf.Rect{W: 150, H: 120})
	pdf.SetXY(20, 180)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFontSize(12.5)
	pdf.Cell(nil, "Уссат мерген")
	pdf.SetXY(20, 205)
	pdf.Cell(nil, "Индивидуальное предприятие")
	pdf.SetXY(20, 230)
	pdf.Cell(nil, "Туркменистан")
	pdf.SetXY(20, 255)
	pdf.Cell(nil, "DK: 10332")
	pdf.SetXY(20, 300)
	pdf.Cell(nil, "Рысгал АКБ")
	pdf.SetXY(20, 325)
	pdf.Cell(nil, "23202934173861109012000")
	pdf.SetXY(20, 350)
	pdf.Cell(nil, "г.Ашхабад., улица 1946(Анкара)., д. 23")
	pdf.SetXY(20, 375)
	pdf.Cell(nil, "BAB: 390101738")
	pdf.SetXY(20, 425)
	pdf.Cell(nil, "Esas:")
	pdf.SetXY(50, 425)
	pdf.Cell(nil, "Ýerli kärhanalar üçin demir ýol kodyny almak hyzmat tölegi")
	pdf.SetXY(486, 50)
	pdf.SetTextColor(0, 0, 255)
	pdf.SetFontSize(15)
	pdf.Cell(nil, "№1290/24")

	pdf.SetTextColor(0, 0, 0)
	pdf.SetFontSize(12.5)
	pdf.SetXY(486, 180)
	pdf.Cell(nil, "08.05.2024")
	pdf.SetXY(456, 205)
	pdf.Cell(nil, "Demirýollary AGPJ")
	pdf.SetXY(440, 230)
	pdf.Cell(nil, "Aşgabat, Türkmenistan")
	pdf.SetXY(470, 255)
	pdf.Cell(nil, "+993 12 383922")
	pdf.SetXY(430, 300)
	pdf.Cell(nil, `"Türkmenbaşy" PTB`)
	pdf.SetXY(400, 325)
	pdf.Cell(nil, "23202934130477301001000")
	pdf.SetXY(420, 350)
	pdf.Cell(nil, "Aşgabat, Türkmenistan")
	pdf.SetXY(420, 375)
	pdf.Cell(nil, "BAB:5 390101304")
	pdf.SetXY(400, 400)
	pdf.Cell(nil, "Salgyt belgisi: 102511004149")

    pdf.WritePdf("test.pdf")
}
