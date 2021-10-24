package graphs

import (
	"math"
	// u "github.com/mayukh42/goals/utils"
)

type Point struct {
	X int
	Y int
	Z int
}

func NewPoint(x, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

func (p *Point) CompareX(p_ *Point) int {
	// -1, 0, 1
	if p_.X < p.X {
		return -1
	} else if p_.X > p.X {
		return 1
	}
	return 0
}

func (p *Point) CompareY(p_ *Point) int {
	// -1, 0, 1
	if p_.Y < p.Y {
		return -1
	} else if p_.Y > p.Y {
		return 1
	}
	return 0
}

func (p *Point) Cartesian(p_ *Point) float64 {
	dX := float64(p.X - p_.X)
	dY := float64(p.Y - p_.Y)
	return math.Sqrt(math.Pow(dX, float64(2)) + math.Pow(dY, float64(2)))
}

func (p *Point) Move(dp *Point) *Point {
	return &Point{
		X: p.X + dp.X,
		Y: p.Y + dp.Y,
	}
}

var SHIFT_MAP = map[int]*Point{
	0: NewPoint(0, 1),
	1: NewPoint(1, 1),
	2: NewPoint(1, 0),
	3: NewPoint(1, -1),
	4: NewPoint(0, -1),
	5: NewPoint(-1, -1),
	6: NewPoint(-1, 0),
	7: NewPoint(-1, 1),
}

/** return all 1-step neighbors of p within the grid
 * 		n	ne	e	se	s	sw	w	nw
 * 		0 	1	2	3	4	5	6	7
 */
func (p *Point) Neighbors(r *RectangleGrid) []*Point {
	// at most there can be 8 neighbors/ walkable cells
	ns := make([]*Point, 8)
	for i := range ns {
		dp := SHIFT_MAP[i]
		pn := p.Move(dp)
		if r.IsInside(pn) {
			ns[i] = pn
		} else {
			ns[i] = nil
		}
	}
	return ns
}

/** ClosestTarget()
 * Neighbor thatis closest to target
 */
func (p *Point) ClosestTarget(r *RectangleGrid, t *Point) *Point {
	if p.CompareX(t) == 0 && p.CompareY(t) == 0 {
		// p == t
		return t
	}

	ns := p.Neighbors(r)
	var c *Point = nil
	var d = math.MaxFloat64

	for _, n := range ns {
		if n == nil {
			continue
		}
		dist := p.Cartesian(n)
		if dist < d {
			c = n
		}
	}

	return c
}

type RectangleGrid struct {
	SW *Point
	NE *Point
}

func NewRectangleGrid(lowerLeft, upperRight *Point) *RectangleGrid {
	return &RectangleGrid{
		SW: lowerLeft,
		NE: upperRight,
	}
}

func (r *RectangleGrid) IsInside(p *Point) bool {
	return p.X <= r.NE.X && p.X >= r.SW.X && p.Y <= r.NE.Y && p.Y >= r.SW.Y
}

func (r *RectangleGrid) ShortestWalk(a, b *Point) []*Point {
	path := make([]*Point, 0)
	c := a
	for c.CompareX(b) != 0 && c.CompareY(b) != 0 {
		p := c.ClosestTarget(r, b)
		path = append(path, p)
		c = p
	}

	return path
}
