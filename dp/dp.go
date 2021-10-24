package dp

import (
	"fmt"
	"log"
	"strings"

	u "github.com/mayukh42/goals/utils"
)

type Seq struct {
	Value   int
	Indices []int
}

func NewSeq() Seq {
	return Seq{
		Indices: []int{},
	}
}

func (s *Seq) Add(i int) Seq {
	s.Indices = append(s.Indices, i)
	return *s
}

func (s *Seq) String() string {
	return fmt.Sprintf("%v", s.Indices)
}

func MaxSeq(seq1, seq2 Seq) Seq {
	l1 := len(seq1.Indices)
	l2 := len(seq2.Indices)
	if l1 > l2 {
		return seq1
	} else {
		return seq2
	}
}

func (s *Seq) CreateSeq(str string) string {
	var b strings.Builder
	for _, i := range s.Indices {
		if i >= 0 && i < len(str) {
			b.WriteByte(str[i])
		}
	}

	log.Printf("sequence range: %v\n", s)
	return b.String()
}

/** LCSeq()
 * returns longest common ubsequence (not substring) of 2 strings
 * 2D mat = classic [m+1][n+1]
 */
func LCSeq(s1, s2 string) string {
	if len(s1) >= u.MAX_DIM || len(s2) >= u.MAX_DIM {
		//
		return ""
	}

	// normalize
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	// transpose
	// tmp := s1
	// s1 = s2
	// s2 = tmp

	mat := [u.MAX_DIM][u.MAX_DIM]Seq{}

	l1 := len(s1)
	l2 := len(s2)

	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			if i == 0 || j == 0 {
				// default
				mat[i][j] = NewSeq()
				continue
			}
			// we already have a seq, mat[i-1][j-1] exists
			x, y := i-1, j-1
			if s1[x] == s2[y] {
				// append the char pos from s1 (we'll later use s1 to create seq)
				mat[i][j] = mat[i-1][j-1].Add(x)
			} else {
				// max(mat[i][j-1], mat[i-1][j])
				mat[i][j] = MaxSeq(mat[i][j-1], mat[i-1][j])
			}
		}
	}

	seq := mat[l1][l2]
	return seq.CreateSeq(s1)
}

/** MaximalSum()
 * returns subarray with max sum or empty array (i.e. sum = 0)
 *
 * [-2, 1, -3, 4, -1, 2, 1, -5, 4] -> [4, −1, 2, 1], 6
 *   0, 1, 0, 4, 3, 5, 6, 1, 5 - maxendshere
 *   0, 1, 1, 4, 4, 5, 6, 6, 6 - maxsofar
 * 		monotonically increasing fn
 * 		to get range, fintd index where max val first seen (== end),
 * 		compute sum by backtracking to arrive at start
 *
 * [1, 2, 3, -2, 9, -1]
 *  1, 3, 6, 4, 13, 12
 *  1, 3, 6, 6, 13, 13
 *
 * Note: we can use just 2 rolling vars to solve the problem,
 * 	but we use Seq to walk through/ visualize the solution
 */
func MaximalSum(xs []int) int {
	l := len(xs)
	if l > u.MAX_DIM {
		return u.MIN_INT
	}

	mat := [u.MAX_DIM]Seq{}

	/** invariant: mat[i] = maximal sum (w/ indices for range) up to i
	 * mat[i].Value = maxsofar
	 * mat[i].Indices = {start, end}
	 */
	maxendshere := 0
	for i, x := range xs {
		mat[i] = NewSeq()

		// empty set has sum = 0
		if maxendshere+x > 0 {
			maxendshere += x
		} else {
			maxendshere = 0
			// potential start index
		}

		// index check
		if i == 0 {
			mat[i].Value = maxendshere
			// add index only if sum exceeds that of an empty array
			if maxendshere > 0 {
				mat[i].Indices = []int{i}
			}
		} else {
			if maxendshere > mat[i-1].Value {
				mat[i].Value = maxendshere
				// maxsofar changed, so record this index too
				mat[i].Indices = []int{i}
			} else {
				mat[i].Value = mat[i-1].Value
				mat[i].Indices = mat[i-1].Indices
			}
		}
	}

	// backtrack to find range
	sum := mat[l-1].Value
	j := mat[l-1].Indices[0] + 1
	i := j - 1
	for sum > 0 {
		sum -= xs[i]
		if sum > 0 {
			i--
		}
	}

	log.Printf("range: [%d:%d] = %v, max sum: %d\n", i, j, xs[i:j], mat[l-1].Value)
	return mat[l-1].Value
}
