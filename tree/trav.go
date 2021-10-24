package tree

import (
	u "github.com/mayukh42/goals/utils"
)

func (n *Node) Infix(acc u.List) u.List {
	var ll, lr u.List

	if n.Left != nil {
		ll = n.Left.Infix(acc)
	}

	if n.Right != nil {
		lr = n.Right.Infix(acc)
	}

	acc = append(acc, ll...)
	acc = append(acc, n.String())
	acc = append(acc, lr...)

	return acc
}

func (n *Node) Size() int {
	return len(n.Infix(u.List{}))
}

func (n *Node) Prefix(acc u.List) u.List {
	var ll, lr u.List

	if n.Left != nil {
		ll = n.Left.Prefix(acc)
	}

	if n.Right != nil {
		lr = n.Right.Prefix(acc)
	}

	acc = append(acc, n.String())
	acc = append(acc, ll...)
	acc = append(acc, lr...)

	return acc
}

func (n *Node) PostfixHeight(acc u.List, lvl uint) (u.List, uint) {
	var ll, lr u.List
	lvlL, lvlR := lvl, lvl

	if n.Left != nil {
		ll, lvlL = n.Left.PostfixHeight(acc, lvl+1)
	}

	if n.Right != nil {
		lr, lvlR = n.Right.PostfixHeight(acc, lvl+1)
	}

	acc = append(acc, ll...)
	acc = append(acc, lr...)
	acc = append(acc, n.String())

	return acc, u.Max(lvlL, lvlR)
}
