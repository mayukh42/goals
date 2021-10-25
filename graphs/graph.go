package graphs

import (
	"fmt"

	u "github.com/mayukh42/goals/utils"
)

type VGraph struct {
	Src      *Node
	Vertices map[u.Any]*Node
}

func NewVGraph(src u.Any, alm map[u.Any]u.List) *VGraph {
	if len(alm) == 0 {
		return nil
	}
	g := &VGraph{
		Vertices: make(map[u.Any]*Node),
	}
	for k := range alm {
		n := NewNode(k)
		g.Vertices[k] = n
	}

	for k, v := range alm {
		nk := g.Vertices[k]
		for _, n := range v {
			nv := g.Vertices[n]
			nk.AList = append(nk.AList, nv)
		}
		if k == src {
			g.Src = nk
		}
	}
	return g
}

type EGraph struct {
	Src   *Node
	Edges map[u.Any]*Edge
}

func NewEGraph(source u.Any, es map[u.Any]u.List) *EGraph {
	edges := make(map[u.Any]*Edge)

	// create node refs
	vs := make(map[u.Any]*Node)

	for label, ends := range es {
		if len(ends) != 2 {
			continue
		}
		var (
			src, dst *Node
			ok       bool
		)
		if _, ok = vs[ends[0]]; ok {
			src = vs[ends[0]]
		} else {
			src = NewNode(ends[0])
		}

		if _, ok = vs[ends[1]]; ok {
			dst = vs[ends[1]]
		} else {
			dst = NewNode(ends[1])
		}
		// label set key is a function of the 2 endpoints
		key := fmt.Sprintf("%v.%v", src.Value, dst.Value)
		edges[key] = &Edge{
			Src:   src,
			Dst:   dst,
			Label: label,
		}
	}

	var src *Node
	if _, ok := vs[source]; ok {
		src = vs[source]
	}
	return &EGraph{
		Src:   src,
		Edges: edges,
	}
}

type Graph struct {
	Src      *Node
	Vertices map[u.Any]*Node
	Edges    map[u.Any]*Edge
}

/** NewGraph()
 * create edges set from label:[src, dst]
 */
func NewGraph(source u.Any, es map[u.Any]u.List) *Graph {
	edges := make(map[u.Any]*Edge)

	// create node refs
	vs := make(map[u.Any]*Node)

	for label, ends := range es {
		if len(ends) != 2 {
			continue
		}
		var (
			src, dst *Node
			ok       bool
		)
		if _, ok = vs[ends[0]]; ok {
			src = vs[ends[0]]
		} else {
			src = NewNode(ends[0])
		}

		if _, ok = vs[ends[1]]; ok {
			dst = vs[ends[1]]
		} else {
			dst = NewNode(ends[1])
		}
		// label set key is a function of the 2 endpoints
		key := fmt.Sprintf("%v.%v", src.Value, dst.Value)
		edges[key] = &Edge{
			Src:   src,
			Dst:   dst,
			Label: label,
		}
	}

	var src *Node
	if _, ok := vs[source]; ok {
		src = vs[source]
	}
	return &Graph{
		Src:      src,
		Edges:    edges,
		Vertices: vs,
	}
}

func (g *Graph) SetSource(s u.Any) {
	if vn, ok := g.Vertices[s]; ok {
		g.Src = vn
	}
}
