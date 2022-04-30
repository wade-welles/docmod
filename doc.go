package docmod

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	extractor "github.com/unidoc/unipdf/extractor"
	model "github.com/unidoc/unipdf/model"
)

func ParsePDFsInFolder(path string) {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	// Iterating through all the files
	for _, f := range files {
		// checking for file extensions
		ext := filepath.Ext(f.Name())
		// if extension relevant

		if ext == ".pdf" {
			// then open file
			f, err := os.Open(f.Name())
			if err != nil {
				panic(err)
			}
			defer f.Close()

			reader, err := model.NewPdfReaderLazy(f)
			if err != nil {
				panic(err)
			}
			// Read page
			p, err := reader.GetPage(1)
			if err != nil {
				panic(err)
			}
			// create extractor for page
			ex, err := extractor.New(p)
			if err != nil {
				panic(err)
			}
			// read text using extractor
			text, err := ex.ExtractText()
			if err != nil {
				panic(err)
			}
			fmt.Printf("and the text of the document is!: %v\n", text)
		}
	}
}
