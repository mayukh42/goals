package algo

import (
	c "github.com/mayukh42/goals/collections"
	g "github.com/mayukh42/goals/graphs"
	t "github.com/mayukh42/goals/tree"
	u "github.com/mayukh42/goals/utils"
)

func BFSTree(n *t.Node) u.List {
	q := c.NewQueue()
	q.Enqueue(n)
	acc := make(u.List, 0)
	for !q.IsEmpty() {
		n, _ := q.Dequeue()
		if n == nil {
			continue
		}
		if n, ok := n.(*t.Node); ok {
			acc = append(acc, n.Value)
			if n.Left != nil {
				q.Enqueue(n.Left)
			}
			if n.Right != nil {
				q.Enqueue(n.Right)
			}
		}
	}
	return acc
}

/** non-recursive version,
 * capturing the recursion in a list
 */
func DFSTree(n *t.Node) u.List {
	s := c.NewStack()
	s.Push(n)
	acc := make(u.List, 0)
	for !s.IsEmpty() {
		n, _ := s.Pop()
		if n == nil {
			continue
		}
		if n, ok := n.(*t.Node); ok {
			acc = append(acc, n.Value)
			if n.Right != nil {
				s.Push(n.Right)
			}
			if n.Left != nil {
				s.Push(n.Left)
			}
		}
	}
	return acc
}

func BFSGraph(n *g.Node) u.List {
	if n == nil {
		return nil
	}
	n.Color = "v"
	q := c.NewQueue()
	q.Enqueue(n)
	acc := make(u.List, 0)
	for !q.IsEmpty() {
		n, _ := q.Dequeue()
		if n == nil {
			continue
		}
		if n, ok := n.(*g.Node); ok {
			acc = append(acc, n.Value)
			for _, nb := range n.AList {
				// check for color
				if nb.Color != nil {
					nb.Color = "v"
					q.Enqueue(nb)
				}
			}
		}
	}
	return acc
}
