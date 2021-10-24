package tree

import (
	"fmt"

	u "github.com/mayukh42/goals/utils"
)

type Node struct {
	Value u.Any
	Left  *Node
	Right *Node
}

func NewNode(v u.Any) *Node {
	return &Node{
		Value: v,
	}
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Value)
}

func (n *Node) Height(h uint) uint {
	hL, hR := h, h
	if n.Left != nil {
		hL = n.Left.Height(h + 1)
	}

	if n.Right != nil {
		hR = n.Right.Height(h + 1)
	}

	return u.Max(hL, hR)
}

func (n *Node) IsBalanced() bool {
	// at least we have n
	hL, hR := uint(1), uint(1)
	if n.Left != nil {
		hL = n.Left.Height(hL + 1)
	}

	if n.Right != nil {
		hR = n.Right.Height(hR + 1)
	}

	return u.AbsDiff(hL, hR) <= 1
}

/** create balanced tree from array
 * 	[1 2 3 4 5 6 7]
 * 				4
 * 			2		6
 * 		1	  3   5		7
 *  TODO: sort list
 */
func CreateFromArray(ls u.List) *Node {
	// [:mid], [mid], [mid+1:]
	n := len(ls)
	if n == 0 {
		return nil
	}
	mid := n / 2
	node := NewNode(ls[mid])
	node.Left = CreateFromArray(ls[:mid])
	// golang is easier on the array index out of bounds exception
	node.Right = CreateFromArray(ls[mid+1:])
	return node
}
