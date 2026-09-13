package markdown

import (
	"os"
	"strings"

	"github.com/yuin/goldmark/v2/ast"
	"github.com/yuin/goldmark/v2/parser"
)

type Document struct {
	root   ast.Node
	source []byte
}

func ParseFile(path string) (*Document, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	parser := parser.New()
	return &Document{
		root:   parser.Parse(file),
		source: file,
	}, nil
}

func DumpAST(doc *Document) string {
	var sb strings.Builder

	doc.root.Dump(doc.source).PrettyPrint(&sb, doc.source)

	return sb.String()
}
