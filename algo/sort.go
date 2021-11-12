package algo

import (
	"log"

	"github.com/mayukh42/goals/utils"
)

/** poc, so using uint8
 * TODO: use u.Any
 */

func CountingSort(xs []uint8) []uint8 {
	min, max := uint8(255), uint8(0)
	for _, x := range xs {
		if x <= min {
			min = x
		}
		if x > max {
			max = x
		}
	}

	bins := make([]uint8, max-min+1)
	for _, x := range xs {
		i := x - min
		bins[i]++
	}

	cs := make([]uint8, 0)
	j := min
	for _, n := range bins {
		for k := n; k > 0; k-- {
			cs = append(cs, j)
		}
		j++
	}

	return cs
}

/** [0..n^2-1]
 * using counting sort
 */
func RadixSort(xs []uint8, n int) []uint8 {
	// radix sort will have fixed bin size
	// radix
	var cs []uint8
	r := 0
	bins := make([][]uint8, 10)
	for _, x := range xs {
		// extract dth digit from right
		ds := Digits(int(x))

		d := 0
		if len(ds) >= n {
			d = int(ds[r])
		}
		// log.Printf("digits(%d): %v, ds(%d): %d", x, ds, r, d)
		if bins[d] == nil {
			bins[d] = []uint8{x}
		} else {
			bins[d] = append(bins[d], x)
		}
	}
	// unpack
	for _, ns := range bins {
		ss := CountingSort(ns)
		cs = append(cs, ss...)
	}
	return cs
}

/** lo: lowest index of array
 * 	hi: length of segment (1 + highest index)
 * 	i.e. xs = [lo, hi)
 * Hoare's pacman scheme
 */
func PartitionInt(xs []int, lo, hi int) int {
	if lo >= hi {
		return -1
	}
	if hi-lo > len(xs) {
		return -1
	}

	// pivot is the mid element
	pi := lo + (hi-lo)>>1
	p := xs[pi]
	xs[pi] = xs[lo]
	xs[lo] = p
	i, j := lo+1, hi-1
	for {
		for i < hi && xs[i] <= p {
			i++
		}
		for j >= lo && xs[j] > p {
			j--
		}
		if i < j {
			// mismatch; swap if invariant of i < j is valid, and advance
			tmp := xs[i]
			xs[i] = xs[j]
			xs[j] = tmp

			i++
			j--
		} else {
			// i >= j, so the pacmen collided
			break
		}
	}
	/** i is the index after all elems are <= p, so swap p with i-1
	 * similarly, j is the index before all elems are > p,
	 * so alternatively swap p with j+1
	 */
	tmp := xs[i-1]
	xs[i-1] = xs[lo]
	xs[lo] = tmp
	return i - 1
}

func QuicksortInt(xs []int, lo, hi int) {
	if lo >= hi {
		return
	}
	// pick a partition at mid point
	pi := PartitionInt(xs, lo, hi)
	log.Printf("[qs] \n\tlo, hi, pi, xs: \n \t%d, %d, %d, %v",
		lo, hi, pi, xs,
	)
	QuicksortInt(xs, lo, pi)
	QuicksortInt(xs, pi+1, hi)
}

/** Partition()
 * Type-generic partition fn
 */
func Partition(xs utils.List, lo, hi int, compFn utils.ComparatorFn) int {
	if lo >= hi {
		return -1
	}
	if hi-lo > len(xs) {
		return -1
	}

	// pivot is the mid element
	pi := lo + (hi-lo)>>1
	p := xs[pi]
	xs[pi] = xs[lo]
	xs[lo] = p
	i, j := lo+1, hi-1
	for {
		for i < hi && compFn(xs[i], p) <= 0 {
			i++
		}
		for j >= lo && compFn(xs[j], p) > 0 {
			j--
		}
		if i < j {
			// mismatch; swap if invariant of i < j is valid, and advance
			tmp := xs[i]
			xs[i] = xs[j]
			xs[j] = tmp

			i++
			j--
		} else {
			// i >= j, so the pacmen collided
			break
		}
	}
	/** i is the index after all elems are <= p, so swap p with i-1
	 * similarly, j is the index before all elems are > p,
	 * so alternatively swap p with j+1
	 */
	tmp := xs[i-1]
	xs[i-1] = xs[lo]
	xs[lo] = tmp
	return i - 1
}

/** Quicksort()
 * Type-generic partition fn
 */
func Quicksort(xs utils.List, lo, hi int, compFn utils.ComparatorFn) {
	if lo >= hi {
		return
	}
	// pick a partition at mid point
	pi := Partition(xs, lo, hi, compFn)
	// log.Printf("[qs] \n\tlo, hi, pi, xs: \n \t%d, %d, %d, %v",
	// 	lo, hi, pi, xs,
	// )
	Quicksort(xs, lo, pi, compFn)
	Quicksort(xs, pi+1, hi, compFn)
}
