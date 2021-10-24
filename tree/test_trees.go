package tree

import (
	"fmt"

	u "github.com/mayukh42/goals/utils"
)

func CreateUnbalanced() *Node {
	return nil
}

func CreateAndInfix() {
	xs := u.List{1, 2, 3, 4, 5, 6, 7}
	tr := CreateFromArray(xs)

	fmt.Printf("root: %v\n", tr)
	ys := tr.Infix(u.List{})
	fmt.Printf("infix: %v\n", ys)
}

func CreateAndPrefix() {
	xs := u.List{1, 2, 3, 4, 5, 6, 7}
	tr := CreateFromArray(xs)

	fmt.Printf("root: %v\n", tr)
	ys := tr.Prefix(u.List{})
	fmt.Printf("prefix: %v\n", ys)
}

func CreateAndPostfix() {
	xs := u.List{1, 2, 3, 4, 5, 6, 7}
	tr := CreateFromArray(xs)

	fmt.Printf("root: %v\n", tr)
	ys, h := tr.PostfixHeight(u.List{}, 1)
	fmt.Printf("height: %d\npostfix: %v\n", h, ys)
}

func HeightBalanced() {
	xs := u.List{1, 2, 3, 4, 5, 6, 7}
	tr := CreateFromArray(xs)

	fmt.Printf("root: %v\n", tr)
	h := tr.Height(1) // since tr is not nil
	isB := tr.IsBalanced()
	fmt.Printf("height: %d\nbalanced: %v\n", h, isB)
}

func TestLCA() {
	xs := u.List{1, 2, 3, 4, 5, 6, 7}
	tr := CreateFromArray(xs)
	ai, bi := 6, 7
	a, b := NewNode(ai), NewNode(bi)

	fmt.Printf("root: %v\n", tr)
	lca := LCA(tr, a, b, u.CompareInt)
	fmt.Printf("lca(%d, %d): %v\n", ai, bi, lca)
}
