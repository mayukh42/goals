package algo

import (
	"fmt"
	"log"
	"strings"

	"github.com/mayukh42/goals/utils"
)

const (
	PAD_BYT = '0'
)

/** radix sort strings list
 */
func FauxLinearSort(xs []string) []string {
	mi := maxLen(xs)
	sc := sumChars(xs)
	log.Printf("original: \n\t %v, \n\ttotal chars %d, \n\tmax str %s, \n\tlength %d", xs, sc, xs[mi], len(xs[mi]))

	ys := binSort(xs, 0, len(xs[mi]))
	// log.Printf("binsort%d: \n\t %v", 0, ys)
	return ys
}

/** padString(string, length)
 * if length > string.length, pad with 0's
 */
func pad(x string, l int) string {
	pl := l - len(x)
	var sb strings.Builder
	sb.WriteString(x)
	if pl > 0 {
		for i := 0; i < pl; i++ {
			sb.WriteByte(PAD_BYT)
		}
	}
	return sb.String()
}

/** getIndex()
 * assume pad char is lexicographically smaller than others
 * if not padding, caller should subtract 1 from the return
 */
func getIndex(b byte) int {
	if b == PAD_BYT {
		return 0
	}
	return int(b) - 96
}

/** assume alphabets in {0, a...z}
 * bin each string based on char i in an array of 26 length
 */
func binSort(xs []string, i, n int) []string {
	// base case
	if i >= n {
		return xs
	}
	// fmt.Printf("i = %d, n = %d\n", i, n)

	fc := false
	// OPT: check if there is any non-pad char in this set
	for _, x := range xs {
		if len(x) > i {
			fc = true
		}
	}

	if !fc {
		// no point in recursing further as we did not get a single non-pad char
		return xs
	}

	bins := make([][]string, 26)
	for _, x := range xs {
		var c byte
		if len(x) > i {
			c = x[i]
		} else {
			continue
		}
		j := getIndex(c) - 1
		// prepend new to head of list
		if bins[j] == nil {
			bins[j] = []string{x}
		} else {
			bins[j] = append(bins[j], x)
		}
	}

	// unpack the bin and recurse
	ys := make([]string, 0)
	for k := range bins {
		ss := binSort(bins[k], i+1, n)
		ys = append(ys, ss...)
	}
	return ys
}

/** maxLen()
 * returns index of max string length
 */
func maxLen(xs []string) int {
	max := ""
	mi := 0
	for i, x := range xs {
		if len(x) > len(max) {
			max = x
			mi = i
		}
	}
	return mi
}

func sumChars(xs []string) int {
	n := 0
	for _, x := range xs {
		n += len(x)
	}
	return n
}

type Suffix string

func NewSuffix(s string) *Suffix {
	sf := Suffix(s)
	return &sf
}

func (s *Suffix) String() string {
	return string(*s)
}

type SuffixArray struct {
	arr  []Suffix
	size int
}

func NewSuffixArray(s string) *SuffixArray {
	n := len(s)
	sa := &SuffixArray{
		arr:  make([]Suffix, n),
		size: n,
	}

	for i := range s {
		sa.arr[i] = *NewSuffix(s[i:])
	}

	return sa
}

func (sa *SuffixArray) String() string {
	var sb strings.Builder
	sb.WriteString("[ \n\t")
	for i, a := range sa.arr {
		sb.WriteString(a.String())
		if i == sa.size-1 {
			sb.WriteString("\n]")
		}
		sb.WriteString("\n\t")
	}
	// sb.WriteString("\n]")
	return sb.String()
}

// TODO: use u.List instead of []Suffix
func (sa *SuffixArray) ToList() utils.List {
	l := make(utils.List, sa.size)
	for i := range sa.arr {
		l[i] = sa.arr[i]
	}
	return l
}

func (s Suffix) ComLen(t Suffix) int {
	p, q := s, t
	if len(t) < len(s) {
		p = t
		q = s
	}

	i := 0
	for i < len(p) {
		if p[i] != q[i] {
			break
		}
		i++
	}

	return i
}

func CompareSuffix(a, b utils.Any) int {
	var (
		x, y     Suffix
		xok, yok bool
	)

	x, xok = a.(Suffix)
	y, yok = b.(Suffix)
	if xok && yok {
		return strings.Compare(x.String(), y.String())
	}
	return -2
}

/** Permutations()
 * return n! permutations of a string of n chars
 */
func Permutations(s string) []string {
	pmts := []string{}
	if len(s) <= 1 {
		pmts = append(pmts, s)
		return pmts
	}

	//
	for i := range s {
		prfx := s[i]
		sfxs := Permutations(fmt.Sprintf("%s%s", s[:i], s[i+1:]))
		for _, sfx := range sfxs {
			pmts = append(pmts, fmt.Sprintf("%c%s", prfx, sfx))
		}
	}

	return pmts
}

func Combinations(s string) []string {
	cmbs := []string{""}
	if len(s) <= 0 {
		return cmbs
	}

	return Subsets(s, cmbs)
}

// return all subsets of s
func Subsets(s string, acc []string) []string {
	l := len(s)
	for i := 0; i < l; i++ {
		curr := fmt.Sprintf("%c", s[i])
		// incrementally create new subsets with the additional char
		us := unionSingle(acc, curr)
		acc = append(acc, us...)
	}

	return acc
}

func unionSingle(ss []string, s string) []string {
	ts := make([]string, len(ss))
	for i := range ss {
		ts[i] = fmt.Sprintf("%s%s", ss[i], s)
	}
	return ts
}

// TODO: Subsets of length n
