package main

import (
	"bytes"
	"fmt"
	"log"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	fmt.Println("vim-go")
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Set global options
	pdfg.Dpi.Set(10)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)
	pdfg.Grayscale.Set(true)

	// pdfg.Dpi.Set(600)
	pdfg.NoCollate.Set(false)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.MarginBottom.Set(40)
	// Create a new input page from an URL
	// page := wkhtmltopdf.NewPage("https://godoc.org/github.com/SebastiaanKlippert/go-wkhtmltopdf")
	// page := wkhtmltopdf.NewPage("https://godoc.org/github.com/SebastiaanKlippert/go-wkhtmltopdf")

	html := "<svg xmlns=\"http://www.w3.org/2000/svg\" version=\"1.1\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xmlns:svgjs=\"http://svgjs.com/svgjs\" width=\"717.5\" height=\"421.3623161315918\" viewBox=\"-115 -58.20539474487305 717.5 421.3623161315918\"><path d=\"M36.110421457544106 -58.205393653062565C71.9438582389125 -49.47458460510318 202.4165632186316 -10.230809047959387 275 0C347.5834367813684 10.230809047959387 483.25 8.5 520 10 \" stroke=\"blue\" fill=\"none\"></path><path d=\"M520 10C521.5 36.5 517.625 146.16666666666666 530 186.66666666666666C542.375 227.16666666666666 591.625 266 602.5 280 \" stroke=\"blue\" fill=\"none\"></path><path d=\"M602.5 280L530 360 \" stroke=\"blue\" fill=\"none\"></path><path d=\"M-19.896005322753986 363.1569189824969C24.338395475659112 359.6833811351224 192.5155992015869 340.47353784737453 275 340C357.4844007984131 339.52646215262547 491.75 357 530 360 \" stroke=\"blue\" fill=\"none\"></path><path d=\"M-115 290L-19.896005322753986 363.1569189824969 \" stroke=\"blue\" fill=\"none\"></path><path d=\"M-115 290C-101.69447892712279 271.5897303173469 -43.982568326573045 202.88540814761873 -26.29652618081863 167.26486878231248C-8.610484035064218 131.64432941700622 -6.454094507391673 86.35027692993121 2.9069476383627375 52.52973756462495C12.267989784117148 18.709198199318692 31.1299003846669 -41.59512397040943 36.110421457544106 -58.205393653062565 \" stroke=\"blue\" fill=\"none\"></path><path d=\"M0 0L0 280 \" stroke=\"red\"></path><path d=\"M0 0L530 0 \" stroke=\"red\"></path><path d=\"M530 0L530 280 \" stroke=\"red\"></path><path d=\"M0 280L530 280 \" stroke=\"red\"></path><path d=\"M0 280L-115 280 \" stroke=\"red\"></path><path d=\"M275 0L275 340 \" stroke=\"red\"></path><path d=\"M-57.5 280L20 0 \" stroke=\"red\"></path><path d=\"M-57.5 280L36.110421457544106 -58.205393653062565 \" stroke=\"red\"></path><path d=\"M20 0L36.110421457544106 -58.205393653062565 \" stroke=\"red\"></path><path d=\"M0 340L530 340 \" stroke=\"red\"></path><path d=\"M530 280L602.5 280 \" stroke=\"red\"></path></svg>"
	// Client code
	// pdfg := wkhtmltopdf.NewPDFPreparer()
	// htmlfile, err := ioutil.ReadFile("./testfiles/htmlsimple.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	page := wkhtmltopdf.NewPageReader(bytes.NewReader([]byte(html)))
	// pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader([]byte(html))))
	// page.FooterRight.Set("[page]")

	// wkhtmltopdf.NewPDFGenerator()
	// page := wkhtmltopdf.NewPage(pdfg.)

	// Set options for this page
	// page.FooterRight.Set("[page]")
	// page.FooterFontSize.Set(10)
	page.Zoom.Set(2)

	// Add to document
	pdfg.AddPage(page)
	pdfg.Dpi.Set(10)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./simplesample.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}
