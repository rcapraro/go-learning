package pdf

import (
	"fmt"
	"gencert/cert"
	"github.com/jung-kurt/gofpdf"
	"os"
	"path"
)

type PdfSaver struct {
	OutputDir string
}

func New(outputDir string) (*PdfSaver, error) {
	var p *PdfSaver
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return p, err
	}
	p = &PdfSaver{
		OutputDir: outputDir,
	}

	return p, nil
}

func (p *PdfSaver) Save(cert cert.Cert) error {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")
	pdf.SetTitle(cert.LabelTitle, false)
	pdf.AddPage()

	// background
	background(pdf)

	// header
	header(pdf, &cert)
	pdf.Ln(30)

	// body
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelPresented, "C")
	pdf.Ln(30)

	// body - student name
	pdf.SetFont("Times", "B", 40)
	pdf.WriteAligned(0, 50, cert.Name, "C")
	pdf.Ln(30)

	// body - participation
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelParticipation, "C")
	pdf.Ln(30)

	// body - date
	pdf.SetFont("Helvetica", "I", 15)
	pdf.WriteAligned(0, 50, cert.LabelDate, "C")

	footer(pdf)

	// save file
	filename := fmt.Sprintf("%v.pdf", cert.LabelTitle)
	outputPath := path.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(outputPath)
	if err != nil {
		return err
	}
	fmt.Printf("Saved certificate to '%v'\n", outputPath)
	return nil
}

func footer(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	imageWidth := 50.0
	fileName := "img/stamp.png"
	pageWidth, pageHeigth := pdf.GetPageSize()
	x := pageWidth - imageWidth - 20.0
	y := pageHeigth - imageWidth - 10.0

	pdf.ImageOptions(fileName,
		x, y, imageWidth, 0,
		false, opts, 0, "")
}

func header(pdf *gofpdf.Fpdf, cert *cert.Cert) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	margin := 30.0
	x := 0.0
	imageWidth := 30.0
	fileName := "img/gopher.png"

	pdf.ImageOptions(fileName,
		x+margin, 20, imageWidth, 0,
		false, opts, 0, "")

	pageWidth, _ := pdf.GetPageSize()
	x = pageWidth - imageWidth
	pdf.ImageOptions(fileName,
		x-margin, 20, imageWidth, 0,
		false, opts, 0, "")
	pdf.SetFont("Helvetica", "", 40)
	pdf.WriteAligned(0, 50, cert.LabelCompletion, "C")
}

func background(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	pageWidth, pageHeight := pdf.GetPageSize()
	pdf.ImageOptions("img/background.png",
		0, 0,
		pageWidth, pageHeight,
		false, opts, 0, "")

}
