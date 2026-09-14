package markdown

import (
	"fmt"
	"seedling/internal/tree"

	"github.com/yuin/goldmark/v2/ast"
)

func Decode(src []byte) (tree.Tree, error) {
	doc := Parse(src)
	return decodeDocument(doc)
}

func decodeDocument(doc *Document) (tree.Tree, error) {
	list, err := getList(doc)
	if err != nil {
		return nil, err
	}

	panic("not yet implemented")
}

func getList(doc *Document) (*ast.List, error) {
	var lists []*ast.List

	for child := range doc.root.Children() {
		if list, ok := child.(*ast.List); ok {
			lists = append(lists, list)
		}
	}

	if len(lists) != 1 {
		// TODO: possibly pass file path through
		return nil, fmt.Errorf("expected a single top-level list, found %d", len(lists))
	}

	return lists[0], nil
}
