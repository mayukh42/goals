package list

import (
	"log"

	u "github.com/mayukh42/goals/utils"
)

/** Rotate()
 * {1, 2, 3, 4, 5, 6, 7, 8, 9}
 * k = 2
 * {8, 9, 1, 2, 3, 4, 5, 6, 7}
 */
func Rotate(xs u.List, k uint) u.List {
	n := uint(len(xs))
	if k > n {
		k = k % n
	}
	log.Printf("[rotate]: n = %d, k = %d\n", n, k)

	ys := make(u.List, n)
	var i uint
	for i = 0; i < n; i++ {
		j := (i + k) % n
		ys[j] = xs[i]
	}
	return ys
}

/** RotateInPlace()
 * {1, 2, 3, 4, 5, 6, 7, 8}
 * k = 2
 * {7, 8, 1, 2, 3, 4, 5, 6}
 * * prefix length = n-k
 */
func RotateInPlace(xs u.List, k uint) u.List {
	n := uint(len(xs))
	if k > n {
		k = k % n
	}
	log.Printf("[rotate]: n = %d, k = %d\n", n, k)

	var i uint

	// rotate full array
	for i = 0; i < n/2; i++ {
		tmp := xs[i]
		xs[i] = xs[n-1-i]
		xs[n-1-i] = tmp
	}

	log.Printf("%v\n", xs)

	// rotate k elements (prefix)
	for i = 0; i < k/2; i++ {
		tmp := xs[i]
		xs[i] = xs[k-1-i]
		xs[k-1-i] = tmp
	}

	log.Printf("%v\n", xs)

	// rotate n-k elements (suffix)
	for i = k; i < k+(n-k)/2; i++ {
		tmp := xs[i]
		// i >= k. to compare prev, consider k=0 in that
		xs[i] = xs[n-1-(i-k)]
		xs[n-1-(i-k)] = tmp
	}

	log.Printf("%v\n", xs)

	return xs
}

/** RotateInPlace()
 * {1, 2, 3, 4, 5, 6, 7, 8}
 * k = 2
 * {7, 8, 1, 2, 3, 4, 5, 6}
 * order of reverses is diff w/ prev fn
 * prefix length = k
 */
func SwapPrefixSuffix(xs u.List, k uint) u.List {
	n := uint(len(xs))
	if k > n {
		k = k % n
	}
	log.Printf("[swap prefix/suffix]: n = %d, k = %d\n", n, k)

	var i uint

	// rotate k elements (prefix)
	for i = 0; i < k/2; i++ {
		tmp := xs[i]
		xs[i] = xs[k-1-i]
		xs[k-1-i] = tmp
	}

	log.Printf("%v\n", xs)

	// rotate n-k elements (suffix)
	for i = k; i < k+(n-k)/2; i++ {
		tmp := xs[i]
		// i >= k. to compare prev, consider k=0 in that
		xs[i] = xs[n-1-(i-k)]
		xs[n-1-(i-k)] = tmp
	}

	log.Printf("%v\n", xs)

	// rotate full array
	for i = 0; i < n/2; i++ {
		tmp := xs[i]
		xs[i] = xs[n-1-i]
		xs[n-1-i] = tmp
	}

	log.Printf("%v\n", xs)

	return xs
}

func ReverseArr(xs u.List) u.List {
	l := len(xs)
	for i := 0; i < (l >> 1); i++ {
		tmp := xs[i]
		xs[i] = xs[l-i-1]
		xs[l-i-1] = tmp
	}
	return xs
}
