package main

import (
	"fmt"

	docmod "github.com/wade-welles/docmod"
)

const cwd = "."

func main() {
	fmt.Println("docmod")
	fmt.Println("======")
	fmt.Println("Checking for PDFs in the current working directory ( " + cwd + " )")

	docmod.ParsePDFsInFolder(cwd)
}
