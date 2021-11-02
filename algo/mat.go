package algo

import "log"

/** SpiralWalk()
 * right -> down -> left -> up
 * until all elements are visited
 * of a 2D array/ matrix
 */
func SpiralWalk(xy [][]int, m, n int) []int {
	v := 0
	walk := make([]int, 0)
	minx, miny := 0, 0
	maxx, maxy := m-1, n-1
	for v < m*n {
		// right
		for i := minx; i <= maxx; i++ {
			walk = append(walk, xy[miny][i])
			v++
			log.Printf("right: %d.%d", v, xy[miny][i])
		}
		// down
		miny++
		for j := miny; j <= maxy; j++ {
			walk = append(walk, xy[j][maxx])
			v++
			log.Printf("down: %d.%d", v, xy[j][maxx])
		}
		// left
		maxx--
		for i := maxx; i >= minx; i-- {
			walk = append(walk, xy[maxy][i])
			v++
			log.Printf("left: %d.%d", v, xy[maxy][i])
		}
		// up
		maxy--
		for j := maxy; j >= miny; j-- {
			walk = append(walk, xy[j][minx])
			v++
			log.Printf("up: %d.%d", v, xy[j][minx])
		}
		minx++
	}
	return walk
}
