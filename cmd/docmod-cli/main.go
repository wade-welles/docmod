package main

import (
	"fmt"

	docgen "github.com/lawstarlabs/docgen"
)

const cwd = "."

func main() {
	fmt.Println("docgen")

	docgen.ParsePDFsInFolder(cwd)
}
