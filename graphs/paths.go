package graphs

import (
	c "github.com/mayukh42/goals/collections"
	u "github.com/mayukh42/goals/utils"
)

// TODO: return path, error
func DjikstraSP(g *Graph, src, dst u.Any) (u.List, float64) {
	path := make(u.List, 0)
	start := g.GetNodeById(src)
	if start == nil {
		return path, 0.0
	}

	end := g.GetNodeById(dst)
	if end == nil {
		path = append(path, start)
		return path, 0.0
	}

	q := c.NewQueue()
	q.Enqueue(start)
	totalDist := 0.0
	for !q.IsEmpty() {
		e, _ := q.Dequeue()
		node, ok := e.(*Node)
		if !ok {
			return path, totalDist
		}

		if node.Equals(end) {
			path = append(path, node)
			return path, totalDist
		}

		nn, dist := g.NearestNeighbor(node.Id)
		if nn == nil {
			return path, totalDist
		}

		totalDist += dist
		q.Enqueue(nn)
		path = append(path, nn)
	}

	return path, totalDist
}
