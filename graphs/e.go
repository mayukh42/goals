package graphs

import (
	"fmt"

	u "github.com/mayukh42/goals/utils"
)

/** Edge
 * for undirected graphs, there should be 2 edges with src/ dst flipped
 */
type Edge struct {
	Src   *Node
	Dst   *Node
	Label u.Any
	Color u.Any
}

func NewEdge(src, dst *Node, label, color u.Any) *Edge {
	return &Edge{
		Src:   src,
		Dst:   dst,
		Label: label,
		Color: color,
	}
}

func (e *Edge) ValueString() string {
	return fmt.Sprintf("%v--%v/%v-->%v", e.Src, e.Label, e.Color, e.Dst)
}
