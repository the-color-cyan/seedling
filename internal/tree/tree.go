package tree

type Node struct {
	parent   *Node
	children *[]Node
	metadata NodeMetadata
}

type NodeMetadata struct {
	kind    NodeType
	name    string
	comment *string
}

type NodeType uint8

const (
	KindUnknown NodeType = iota
	KindFile
	KindDirectory
)
