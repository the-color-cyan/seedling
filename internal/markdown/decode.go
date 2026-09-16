package markdown

import (
	"errors"
	"fmt"
	"seedling/internal/tree"

	"github.com/yuin/goldmark/v2/ast"
)

var ErrListCount error = errors.New("expected a single top-level list")

type decoder struct{ src []byte }

func Decode(src []byte) (tree.Forest, error) {
	doc := Parse(src)
	return decoder{src}.decodeDocument(doc)
}

func (d decoder) decodeDocument(doc *Document) (tree.Forest, error) {
	list, err := topLevelList(doc)
	if err != nil {
		return nil, err
	}

	return d.decodeList(list, nil)
}

func (d decoder) decodeList(list *ast.List, parent *tree.Node) (tree.Forest, error) {
	if list.IsOrdered() {
		return nil, errors.New("ordered lists are unsupported")
	}

	nodes := make(tree.Forest, 0, list.ChildCount())

	for child := range list.Children() {
		item, ok := child.(*ast.ListItem)
		if !ok {
			continue
		}

		node, err := d.decodeItem(item, parent)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}

	return nodes, nil
}

func (d decoder) decodeItem(item *ast.ListItem, parent *tree.Node) (*tree.Node, error) {
	var paragraph *ast.Paragraph
	var list *ast.List
	var value string
	var err error

	for child := range item.Children() {
		switch child := child.(type) {
		case *ast.Paragraph:
			if paragraph != nil {
				return nil, errors.New("ListItem contains multiple child Paragraphs")
			}
			paragraph = child

			value, err = d.decodeParagraph(child)
			if err != nil {
				return nil, err
			}

		case *ast.List:
			if list != nil {
				return nil, errors.New("ListItem contains multiple child Lists")
			}
			list = child

		default:
			return nil, fmt.Errorf("unsupported child of ListItem: %T", child.Kind())
		}
	}

	name, comment := parseValue(value)

	// create node with nil children first
	node := tree.NewNode(name, comment, parent, nil)

	// then decode nested List
	if list != nil {
		decoded, err := d.decodeList(list, node)
		if err != nil {
			return nil, err
		}

		node.Children = decoded
	}

	return node, nil

}

// returns name and comment parsed from Paragraph's Text Value
func (d decoder) decodeParagraph(paragraph *ast.Paragraph) (string, error) {
	var value string

	for child := range paragraph.Children() {
		if value != "" {
			return "", errors.New("Paragraph contains multiple child Texts")
		}

		if text, ok := child.(*ast.Text); ok {
			value = text.Value.Str(d.src)
		}
	}

	return value, nil
}

func topLevelList(doc *Document) (*ast.List, error) {
	var lists []*ast.List

	for child := range doc.root.Children() {
		if list, ok := child.(*ast.List); ok {
			lists = append(lists, list)
		}
	}

	if len(lists) != 1 {
		// TODO: possibly pass file path through
		return nil, fmt.Errorf("%w - found %d", ErrListCount, len(lists))
	}
	return lists[0], nil
}
