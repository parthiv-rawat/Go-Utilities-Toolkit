package pdfgen

import "github.com/jung-kurt/gofpdf"

func GeneratePDF(filename, text string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, text)
	return pdf.OutputFileAndClose(filename)
}
