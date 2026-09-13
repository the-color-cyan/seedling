package main

import (
	"fmt"
	"os"
	"seedling/internal/markdown"
)

const header string = `
	|-----------|
	|--ASTDUMP--|
	|-----------|

	`

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: astdump <markdown-file>")
		os.Exit(2)
	}

	doc, err := markdown.ParseFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	println(header)
	println(markdown.DumpAST(doc))
}
