package tree

import (
	u "github.com/mayukh42/goals/utils"
)

/** Least Common Ancestor
 * BST
 */
func LCA(r, a, b *Node, f u.ComparatorFn) *Node {
	if f(a.Value, r.Value) < 0 && f(b.Value, r.Value) < 0 {
		// LCA is in left subtree
		if r.Left != nil {
			return LCA(r.Left, a, b, f)
		}
	} else if f(a.Value, r.Value) > 0 && f(b.Value, r.Value) > 0 {
		// LCA is in right subtree
		if r.Right != nil {
			return LCA(r.Right, a, b, f)
		}
	}

	// a <= r <= b, current node is LCA
	return r
}
