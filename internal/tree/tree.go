package tree

import (
	"strings"
)

type Node struct {
	Kind     NodeType
	Name     string
	Comment  string
	Parent   *Node
	Children Forest
}

type NodeType uint8

const (
	KindFile NodeType = iota
	KindDirectory
)

type Forest []*Node

func NewNode(name string, comment string, parent *Node, children Forest) *Node {
	var kind NodeType

	if len(children) > 0 || strings.HasSuffix(name, "/") {
		kind = NodeType(KindDirectory)
	}

	return &Node{
		Kind:     kind,
		Name:     name,
		Comment:  comment,
		Parent:   parent,
		Children: children,
	}
}
