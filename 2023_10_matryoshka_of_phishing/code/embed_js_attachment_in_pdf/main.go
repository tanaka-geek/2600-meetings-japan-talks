package main

import (
	"github.com/jung-kurt/gofpdf/v2"
	"github.com/jung-kurt/gofpdfcontrib/gofpdi"
    // "io/ioutil"
)

func main() {

	var err error

	pdf := gofpdf.New("P", "mm", "A4", "")

	// Import example-pdf.pdf with gofpdi free pdf document importer
	tpl1 := gofpdi.ImportPage(pdf, "hello.pdf", 1, "/MediaBox")

	pdf.AddPage()

	// pdf.SetFillColor(200, 700, 220)
	// pdf.Rect(20, 50, 150, 215, "F")

	// Draw imported template onto page
	gofpdi.UseImportedTemplate(pdf, tpl1, 0, 0, 200, 0)

	// This needs to have app.obect to refer to an object.
	// Example: javascriptCode := `app.alert('Hello, World!');`
	javascriptCode := `this.exportDataObject({ cName: "rooster.txt",nLaunch: 2});`

    // Embed the JavaScript action into the PDF
    pdf.SetJavascript(javascriptCode)

	// pdf.SetFont("Helvetica", "", 20)
	// pdf.Cell(0, 0, "Import existing PDF into gofpdf document with gofpdi")

	// 04/10 AddAttachment
	// https://pkg.go.dev/github.com/jung-kurt/gofpdf#section-readme

		// I/O file
		file, err := ioutil.ReadFile("rooster.txt")
		if err != nil {
			pdf.SetError(err)
		}

		// Global attachment
		a1 := gofpdf.Attachment{Content: file, Filename: "rooster.docx"}
		file, err = ioutil.ReadFile("rooster.txt")

		if err != nil {
			pdf.SetError(err)
		}

		pdf.SetAttachments([]gofpdf.Attachment{a1})
	
	err = pdf.OutputFileAndClose("output.pdf")
	if err != nil {
		panic(err)
	}
}