package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jung-kurt/gofpdf"
	"github.com/waigani/diffparser"
	"os"
	"strconv"
)

var (
	debugParsing = false
)

func main() {
	// read params
	var diffFilePath string
	var outputFilePath string
	var format string

	flag.StringVar(&diffFilePath, "d", "", "diff file path")
	flag.StringVar(&outputFilePath, "o", "", "output file path")
	flag.StringVar(&format, "f", "pdf", "format [pdf]")

	flag.Parse()

	if diffFilePath == "" || outputFilePath == "" || format == "" {
		log.Fatal("You need pass all params to work, check params using help command: godiffexporter -h")
	}

	fileContentRaw, _ := ioutil.ReadFile(diffFilePath)
	diff, err := diffparser.Parse(string(fileContentRaw))

	if err != nil {
		log.Fatalf("Error while parse diff file: %v", err)
	}

	fmt.Println("Parsing...")

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFontLocation(os.Getenv("GOPATH") + "src/github.com/prsolucoes/godiffexporter/fonts")
	pdf.AddFont("Helvetica-1251", "", "helvetica_1251.json")
	tr := pdf.UnicodeTranslatorFromDescriptor("")

	// title
	pdf.SetFont("Helvetica", "B", 16)
	pdf.CellFormat(0, 16, "Diff Exporter", "", 2, "C", false, 0, "")

	// subtitle
	pdf.SetFont("Helvetica", "B", 10)
	pdf.CellFormat(0, 10, "https://github.com/prsolucoes/godiffexporter", "", 2, "C", false, 0, "https://github.com/prsolucoes/godiffexporter")
	pdf.Ln(4)

	// iterate over all files
	for fileCount, file := range diff.Files {

		if fileCount > 0 {
			pdf.Ln(8)
		}

		if debugParsing {
			fmt.Println("-------------------------------------------------------------------------------")
		}

		if file.OrigName == file.NewName || file.NewName == "" {
			if debugParsing {
				fmt.Printf("| File: %v\n", file.OrigName)
			}

			pdf.SetFont("Helvetica", "B", 11)
			pdf.SetFillColor(222, 222, 222)
			pdf.MultiCell(0, 11, "File: " + tr(file.OrigName), "1", "L", true)
		} else {
			if debugParsing {
				fmt.Printf("| Renamed File: %v\n", file.OrigName)
				fmt.Printf("| To File: %v\n", file.NewName)
			}

			pdf.SetFont("Helvetica", "B", 11)
			pdf.SetFillColor(222, 222, 222)
			pdf.MultiCell(0, 11, "Renamed file: " + tr(file.OrigName), "1", "L", true)
			pdf.MultiCell(0, 11, "To file: " + tr(file.NewName), "1", "L", true)
		}

		if debugParsing {
			fmt.Println("-------------------------------------------------------------------------------")
		}

		// iterate over all file hunks
		for _, hunk := range file.Hunks {
			newRange := hunk.NewRange
			oldRange := hunk.OrigRange

			currentLine := -1

			for _, lineOld := range oldRange.Lines {
				currentLine = lineOld.Number
				exportLineToPDF(pdf, lineOld)

				for _, lineNew := range newRange.Lines {
					if lineNew.Number <= lineOld.Number && lineNew.Number >= currentLine {
						exportLineToPDF(pdf, lineNew)
					}
				}
			}
		}
	}

	err = pdf.OutputFileAndClose(outputFilePath)

	if err != nil {
		log.Fatalf("Error while generate PDF file: %v", err)
	} else {
		fmt.Println("Finished")
	}
}

func exportLineToPDF(pdf *gofpdf.Fpdf, line *diffparser.DiffLine) {
	tr := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetFont("Helvetica", "", 10)
	prefix := ""

	if line.Mode == diffparser.ADDED {
		if debugParsing {
			fmt.Println("+ " + line.Content)
		}

		pdf.SetFillColor(221, 255, 221)
		prefix = "+ "
	} else if line.Mode == diffparser.REMOVED {
		if debugParsing {
			fmt.Println("- " + line.Content)
		}

		pdf.SetFillColor(254, 232, 233)
		prefix = "- "
	} else if line.Mode == diffparser.UNCHANGED {
		if debugParsing {
			fmt.Println(" " + line.Content)
		}

		pdf.SetFillColor(255, 255, 255)
		prefix = "  "
	}

	pdf.MultiCell(0, 10, tr("(" + strconv.Itoa(line.Number) + ") " + prefix + line.Content), "1", "L", true)
}