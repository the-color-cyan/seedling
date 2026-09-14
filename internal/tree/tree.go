package tree

type Node struct {
	Kind     NodeType
	Name     string
	Comment  *string
	Parent   *Node
	Children *[]Node
}

type NodeType uint8

const (
	KindUnknown NodeType = iota
	KindFile
	KindDirectory
)

type Tree []Node
