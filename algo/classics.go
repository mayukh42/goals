package algo

import (
	"log"

	u "github.com/mayukh42/goals/utils"
)

func BinarySearch(xs u.List, x u.Any, cf u.ComparatorFn, lo, hi int) bool {
	if len(xs) == 0 {
		return false
	}

	if lo > hi {
		return false
	}

	mi := lo + (hi-lo)/2
	log.Printf("lo, mi, hi = %d, %d, %d", lo, mi, hi)

	if cf(x, xs[mi]) < 0 {
		// search left
		return BinarySearch(xs, x, cf, lo, mi-1)
	} else if cf(x, xs[mi]) > 0 {
		// search right
		return BinarySearch(xs, x, cf, mi+1, hi)
	} else {
		// found
		return true
	}
}

/** LRSeq()
 * longest repeated subsequence
 * in a string
 */
func LRSeq(s string) string {
	sa := NewSuffixArray(s)
	log.Printf("suffix array of %s: %v", s, sa)

	ss := sa.ToList()
	Quicksort(ss, 0, sa.size, CompareSuffix)

	// find index having max common length with next neighbor
	comlen, idx := 0, 0
	for i := 0; i < sa.size-1; i++ {
		j := i + 1
		var (
			x, y Suffix
		)
		x, _ = ss[i].(Suffix)
		y, _ = ss[j].(Suffix)
		comxy := x.ComLen(y)
		// log.Printf("%d: comlen(%s,%s): %d", comlen, x, y, comxy)

		if comxy > comlen {
			comlen = comxy
			idx = i
		}
	}

	return u.AnyToString(ss[idx])
}
