package list

import (
	"fmt"
	"strings"
	"time"

	u "github.com/mayukh42/goals/utils"
)

type List u.List

type Node struct {
	Value interface{}
	Next  *Node
}

func (n *Node) ValueStr() string {
	return fmt.Sprintf("%v, ", n.Value)
}

func NewNode(v interface{}) *Node {
	return &Node{
		Value: v,
		Next:  nil,
	}
}

func DeleteNode(n *Node) {
	curr := n
	for curr != nil {
		fmt.Printf("curr: %v\n", curr.ValueStr())
		curr = Unlink(curr)
	}
}

/** Link()
 */
func Link(first, second *Node) *Node {
	if second == nil {
		return first
	}

	if first == nil {
		return second
	}

	first.Next = second
	return second
}

// WIP
func LinkAndShift(head, elem *Node) (*Node, *Node) {
	var tail *Node
	if head == nil {
		head = elem
		tail = head
	} else {
		// WIP
		tail = head.Next
		tail = Link(tail, elem)
	}

	return head, tail
}

func Unlink(head *Node) *Node {
	if head != nil {
		tail := head.Next
		head.Next = nil
		return tail
	}

	return head
}

// category theory [?]
func CreateFromArray(xs List) *Node {
	var last, head *Node
	for _, x := range xs {
		n := NewNode(x)
		if head == nil {
			head = n
			last = head
		} else {
			last = Link(last, n)
		}
	}
	return head
}

func (n *Node) String() string {
	s := &strings.Builder{}
	s.WriteString("[ ")

	curr := n
	for curr != nil {
		s.WriteString(curr.ValueStr())
		curr = curr.Next
	}

	s.WriteString("]")
	return s.String()
}

type DNode struct {
	Value u.Any
	Time  time.Time
	Prev  *DNode
	Next  *DNode
}

func NewDNode(v interface{}) *DNode {
	return &DNode{
		Value: v,
		Time:  time.Now().UTC(),
		Next:  nil,
		Prev:  nil,
	}
}

type DList struct {
	Head *DNode
	Tail *DNode
	Size uint
}

func NewDList() *DList {
	return &DList{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
}

// Suffix Link ->
func (dn *DNode) Link(n *DNode) *DNode {
	if n == nil {
		return dn
	}

	dn.Next = n
	n.Prev = dn
	// n.Next = nil

	return n
}

func (dn *DNode) Sanitize() *DNode {
	dn.Next = nil
	dn.Prev = nil

	return dn
}

func (dn *DNode) ValueStr() string {
	return fmt.Sprintf("%v", dn.Value)
}

func (dn *DNode) Equals(on *DNode) bool {
	if on != nil {
		return dn.ValueStr() == on.ValueStr()
	}
	return false
}

func (dl *DList) Suffix(n *DNode) *DList {
	if n == nil {
		return dl
	}

	n.Sanitize()
	if dl.Head == nil {
		// invariant: dl.Head and dl.Tail can either be both nil or none nil
		dl.Head = n
		dl.Tail = n
	} else {
		nt := dl.Tail.Link(n)
		dl.Tail = nt
	}

	dl.Size++
	return dl
}

func (dl *DList) String() string {
	s := &strings.Builder{}
	s.WriteString("[ ")
	s.WriteString(fmt.Sprintf("(%d) ", dl.Size))

	curr := dl.Head
	for curr != nil {
		s.WriteString(curr.ValueStr())
		if curr == dl.Head {
			s.WriteRune('h')
		}
		if curr == dl.Tail {
			s.WriteRune('t')
		}
		s.WriteString(" <> ")
		curr = curr.Next
	}

	s.WriteString("]")
	return s.String()
}

func (dl *DList) Prefix(n *DNode) *DList {
	if n == nil {
		return dl
	}

	n.Sanitize()
	if dl.Head == nil {
		// invariant: dl.Head and dl.Tail can either be both nil or none nil
		dl.Head = n
		dl.Tail = n
	} else {
		n.Link(dl.Head)
		dl.Head = n
	}

	dl.Size++
	return dl
}

func (dl *DList) RemoveHead() *DNode {
	if dl.Head != nil {
		e := dl.Head
		dl.Head = dl.Head.Next
		dl.Head.Prev = nil
		dl.Size--
		return e
	}
	return nil
}

func (dl *DList) RemoveTail() *DNode {
	if dl.Tail != nil {
		e := dl.Tail
		dl.Tail = dl.Tail.Prev
		dl.Tail.Next = nil
		dl.Size--
		return e
	}
	return nil
}

func (dl *DList) RemoveNode(n *DNode) *DNode {
	if dl.Head.Equals(n) {
		return dl.RemoveHead()
	}

	if dl.Tail.Equals(n) {
		return dl.RemoveTail()
	}

	// somewhere in the middle
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	dl.Size--

	n.Sanitize()
	return n
}

func (dl *DList) GetSize() int {
	return int(dl.Size)
}

// Queue, FIFO
func CreateDListSuffixed(xs u.List) *DList {
	dl := NewDList()
	for _, x := range xs {
		dl = dl.Suffix(NewDNode(x))
	}
	return dl
}

// Stack, LIFO
func CreateDListPrefixed(xs u.List) *DList {
	dl := NewDList()
	for _, x := range xs {
		dl = dl.Prefix(NewDNode(x))
	}
	return dl
}
