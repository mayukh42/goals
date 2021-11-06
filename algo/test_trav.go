package algo

import (
	"fmt"

	g "github.com/mayukh42/goals/graphs"
	t "github.com/mayukh42/goals/tree"
	u "github.com/mayukh42/goals/utils"
)

func testTreeBFS() {
	xs := u.List{1, 2, 3, 4, 5, 6, 7}
	tr := t.CreateFromArray(xs)

	fmt.Printf("root: %v\n", tr)
	ys := tr.Infix(u.List{})
	fmt.Printf("infix: %v\n", ys)

	acc := BFSTree(tr)
	fmt.Printf("level order: %v\n", acc)
}

func testTreeDFS() {
	xs := u.List{1, 2, 3, 4, 5, 6, 7}
	tr := t.CreateFromArray(xs)

	fmt.Printf("root: %v\n", tr)
	ys := tr.Infix(u.List{})
	fmt.Printf("infix: %v\n", ys)

	acc := DFSTree(tr)
	fmt.Printf("depth order: %v\n", acc)
}

func testGraphBFS() {
	gr := g.CreateGraphAL()
	l := BFSGraph(gr.Src)

	fmt.Printf("bfs order: %v\n", l)
}
