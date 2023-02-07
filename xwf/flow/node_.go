package flow

import (
	"go-phoenix/xwf/enum"
)

func (node *Node) DiagramId() string {
	return node.diagramId
}

func (node *Node) Key() int {
	return node.key
}

func (node *Node) Code() string {
	return node.code
}

func (node *Node) Name() string {
	return node.name
}

func (node *Node) Category() enum.Category {
	return node.category
}
