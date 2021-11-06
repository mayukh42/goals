package algo

import (
	"fmt"
	"log"
)

/** TODO: 2D mat struct?
 */

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

/** linear saddlaback search
 * elements do not overlap with prev row/ col, i.e.
 * 	A[i][j] <= A[i][j+1]
 * 	A[i][j] <= A[i+1][j]
 * 	A[i][j] <= A[i+1][j+1]
 * O(n + m), n rows and m cols
 */
func SaddlebackSearchL(mat [][]int, r, c, x int) (int, int) {
	// find rows i and j such that ri_max <= x <= rj_max
	var i, j = 0, 0
	for i = 0; i < r-1; i++ {
		if mat[i][c-1] < x && mat[i+1][c-1] >= x {
			// x is in row i+1
			j = i + 1
			break
		}
	}
	// search through cols of row j
	for i = 0; i < c; i++ {
		if mat[j][i] == x {
			return j, i
		}
	}
	return -1, -1
}

func MatAt(mat [][]int, r, c, i, j int) (int, error) {
	// return element of 2D matrix rxc at i, j
	if i < r && j < c {
		return mat[i][j], nil
	}
	return 0, fmt.Errorf("index out of range")
}

func SaddlebackSearch(mat [][]int, r, c, x int) (int, int) {
	/** start from top right i,j (or bottom left) p,q
	 * and advance 1 step left/ down
	 * or right/ up
	 * to shrink the search space
	 * ranges:
	 * 	i = [0..r-1]
	 * 	j = [c-1..0]
	 * 	p = [r-1..0]
	 * 	q = [0..c-1]
	 */
	var i, j, p, q = 0, c - 1, r - 1, 0
	for mat[i][j] != x {
		// walk from top right
		if i < r && j >= 0 && x > mat[i][j] {
			i++
		}
		if i < r && j >= 0 && x < mat[i][j] {
			j--
		}
		// walk from bottom left
		if p >= 0 && q < c && x < mat[p][q] {
			p--
		}
		if p >= 0 && q < c && x > mat[p][q] {
			q++
		}
		// break before next iteration if necessary
		if i >= r || j < 0 || p < 0 || q >= c {
			break
		}
	}

	log.Printf("i, j: %d, %d", i, j)
	log.Printf("p, q: %d, %d", p, q)

	/** if first row of occurrence is desired, check for mat[i,j]
	 * 		return saddled back coords from top right, i.e. i,j,
	 * 	else if first col of occurrence is desired, check for mat[p,q]
	 * 		return saddled back coords from bottom left, i.e. p,q
	 */
	// trick to return -1, -1 if element is not found
	if i >= r || j < 0 {
		return -1, -1
	}
	return i, j

}
