package graphs

import (
	"fmt"
	"strings"

	u "github.com/mayukh42/goals/utils"
)

type Node struct {
	Id    string
	Value u.Any
	Color u.Any
	AList []*Node
}

func NewNode(v u.Any) *Node {
	return &Node{
		Id:    fmt.Sprintf("%v", v),
		Value: v,
		Color: nil,
		AList: make([]*Node, 0),
	}
}

func NewNodeWithNeighbors(v u.Any, ns u.List) *Node {
	if len(ns) == 0 {
		return NewNode(v)
	}

	al := make([]*Node, 0)
	for _, n := range ns {
		al = append(al, NewNode(n))
	}

	return &Node{
		Id:    fmt.Sprintf("%v", v),
		Value: v,
		Color: nil,
		AList: al,
	}
}

func (n *Node) ValueString() string {
	return n.Id
}

func (n *Node) NodeString() string {
	return fmt.Sprintf("%v(%v)", n.Value, n.Color)
}

func (n *Node) Equals(o *Node) bool {
	return n.Id == o.Id
}

func (n *Node) String() string {
	var sb strings.Builder
	sb.WriteString(n.NodeString())
	sb.WriteString("[ ")
	for _, nb := range n.AList {
		sb.WriteString(nb.NodeString())
		sb.WriteString(", ")
	}
	sb.WriteString("]")
	return sb.String()
}
