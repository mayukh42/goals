package algo

import (
	"log"
	"strings"
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
