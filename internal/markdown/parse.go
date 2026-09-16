package markdown

import (
	"strings"

	"github.com/yuin/goldmark/v2/ast"
	"github.com/yuin/goldmark/v2/parser"
)

type Document struct {
	root   ast.Node
	source []byte
}

func Parse(src []byte) *Document {
	return &Document{
		root:   parser.New().Parse(src),
		source: src,
	}
}

func DumpAST(doc *Document) string {
	var sb strings.Builder
	doc.root.Dump(doc.source).PrettyPrint(&sb, doc.source)

	return sb.String()
}

func parseValue(value string) (name string, comment string) {
	name, comment, _ = strings.Cut(value, "#")
	return
}
