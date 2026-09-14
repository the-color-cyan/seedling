package markdown

import (
	"fmt"
	"seedling/internal/tree"

	astquery "github.com/yuin/goldmark-astquery"
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
	listNode := astquery.Match(doc.root, doc.source,
		astquery.ForEachDescendant(
			astquery.Bind("lists",
				astquery.NodeKind(ast.KindList),
			),
		),
	)["lists"]
	if len(listNode) != 1 {
		// TODO: possibly pass file path through
		return nil, fmt.Errorf("expected a single list, found %d", len(listNode))
	}

	list, ok := listNode[0].(*ast.List)
	if !ok {
		return nil, fmt.Errorf("expected *ast.List, got %T", listNode)
	}

	return list, nil
}
