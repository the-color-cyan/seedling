package main

import (
	"fmt"
	"os"
	"seedling/internal/markdown"
)

const testPath = "testdata/test_tree.md"

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: astdump <markdown-file>")
		os.Exit(2)
	}

	doc, err := markdown.ParseFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	println("--ASTDUMP--")
	println("-----------")
	println()
	println(markdown.DumpAST(doc))
}
