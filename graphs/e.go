package graphs

import (
	"fmt"

	u "github.com/mayukh42/goals/utils"
)

/** Edge
 * for undirected graphs, there should be 2 edges with src/ dst flipped
 */
type Edge struct {
	Id     string
	Src    *Node
	Dst    *Node
	Label  u.Any
	Color  u.Any
	Weight float64
}

func NewEdge(src, dst *Node, label, color u.Any, w float64) *Edge {
	return &Edge{
		Id:     fmt.Sprintf("%v.%v", src.Value, dst.Value),
		Src:    src,
		Dst:    dst,
		Label:  label,
		Color:  color,
		Weight: w,
	}
}

func (e *Edge) ValueString() string {
	return fmt.Sprintf("%v--%v/%v(%.2f)-->%v", e.Src, e.Label, e.Color, e.Weight, e.Dst)
}
