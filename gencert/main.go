package main

import (
	"flag"
	"fmt"
	"gencert/cert"
	"gencert/html"
	"gencert/pdf"
	"os"
)

func main() {

	file := flag.String("file", "", "CSV Students file.")
	outputType := flag.String("type", "pdf", "Output type of the certificate.")
	flag.Parse()

	if len(*file) <= 0 {
		fmt.Printf("Invalid file. got=%v\n", *file)
		os.Exit(1)
	}

	var saver cert.Saver
	var err error
	switch *outputType {
	case "html":
		saver, err = html.New("output")
	case "pdf":
		saver, err = pdf.New("output")
	default:
		err = fmt.Errorf("unknown output type. got='%v'", *outputType)
	}
	if err != nil {
		fmt.Printf("Could not create generator: %v", err)
		os.Exit(1)
	}

	certs, err := cert.ParseCsv(*file)
	if err != nil {
		fmt.Printf("Could not parse CSV file: %v", err)
		os.Exit(1)
	}

	for _, c := range certs {
		err = saver.Save(*c)
		if err != nil {
			fmt.Printf("Could not save Certificate. got=%v\n", err)
		}
	}
}
