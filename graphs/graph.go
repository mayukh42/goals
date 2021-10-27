package graphs

import (
	"fmt"
	"math"

	u "github.com/mayukh42/goals/utils"
)

type Graph struct {
	Src      *Node
	Vertices map[string]*Node
	Edges    map[string]*Edge
}

/** NewGraph()
 * create edges set from label:[src, dst]
 */
func NewGraph(source u.Any, es map[u.Any]u.List) *Graph {
	edges := make(map[string]*Edge)

	// create node refs
	vs := make(map[string]*Node)

	for label, ends := range es {
		if len(ends) != 2 {
			continue
		}
		var (
			src, dst *Node
			ok       bool
		)
		e0 := fmt.Sprintf("%v", ends[0])
		e1 := fmt.Sprintf("%v", ends[1])
		if _, ok = vs[e0]; ok {
			src = vs[e0]
		} else {
			src = NewNode(ends[0])
		}

		if _, ok = vs[e1]; ok {
			dst = vs[e1]
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
	s := fmt.Sprintf("%v", source)
	if _, ok := vs[s]; ok {
		src = vs[s]
	}

	return &Graph{
		Src:      src,
		Edges:    edges,
		Vertices: vs,
	}
}

func (g *Graph) SetSource(src u.Any) {
	s := fmt.Sprintf("%v", src)
	if vn, ok := g.Vertices[s]; ok {
		g.Src = vn
	}
}

/** get adjacency list from sets G(V, E)
 * TODO: sort by weight
 */
func (g *Graph) AdjacencyList() map[u.Any]u.List {
	al := make(map[u.Any]u.List)
	for _, v := range g.Edges {
		// k = src.dst, v.Src, v.Dst etc
		var (
			ns u.List
			ok bool
		)
		if ns, ok = al[v.Src]; ok {
			ns = append(ns, v.Dst)
		} else {
			ns = u.List{v.Dst}
		}
		al[v.Src] = ns
		/** no need to do for dst separately for this edge, since
		 * it will be another edge if undirected
		 */
	}
	return al
}

/** get adjacency list from sets G(V, E)
 * TODO: sort by weight
 */
func (g *Graph) AdjacencyListSingleNode(nv u.Any) map[u.Any]u.List {
	al := make(map[u.Any]u.List)
	for _, v := range g.Edges {
		// k = src.dst, v.Src, v.Dst etc
		nvs := fmt.Sprintf("%v", nv)
		if v.Src.Id != nvs {
			continue
		}
		var (
			ns u.List
			ok bool
		)
		if ns, ok = al[v.Src]; ok {
			ns = append(ns, v.Dst)
		} else {
			ns = u.List{v.Dst}
		}
		al[v.Src] = ns
		/** no need to do for dst separately for this edge, since
		 * it will be another edge if undirected
		 */
	}
	return al
}

func (g *Graph) NearestNeighbor(v u.Any) *Node {
	var (
		nn   *Node
		dist float64
	)
	dist = math.MaxFloat64
	for _, e := range g.Edges {
		vs := fmt.Sprintf("%v", v)
		if e.Src.ValueString() != vs {
			continue
		}
		// e0 matches v, so e1 is a neighbor
		if e.Weight < dist {
			nn = e.Dst
			dist = e.Weight
		}
	}
	return nn
}
