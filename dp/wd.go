package dp

import (
	"strings"

	u "github.com/mayukh42/goals/utils"
)

var COST = map[string]int{
	"sub": 1,
	"ins": 1,
	"del": 1,
}

/** Levenshtein1 Word Distance
 * -> insert
 * |  delete
 * -  substitute
 */
func Levenshtein1(w1, w2 string) int {
	if len(w1) > u.MAX_DIM || len(w2) > u.MAX_DIM {
		//
		return -1
	}

	// normalize
	w1 = strings.ToLower(w1)
	w2 = strings.ToLower(w2)

	// transpose
	tmp := w1
	w1 = w2
	w2 = tmp

	mat := [u.MAX_DIM][u.MAX_DIM]int{}

	l1 := len(w1)
	l2 := len(w2)

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			xd, yd := i-1, j-1
			// index checks

			// diagonal (nw)
			if xd < 0 {
				xd = 0
			}
			if yd < 0 {
				yd = 0
			}

			// nw, w, n
			nw := mat[xd][yd]

			w := mat[xd][j]
			// avoid using mat[i][j]
			if xd == i {
				w = nw
			}

			// avoid using mat[i][j]
			n := mat[i][yd]
			if yd == j {
				n = nw
			}

			// makes sense only if all costs are same
			min := int(u.Min3(uint(nw), uint(w), uint(n)))
			if w1[i] == w2[j] {
				// no added cost, just min
				mat[i][j] = min
			} else {
				// cost in all directions same
				mat[i][j] = 1 + min
			}
		}
	}

	u.Print2DMatrix(mat, l1, l2)

	// answer [TODO: construct Hamming Distance]
	return mat[l1-1][l2-1]
}

func WordChangeCost(nw, n, w int, add bool) int {
	// if all costs same, then no further cost too
	if !add {
		return int(u.Min3(uint(nw), uint(w), uint(n)))
	}
	// nw: sub, n: ins, w: del
	sub := uint(nw + COST["sub"])
	ins := uint(n + COST["ins"])
	del := uint(w + COST["del"])

	return int(u.Min3(sub, ins, del))
}

/** Levenshtein Word Distance
 * -> insert
 * |  delete
 * -  substitute
 */
func Levenshtein(w1, w2 string) int {
	if len(w1) > u.MAX_DIM || len(w2) > u.MAX_DIM {
		//
		return -1
	}

	// normalize
	w1 = strings.ToLower(w1)
	w2 = strings.ToLower(w2)

	// transpose
	tmp := w1
	w1 = w2
	w2 = tmp

	mat := [u.MAX_DIM][u.MAX_DIM]int{}

	l1 := len(w1)
	l2 := len(w2)

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			xd, yd := i-1, j-1
			// index checks

			// diagonal (nw)
			if xd < 0 {
				xd = 0
			}
			if yd < 0 {
				yd = 0
			}

			// nw, w, n
			nw := mat[xd][yd]

			w := mat[xd][j]
			// avoid using mat[i][j]
			if xd == i {
				w = nw
			}

			// avoid using mat[i][j]
			n := mat[i][yd]
			if yd == j {
				n = nw
			}

			if w1[i] == w2[j] {
				// no added cost over nw
				mat[i][j] = WordChangeCost(nw, n, w, false)
			} else {
				// find cost from map, if all = 1 then Levenshtein1
				mat[i][j] = WordChangeCost(nw, n, w, true)
			}
		}
	}

	u.Print2DMatrix(mat, l1, l2)

	// answer [TODO: construct Hamming Distance]
	return mat[l1-1][l2-1]
}
