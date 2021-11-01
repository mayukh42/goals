package graphs

import (
	"fmt"
	"strings"

	u "github.com/mayukh42/goals/utils"
)

/** union find (partition) data structure
 * linked list
 * tree*
 *
 * requires a ref map for nodes/ indices
 * 	maintained on the caller side
 */

type UFNode struct {
	Value   u.Any
	LinkPtr *UFNode
	RootPtr *UFNode
	Size    uint
}

func NewUFNode(v u.Any) *UFNode {
	n := &UFNode{
		Value: v,
		Size:  1,
	}
	n.LinkPtr = n
	// TODO: implement path compression
	n.RootPtr = nil
	return n
}

func (n *UFNode) IsSingleton() bool {
	return n == n.LinkPtr
}

func (n *UFNode) ValueStr() string {
	return fmt.Sprintf("%v{%d} ", n.Value, n.Size)
}

func (n *UFNode) String() string {
	var sb strings.Builder
	sb.WriteString("[ ")
	sb.WriteString(n.ValueStr())
	if n.LinkPtr != nil {
		sb.WriteString("pr:")
		sb.WriteString(n.LinkPtr.ValueStr())
	}
	if n.RootPtr != nil {
		sb.WriteString("rt:")
		sb.WriteString(n.RootPtr.ValueStr())
	}
	sb.WriteString(" ]")
	return sb.String()
}

type UFDS struct {
	Id    u.Any
	Root  *UFNode
	Nodes []*UFNode
	Size  int
}

/** Create empty UFDS wrapper struct
 */
func NewUFDS(id u.Any) *UFDS {
	return &UFDS{
		Id:    id,
		Root:  nil,
		Nodes: make([]*UFNode, 0),
		Size:  0,
	}
}

/** MakeSet(e): Create a singleton set containing
 * 	the element e and return the position storing e
 * 	in this set
 */
func (uf *UFDS) MakeSet(v u.Any) (*UFNode, int) {
	n := NewUFNode(v)
	pos := uf.Size
	uf.Nodes = append(uf.Nodes, n)
	if uf.Root == nil {
		uf.Root = n
	}
	uf.Size += 1
	return n, pos
}

func (uf *UFDS) MakeSetWithNode(n *UFNode) (*UFNode, int) {
	pos := uf.Size
	uf.Nodes = append(uf.Nodes, n)
	uf.Size += 1
	return n, pos
}

/** Union(A,B): Return the set A U B, naming the
 * 	result “A” or “B”
 */
func (uf *UFDS) Union(u1, u2 *UFNode) *UFNode {
	var lrg, sml *UFNode
	if u1.Size >= u2.Size {
		lrg = u1
		sml = u2
	} else {
		lrg = u2
		sml = u1
	}

	// add sml to lrg
	sml.LinkPtr = lrg
	lrg.Size += sml.Size
	return lrg
}

/** Find(e): Return the set containing the element e
 */
func (uf *UFDS) Find(e *UFNode) *UFNode {
	root := e
	for !root.IsSingleton() {
		root = root.LinkPtr
	}
	return root
}

func (uf *UFDS) Consistent() bool {
	return uf.Size == int(uf.Root.Size)
}

func (uf *UFDS) NodeAt(i int) *UFNode {
	if i < uf.Size {
		return uf.Nodes[i]
	}
	return nil
}

func (uf *UFDS) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("ufds.%v{%d}[root:%v] \n\t",
		uf.Id,
		uf.Size,
		uf.Root,
	))
	// invariant: u.Size == len(u.Nodes)
	if uf.Size > 1 {
		root := uf.Nodes[0]
		sb.WriteString(fmt.Sprintf("root: %v \n\t", root.Value))
	}
	for _, n := range uf.Nodes {
		sb.WriteString(fmt.Sprintf("%v \n\t", n))
	}
	sb.WriteString("\n")
	return sb.String()
}
