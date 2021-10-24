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
